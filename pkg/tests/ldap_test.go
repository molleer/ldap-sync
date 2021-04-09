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

func TestUpdateUser(t *testing.T) {
	newDummy := dummyUser
	newDummy.Email = "hello@chalmers.it"
	assert.NoError(t, service.UpdateUser(newDummy), "An error occured when updating user")
	user, err := service.GetITUser(newDummy.Cid)
	assert.NoError(t, err, "An error accured when fetching updated user")
	assert.Equal(t, newDummy.Email, user.Email, "The user email was not updated")
}

func TestAddSuperGroup(t *testing.T) {
	err := service.AddSuperGroup(dummySuperGroup)
	assert.NoError(t, err, "An error occured when adding new super group")
}

func TestGetSuperGroup(t *testing.T) {
	superGroup, err := service.GetSuperGroup(dummySuperGroup.Name)
	assert.NoError(t, err, "An error occured while fetching super group")
	assert.Equal(t, dummySuperGroup.Email, superGroup.Email, "The wrong super group was fetched")
}

func TestGetSuperGroups(t *testing.T) {
	superGroups, err := service.GetSuperGroups()
	assert.NoError(t, err, "An error occured when fetching all superg groups")
	assert.Contains(t, superGroups, dummySuperGroup, "Failed to fetch dmmy super group")
}

func TestAddGroup(t *testing.T) {
	err := service.AddGroup(dummyGroup)
	assert.NoError(t, err, "An error occured when adding new group")
}

func TestDeleteGroup(t *testing.T) {
	err := service.DeleteGroup(dummyGroup.Name, dummyGroup.SuperGroup.Name)
	assert.NoError(t, err, "An error occured while deleting group")
}

func TestDeleteSuperGroup(t *testing.T) {
	err := service.DeleteSuperGroup(dummySuperGroup.Name)
	assert.NoError(t, err, "An error occured when deleting super group")
}

func TestDeleteUser(t *testing.T) {
	err := service.DeleteUser(dummyUser.Cid)
	assert.NoError(t, err, "An error accured when deleting a user")
}
