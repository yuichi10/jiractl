package usecase

import "github.com/yuichi10/jiractl/entity"

type ICredentialDataAccess interface {
	AddCredential(credentialName, userName, basicAuth string) entity.Credential
}
