package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type authInput struct {
	userName       string
	password       string
	credentialName string
}

func (a authInput) GetUserName() string {
	return a.userName
}

func (a authInput) GetPassword() string {
	return a.password
}

func (a authInput) GetCredentialName() string {
	return a.credentialName
}

func TestBasicAuth(t *testing.T) {
	userID := "userid"
	password := "password"
	expect := "dXNlcmlkOnBhc3N3b3Jk"
	b := basicToken(userID, password)
	assert.Equal(t, b, expect, "they should be equal")
}
