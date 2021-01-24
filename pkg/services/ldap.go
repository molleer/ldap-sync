package services

import (
	"crypto/tls"

	"github.com/spf13/viper"
	"gopkg.in/ldap.v2"
)

type LdapConfig struct {
	Url        string
	ServerName string
	Tls        bool
}

type LoginConfig struct {
	UserName string
	Password string
}

func LDAPConnect() (*ldap.Conn, error) {
	config := LdapConfig{
		Url:        viper.GetString("ldap.url"),
		ServerName: viper.GetString("ldap.servername"),
		Tls:        viper.GetBool("ldap.tls"),
	}

	return ldap.DialTLS("tcp",
		config.Url,
		&tls.Config{ServerName: config.ServerName, InsecureSkipVerify: !config.Tls})
}

func LoginUser(l *ldap.Conn, config LoginConfig) error {
	return l.Bind(config.UserName, config.Password)
}
