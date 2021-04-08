package services

import (
	"math/rand"

	"gopkg.in/ldap.v2"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func ToPartial(attribs []ldap.Attribute) []ldap.PartialAttribute {
	res := make([]ldap.PartialAttribute, len(attribs))
	for i := 0; i < len(attribs); i++ {
		res[i].Type = attribs[i].Type
		res[i].Vals = attribs[i].Vals
	}
	return res
}

func RemoveEmpty(attribs []ldap.Attribute) []ldap.Attribute {
	res := make([]ldap.Attribute, 0)
	for _, attr := range attribs {
		if len(attr.Vals) > 0 && attr.Vals[0] != "" {
			res = append(res, attr)
		}
	}
	return res
}
