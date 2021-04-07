package services

import "errors"

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

// CRUD Group =========================================

func (s *ServiceLDAP) AddGroup(group FKITGroup) error {
	//TODO
	return errors.New("Not yet implemented")
}

func (s *ServiceLDAP) GetGroups() ([]FKITGroup, error) {
	//TODO
	return nil, errors.New("Not yet implemented")
}

func (s *ServiceLDAP) GetGroup(groupName string) (FKITGroup, error) {
	//TODO
	return FKITGroup{}, errors.New("Not yet implemented")
}

func (s *ServiceLDAP) DeleteGroup(name string) error {
	//TODO
	return errors.New("Not yet implemented")
}
