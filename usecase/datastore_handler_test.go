package usecase

import "github.com/yuichi10/jiractl/entity"

type credentialDataAccess struct {
}

func AddCredential(credentialName, userName, basicAuth string) entity.Credential {
	return entity.Credential{}
}
