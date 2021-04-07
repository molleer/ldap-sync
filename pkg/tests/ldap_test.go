package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/molleer/ldap-sync/pkg/config"
	"github.com/molleer/ldap-sync/pkg/services"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var service *services.ServiceLDAP
var dummyUser = services.ITUser{
	Cid:            "wmacmak",
	Gdpr:           true,
	AcceptanceYear: 2002,
	FirstName:      "Wyatt",
	Email:          "wmacmak@student.chalmers.se",
	LastName:       "MacMakin",
	Nick:           "Chokladkaka",
	Phone:          "123456789",
}

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
	assert.NoError(t, err, "Failed to dial LDAP server")
	defer l.Connection.Close()

	err = l.LoginUser(services.LoginConfig{
		UserName: viper.GetString("ldap.user"),
		Password: viper.GetString("ldap.password"),
	})
	assert.NoError(t, err, "Failed to login as admin")
}

func TestGetUsers(t *testing.T) {
	_, err := service.GetITUsers()
	assert.NoError(t, err, "An error ocurred when fetching user")
}

func TestNextUid(t *testing.T) {
	uid, err := service.NextUid()
	assert.NoError(t, err, "An error accured while fetching next uidNumber")
	log.Printf("Next uid: %v\n", uid)
}

func TestAddUser(t *testing.T) {
	uid, err := service.NextUid()
	assert.NoError(t, err, "Cound not get next uidNumber when adding new user")
	err = service.AddITUser(dummyUser, uid)
	assert.NoError(t, err, "An error ocurred when adding users")
}

func TestGetUser(t *testing.T) {
	user, err := service.GetITUser(dummyUser.Cid)
	assert.NoError(t, err, "An error accured while fetching user")
	assert.Equal(t, dummyUser.Email, user.Email, "The wrong user was fetched")
}

func TestDeleteUser(t *testing.T) {
	err := service.DeleteUser(dummyUser.Cid)
	assert.NoError(t, err, "An error accured when deleting a user")
}
