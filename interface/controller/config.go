package controller

import (
	"github.com/yuichi10/jiractl/interface/database"
	"github.com/yuichi10/jiractl/usecase"
)

// CredentialInput is user input for credential
type CredentialInput struct {
	userName string
	password string
	credName string
}

// GetUserName return user name
func (c CredentialInput) GetUserName() string {
	return c.userName
}

// GetPassword return password
func (c CredentialInput) GetPassword() string {
	return c.password
}

// GetCredentialName return credential name
func (c CredentialInput) GetCredentialName() string {
	return c.credName
}

// SetCredentialController is controller for set credential
func SetCredentialController(credName, userName, password string, ds database.IDataStore) {
	c := database.NewConfig(ds)
	input := CredentialInput{userName: userName, password: password, credName: credName}
	usecase.SetCredential(input, c)
}

// JiraURLInput is user input for jira
type JiraURLInput struct {
	jiraURL string
	name    string
}

// GetJiraURL return jira url
func (j JiraURLInput) GetJiraURL() string {
	return j.jiraURL
}

// GetName return name
func (j JiraURLInput) GetName() string {
	return j.name
}

// SetJiraURL set jira url
func SetJiraURL(name, url string, ds database.IDataStore) {
	c := database.NewConfig(ds)
	input := &JiraURLInput{jiraURL: url, name: name}
	usecase.SetJiraURL(input, c)
}

// ContextInput is user input for context
type ContextInput struct {
	Name    string
	User    string
	jiraURL string
}

// GetName return context name
func (c ContextInput) GetName() string {
	return c.Name
}

// GetUser return user name
func (c ContextInput) GetUser() string {
	return c.User
}

// GetJiraURLName return jira url name
func (c ContextInput) GetJiraURLName() string {
	return c.jiraURL
}

// SetContext is controller for setting context
func SetContext(contextName, user, url string, ds database.IDataStore) {
	c := database.NewConfig(ds)
	input := &ContextInput{Name: contextName, User: user, jiraURL: url}
	usecase.SetContext(input, c)
}

// CurrentContextInput is user input for current context
type CurrentContextInput struct {
	Name string
}

// GetName return context name
func (c CurrentContextInput) GetName() string {
	return c.Name
}

// SetCurrentContext is controller for setting current context
func SetCurrentContext(contextName string, ds database.IDataStore) {
	c := database.NewConfig(ds)
	input := &ContextInput{Name: contextName}
	usecase.SetCurrentContext(input, c)
}
