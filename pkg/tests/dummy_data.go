package main

import "github.com/molleer/ldap-sync/pkg/services"

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

var dummySuperGroup = services.FKITSuperGroup{
	Name:       "lolkit",
	Type:       "COMMITTEE",
	Email:      "lolkit@chalmers.it",
	PrettyName: "lolKIT",
}

var dummyGroup = services.FKITGroup{
	Name: "lolkit01",
	Description: services.SvEn{
		Sv: "LOL saker",
	},
	PrettyName: "lolKIT 01/02",
	Function: services.SvEn{
		Sv: "Att inte implodera",
	},
	Email: "lolkit01@chalmers.it",
	GroupMembers: []services.FKITUser{
		{
			Cid: dummyUser.Cid,
			Post: services.Post{
				Sv: "Hubbenansvarig",
			},
		},
	},
	SuperGroup: dummySuperGroup,
}
