package services

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/ldap.v2"
)

type FKITGroupDTO struct {
	ID              string         `json:"id"`
	BecomesActive   int64          `json:"becomesActive"`
	BecomesInactive int64          `json:"becomesInactive"`
	Description     SvEn           `json:"description"`
	Email           string         `json:"email"`
	Function        SvEn           `json:"function"`
	Name            string         `json:"name"`
	PrettyName      string         `json:"prettyName"`
	AvatarURL       interface{}    `json:"avatarURL"`
	SuperGroup      FKITSuperGroup `json:"superGroup"`
	Active          bool           `json:"active"`
}

type FKITGroup struct {
	ID               string         `json:"id"`
	BecomesActive    int64          `json:"becomesActive"`
	BecomesInactive  int64          `json:"becomesInactive"`
	Description      SvEn           `json:"description"`
	Email            string         `json:"email"`
	Function         SvEn           `json:"function"`
	Name             string         `json:"name"`
	PrettyName       string         `json:"prettyName"`
	AvatarURL        interface{}    `json:"avatarURL"`
	SuperGroup       FKITSuperGroup `json:"superGroup"`
	Active           bool           `json:"active"`
	GroupMembers     []FKITUser     `json:"groupMembers"`
	NoAccountMembers []interface{}  `json:"noAccountMembers"`
}

func getMembers(users []FKITUser, baseDN string) []string {
	memberStrings := make([]string, len(users))
	for i := 0; i < len(users); i++ {
		memberStrings[i] = fmt.Sprintf("uid=%s,%s", users[i].Cid, baseDN)
	}
	return memberStrings
}

func getPositions(users []FKITUser) []string {
	posts := make([]string, len(users))
	for i := 0; i < len(users); i++ {
		posts[i] = fmt.Sprintf("%s;%s", users[i].Post.Sv, users[i].Cid)
	}
	return posts
}

func (group *FKITGroup) ToLdapAttrib(gidNumber int, userBaseDN string) []ldap.Attribute {
	groupType := "committee"
	if group.SuperGroup.Type != "" {
		groupType = strings.ToLower(group.SuperGroup.Type)
	}

	attributes := []ldap.Attribute{
		{Type: "cn", Vals: []string{group.Name}},
		{Type: "description", Vals: []string{group.Description.Sv}},
		{Type: "displayname", Vals: []string{group.PrettyName}},
		{Type: "function", Vals: []string{group.Function.Sv}},
		{Type: "mail", Vals: []string{group.Email}},
		{Type: "member", Vals: getMembers(group.GroupMembers, userBaseDN)},
		{Type: "objectclass", Vals: []string{"groupOfNames", "itGroup", "top"}},
		{Type: "position", Vals: getPositions(group.GroupMembers)},
		{Type: "type", Vals: []string{groupType}},
	}

	/*TODO
	if production {
		attributes[6].Vals = append(attributes[6].Vals, "posixGroup")
		attributes = append(attributes, {
			Type: "gidnumber",
			Vals: []string{fmt.Sprintf("%v", uidNumber)}
		})
	}
	*/

	return attributes
}

func NewGroup(entry *ldap.Entry) FKITGroup {
	return FKITGroup{
		Description: SvEn{
			Sv: entry.GetAttributeValue("description"),
		},
		Email: entry.GetAttributeValue("mail"),
		Function: SvEn{
			Sv: entry.GetAttributeValue("function"),
		},
		Name:       entry.GetAttributeValue("cn"),
		PrettyName: entry.GetAttributeValue("displayname"),
		//SuperGroup: 	"",
		//Active: 		"",
		//GroupMembers: "",
	}
}

// CRUD Group =========================================

func (s *ServiceLDAP) AddGroup(group FKITGroup) error {
	return s.Connection.Add(&ldap.AddRequest{
		DN:         fmt.Sprintf("cn=%s,ou=%s,ou=fkit,%s", group.Name, group.SuperGroup.Name, s.GroupsConfig.BaseDN),
		Attributes: group.ToLdapAttrib(10, s.UsersConfig.BaseDN), //TODO: 10 is never used
	})
}

func (s *ServiceLDAP) GetGroups() ([]FKITGroup, error) {
	//TODO
	return nil, errors.New("Not yet implemented")
}

func (s *ServiceLDAP) GetGroup(groupName string) (FKITGroup, error) {
	//TODO
	return FKITGroup{}, errors.New("Not yet implemented")
}

func (s *ServiceLDAP) DeleteGroup(groupName string, superGroupName string) error {
	return s.Connection.Del(&ldap.DelRequest{
		DN: fmt.Sprintf("cn=%s,ou=%s,ou=fkit,%s", groupName, superGroupName, s.GroupsConfig.BaseDN),
	})
}
