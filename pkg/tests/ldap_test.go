package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/molleer/ldap-sync/pkg/config"
	"github.com/molleer/ldap-sync/pkg/services"
	"github.com/spf13/viper"
)

func TestConnection(t *testing.T) {
	err := config.LoadConfig()
	assert.Equal(t, err, nil, "Failed to load config")

	l, err := services.LDAPConnect()
	assert.Equal(t, err, nil, "Failed to dial LDAP server")
	defer l.Close()

	err = services.LoginUser(l, services.LoginConfig{
		UserName: viper.GetString("ldap.user"),
		Password: viper.GetString("ldap.password"),
	})
	assert.Equal(t, err, nil, "Failed to login as admin")
}
