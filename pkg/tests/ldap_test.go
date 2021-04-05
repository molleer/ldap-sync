package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/molleer/ldap-sync/pkg/config"
	"github.com/molleer/ldap-sync/pkg/services"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var service *services.ServiceLDAP

func TestMain(m *testing.M) {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config")
		panic(err)
	}

	service, _ = services.NewLDAPService()
	service.LoginUser(services.LoginConfig{
		UserName: viper.GetString("ldap.user"),
		Password: viper.GetString("ldap.password"),
	})

	os.Exit(m.Run())
}

func TestConnection(t *testing.T) {
	l, err := services.NewLDAPService()
	assert.Equal(t, err, nil, "Failed to dial LDAP server")
	defer l.Connection.Close()

	err = l.LoginUser(services.LoginConfig{
		UserName: viper.GetString("ldap.user"),
		Password: viper.GetString("ldap.password"),
	})
	assert.Equal(t, err, nil, "Failed to login as admin")
}

func TestGetUsers(t *testing.T) {
	_, err := service.GetITUsers()
	assert.Equal(t, err, nil, "An error ocurred when fetching user")
}
