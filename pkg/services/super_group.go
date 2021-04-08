package services

import (
	"errors"
	"fmt"

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
	//TODO
	return nil, errors.New("Not yet implemented")
}

func (s *ServiceLDAP) GetSuperGroup(superGroupName string) (FKITSuperGroup, error) {
	//TODO
	return FKITSuperGroup{}, errors.New("Not yet implemented")
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
