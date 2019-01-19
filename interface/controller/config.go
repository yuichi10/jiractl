package controller

import "github.com/yuichi10/jiractl/usecase"

type Credential struct {
	userName string
	password string
	credName string
}

func (c Credential) GetUserName() string {
	return c.userName
}

func (c Credential) GetPassword() string {
	return c.password
}

func (c Credential) GetCredentialName() string {
	return c.credName
}

func SetCredentialController(credName, userName, password string) {
	cred := Credential{credName: credName, userName: userName, password: password}
	usecase.SetCredential(cred)
}
