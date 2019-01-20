package controller

import (
	"github.com/yuichi10/jiractl/interface/database"
	"github.com/yuichi10/jiractl/usecase"
)

type CredentialInput struct {
	userName string
	password string
	credName string
}

func (c CredentialInput) GetUserName() string {
	return c.userName
}

func (c CredentialInput) GetPassword() string {
	return c.password
}

func (c CredentialInput) GetCredentialName() string {
	return c.credName
}

func SetCredentialController(credName, userName, password string, ds database.IDataStore) {
	credData := database.NewCredential(ds)
	input := CredentialInput{userName: userName, password: password, credName: credName}
	usecase.SetCredential(input, credData)
}