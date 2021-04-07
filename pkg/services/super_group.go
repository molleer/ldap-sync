package services

import "errors"

type FKITSuperGroup struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	PrettyName string `json:"prettyName"`
	Type       string `json:"type"`
	Email      string `json:"email"`
}

// CRUD Super Group ==============================================

func (s *ServiceLDAP) AddSuperGroup(group FKITSuperGroup) error {
	//TODO
	return errors.New("Not yet implemented")
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
	//TODO
	return errors.New("Not yet implemented")
}
