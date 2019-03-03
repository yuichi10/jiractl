package usecase

import "github.com/yuichi10/jiractl/entity"

// ICredentialDataAccess is for add credential data
type ICredentialDataAccess interface {
	AddCredential(credentialName, userName, basicAuth string) entity.Credential
}

// IJiraURLDataAccess is for add jira info
type IJiraURLDataAccess interface {
	AddJiraURL(name, url string) *entity.JiraURL
}

// IContextDataAccess is for add context info
type IContextDataAccess interface {
	AddContext(context, user, url string) *entity.Context
}

// ICurrentContextDataAccess is for add current context info
type ICurrentContextDataAccess interface {
	AddCurrentContext(name string) *entity.Current
	GetCurrentContext() (*entity.Current, error)
}

type ITeamDataAccess interface {
	AddTeam(board, sprint string) *entity.JiraTeam
}
