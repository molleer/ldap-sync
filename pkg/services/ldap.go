package services

import (
	"crypto/tls"
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
