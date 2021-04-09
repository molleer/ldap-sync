package services

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/ldap.v2"
)

type FKITSuperGroup struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	PrettyName string `json:"prettyName"`
	Type       string `json:"type"`
	Email      string `json:"email"`
}

// CRUD Super Group ==============================================

func (s *ServiceLDAP) AddSuperGroup(superGroup FKITSuperGroup) error {
	//Creates super group directory in ldap
	err := s.Connection.Add(&ldap.AddRequest{
		DN: fmt.Sprintf("ou=%s,ou=fkit,%s", superGroup.Name, s.GroupsConfig.BaseDN),
		Attributes: []ldap.Attribute{
			{Type: "ou", Vals: []string{superGroup.Name}},
			{Type: "objectclass", Vals: []string{"organizationalUnit", "top"}},
		},
	})

	if err != nil {
		return err
	}

	//Creates active group
	group := FKITGroup{
		Name:       superGroup.Name,
		PrettyName: superGroup.PrettyName,
		Email:      superGroup.Email,
		SuperGroup: superGroup,
		Description: SvEn{
			Sv: fmt.Sprintf("%s saker", superGroup.PrettyName),
		},
	}
	groupAttribs := group.ToLdapAttrib(10, "") //TODO: 10 is never used
	groupAttribs[5].Vals = []string{fmt.Sprintf("cn=digit,%s", s.DBConfig.BaseDN)}
	groupAttribs = RemoveEmpty(groupAttribs)

	//Adds active group to the super group directory
	err = s.Connection.Add(&ldap.AddRequest{
		DN:         fmt.Sprintf("cn=%s,ou=%s,ou=fkit,%s", group.Name, group.SuperGroup.Name, s.GroupsConfig.BaseDN),
		Attributes: groupAttribs,
	})

	//If failed to create active group, it will be removed
	if err != nil {
		s.DeleteSuperGroup(superGroup.Name)
		return err
	}

	return nil
}

func (s *ServiceLDAP) GetSuperGroups() ([]FKITSuperGroup, error) {
	res, err := s.Connection.Search(ldap.NewSearchRequest(
		s.GroupsConfig.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(ou=*)",
		[]string{"ou"},
		nil,
	))
	if err != nil {
		return nil, err
	}

	superGroups := make([]FKITSuperGroup, 0)

	for _, e := range res.Entries {
		name := SplitDN(e.DN)[1]
		group, err := s.GetSuperGroup(name)
		if err != nil {
			continue
		}
		superGroups = append(superGroups, group)
	}

	return superGroups, nil
}

func (s *ServiceLDAP) GetSuperGroup(superGroupName string) (FKITSuperGroup, error) {
	group, err := s.Connection.Search(ldap.NewSearchRequest(
		fmt.Sprintf("cn=%s,ou=%s,ou=fkit,%s", superGroupName, superGroupName, s.GroupsConfig.BaseDN),
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(cn=%s)", superGroupName),
		[]string{"mail", "type", "displayName"},
		nil,
	))

	if err != nil {
		return FKITSuperGroup{}, err
	}

	if group == nil || len(group.Entries) == 0 {
		return FKITSuperGroup{}, errors.New(fmt.Sprintf("Failed to fins super group: %s", superGroupName))
	}

	return FKITSuperGroup{
		Name:       superGroupName,
		PrettyName: group.Entries[0].GetAttributeValue("displayName"),
		Email:      group.Entries[0].GetAttributeValue("mail"),
		Type:       strings.ToUpper(group.Entries[0].GetAttributeValue("type")),
	}, err
}

func (s *ServiceLDAP) DeleteSuperGroup(superGroupName string) error {
	subGroups, err := s.Connection.Search(ldap.NewSearchRequest(
		fmt.Sprintf("ou=%s,ou=fkit,%s", superGroupName, s.GroupsConfig.BaseDN),
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, "(cn=*)",
		[]string{"cn"},
		nil,
	))

	if err != nil {
		return err
	}

	for _, group := range subGroups.Entries {
		s.Connection.Del(&ldap.DelRequest{
			DN: group.DN,
		})
	}

	return s.Connection.Del(&ldap.DelRequest{
		DN: fmt.Sprintf("ou=%s,ou=fkit,%s", superGroupName, s.GroupsConfig.BaseDN),
	})
}
