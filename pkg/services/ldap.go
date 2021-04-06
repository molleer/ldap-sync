package services

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/spf13/viper"
	"gopkg.in/ldap.v2"
)

type ServiceLDAP struct {
	Connection   *ldap.Conn
	DBConfig     ServerConfig
	GroupsConfig EntryConfig
	UsersConfig  EntryConfig
}

type ServerConfig struct {
	Url        string
	ServerName string
}

type EntryConfig struct {
	BaseDN     string
	Filter     string
	Attributes []string
}

type CustomEntryConfig struct {
	BaseDN       string
	Filter       string
	ParentFilter string
	Attributes   []string
	Mail         string
}

func NewLDAPService() (*ServiceLDAP, error) {
	dbConfig := ServerConfig{
		Url:        viper.GetString("ldap.url"),
		ServerName: viper.GetString("ldap.servername"),
	}

	groupsConfig := EntryConfig{
		BaseDN:     viper.GetString("ldap.groups.basedn"),
		Filter:     viper.GetString("ldap.groups.filter"),
		Attributes: viper.GetStringSlice("ldap.groups.attributes"),
	}

	usersConfig := EntryConfig{
		BaseDN:     viper.GetString("ldap.users.basedn"),
		Filter:     viper.GetString("ldap.users.filter"),
		Attributes: viper.GetStringSlice("ldap.users.attributes"),
	}

	config := LdapConfig{
		Url:        viper.GetString("ldap.url"),
		ServerName: viper.GetString("ldap.servername"),
		Tls:        viper.GetBool("ldap.tls"),
	}

	l, err := ldap.DialTLS("tcp",
		config.Url,
		&tls.Config{ServerName: config.ServerName, InsecureSkipVerify: !config.Tls})

	return &ServiceLDAP{
		Connection:   l,
		DBConfig:     dbConfig,
		UsersConfig:  usersConfig,
		GroupsConfig: groupsConfig,
	}, err
}

func (s *ServiceLDAP) LoginUser(config LoginConfig) error {
	return s.Connection.Bind(config.UserName, config.Password)
}

func (s *ServiceLDAP) GetITUsers() ([]ITUser, error) {
	request := ldap.NewSearchRequest(
		s.UsersConfig.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		s.UsersConfig.Filter,
		[]string{"uid", "givenName", "sn", "acceptedUserAgreement", "admissionYear",
			"nickname", "mail", "telephoneNumber", "preferredLanguage"},
		nil,
	)

	users, err := s.Connection.Search(request)
	if err != nil {
		return nil, err
	}

	itUsers := []ITUser{}
	for _, entry := range users.Entries {
		userAgreement, _ := strconv.ParseBool(entry.GetAttributeValue("acceptedUserAgreement"))
		admissionYear, _ := strconv.Atoi(entry.GetAttributeValue("admissionYear"))
		itUsers = append(itUsers, ITUser{
			Cid:                   entry.GetAttributeValue("uid"),
			FirstName:             entry.GetAttributeValue("givenName"),
			LastName:              entry.GetAttributeValue("sn"),
			UserAgreement:         userAgreement,
			AcceptanceYear:        admissionYear,
			Nick:                  entry.GetAttributeValue("nickname"),
			Email:                 entry.GetAttributeValue("mail"),
			Phone:                 entry.GetAttributeValue("telephoneNumber"),
			Language:              entry.GetAttributeValue("preferredLanguage"),
			AccountLocked:         false,
			Activated:             true,
			Enabled:               true,
			AccountNonLocked:      true,
			CredentialsNonExpired: true,
		})
	}

	return itUsers, nil
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func (s *ServiceLDAP) AddITUser(user ITUser) error {
	/*
		dn: uid=bill,ou=people,dc=chalmers,dc=it
		changetype: add
		accepteduseragreement: TRUE
		admissionyear: 2001
		cn: %{firstname} '%{nickname}' %{lastname}
		gidnumber: 4500
		givenname: Bill
		homedirectory: /home/chalmersit/bill
		loginshell: /bin/bash
		mail: bill@student.chalmers.se
		nickname: Billy
		objectclass: chalmersstudent
		objectclass: posixAccount
		objectclass: top
		preferredlanguage: sv
		sn: Billysson
		telephonenumber: 07212345667
		uid: bill
		uidnumber: 10000
		userpassword: {SSHA}xeHcFuIwJ8R2JffkdS1YEgPUHzN67kvx
	*/

	gdpr := ""
	if user.Gdpr {
		gdpr = "TRUE"
	} else {
		gdpr = "FALSE"
	}

	return s.Connection.Add(&ldap.AddRequest{
		DN: fmt.Sprintf("uid=%s,%s", user.Cid, s.UsersConfig.BaseDN),
		Attributes: []ldap.Attribute{
			{Type: "accepteduseragreement", Vals: []string{gdpr}},
			{Type: "admissionyear", Vals: []string{strconv.FormatInt(int64(user.AcceptanceYear), 10)}},
			{Type: "cn", Vals: []string{"%{firstname} '%{nickname}' %{lastname}"}},
			{Type: "gidnumber", Vals: []string{"4500"}},
			{Type: "givenname", Vals: []string{user.FirstName}},
			{Type: "homedirectory", Vals: []string{fmt.Sprintf("/home/chalmersit/%s", user.Cid)}},
			{Type: "loginshell", Vals: []string{"/bin/bash"}},
			{Type: "mail", Vals: []string{user.Email}},
			{Type: "nickname", Vals: []string{user.Nick}},
			{Type: "objectclass", Vals: []string{"chalmersstudent", "posixAccount", "top"}},
			{Type: "sn", Vals: []string{user.LastName}},
			{Type: "telephonenumber", Vals: []string{user.Phone}},
			{Type: "uid", Vals: []string{user.Cid}},
			{Type: "uidnumber", Vals: []string{"10010"}},
			{Type: "userpassword", Vals: []string{fmt.Sprintf("{SSHA}%s", RandomString(20))}},
		},
	})

	/*request := ldap.NewSearchRequest(
		s.UsersConfig.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		s.UsersConfig.Filter,
		[]string{"uid", "givenName", "sn", "acceptedUserAgreement", "admissionYear",
			"nickname", "mail", "telephoneNumber", "preferredLanguage"},
		nil,
	)

	users, err := s.Connection.Search(request)
	if err != nil {
		return nil, err
	}

	itUsers := []ITUser{}
	for _, entry := range users.Entries {
		userAgreement, _ := strconv.ParseBool(entry.GetAttributeValue("acceptedUserAgreement"))
		admissionYear, _ := strconv.Atoi(entry.GetAttributeValue("admissionYear"))
		itUsers = append(itUsers, ITUser{
			Cid:                   entry.GetAttributeValue("uid"),
			FirstName:             entry.GetAttributeValue("givenName"),
			LastName:              entry.GetAttributeValue("sn"),
			UserAgreement:         userAgreement,
			AcceptanceYear:        admissionYear,
			Nick:                  entry.GetAttributeValue("nickname"),
			Email:                 entry.GetAttributeValue("mail"),
			Phone:                 entry.GetAttributeValue("telephoneNumber"),
			Language:              entry.GetAttributeValue("preferredLanguage"),
			AccountLocked:         false,
			Activated:             true,
			Enabled:               true,
			AccountNonLocked:      true,
			CredentialsNonExpired: true,
		})
	}

	return itUsers, nil*/
}
