package main

import (
	"crypto/tls"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/molleer/ldap-sync/pkg/config"
	"github.com/molleer/ldap-sync/pkg/services"
	"github.com/spf13/viper"
	"gopkg.in/ldap.v2"
)

func TestConnection(t *testing.T) {
	err := config.LoadConfig()
	assert.Equal(t, err, nil, "Failed to load config")

	config := services.LdapConfig{
		Url:        viper.GetString("ldap.url"),
		ServerName: viper.GetString("ldap.servername"),
		Tls:        viper.GetBool("ldap.tls"),
	}

	loginConfig := services.LoginConfig{
		UserName: viper.GetString("ldap.user"),
		Password: viper.GetString("ldap.password"),
	}

	l, err := ldap.DialTLS("tcp",
		config.Url,
		&tls.Config{ServerName: config.ServerName, InsecureSkipVerify: !config.Tls})

	assert.Equal(t, err, nil, "Failed to dial LDAP server")
	defer l.Close()

	l.Bind(loginConfig.UserName, loginConfig.Password)
	assert.Equal(t, err, nil, "Failed to login as admin")
}
