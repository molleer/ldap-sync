package main

import (
	"testing"

	"github.com/molleer/ldap-sync/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestSplitDN(t *testing.T) {
	assert.Equal(t,
		[]string{"uid", "hi", "cn", "group", "ou", "root", "dn", "localhost", "cn", "com"},
		services.SplitDN("uid=hi,cn=group,ou=root,dn=localhost,cn=com"),
		"Failed to split  DN")
}
