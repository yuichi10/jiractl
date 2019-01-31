package usecase

import (
	"encoding/base64"
	"fmt"
)

// IAuthInput is user input data for credential
type IAuthInput interface {
	GetUserName() string
	GetPassword() string
	GetCredentialName() string
}

// IJiraURLInput is user input data for jira info
type IJiraURLInput interface {
	GetJiraURL() string
	GetName() string
}

// IContextInput is user input data for context info
type IContextInput interface {
	GetName() string
	GetJiraURLName() string
	GetUser() string
}

// ICurrentContextInput is input dat for current context data
type ICurrentContextInput interface {
	GetName() string
}

func basicToken(userID, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", userID, password)))
}

// SetCredential set credential info to datastore
func SetCredential(ia IAuthInput, cda ICredentialDataAccess) {
	cda.AddCredential(ia.GetCredentialName(), ia.GetUserName(), basicToken(ia.GetUserName(), ia.GetPassword()))
}

// SetJiraURL set url info to datastore
func SetJiraURL(input IJiraURLInput, da IJiraURLDataAccess) {
	da.AddJiraURL(input.GetName(), input.GetJiraURL())
}

// SetContext set context data to datastore
func SetContext(input IContextInput, da IContextDataAccess) {
	da.AddContext(input.GetName(), input.GetUser(), input.GetJiraURLName())
}

// SetCurrentContext set current context data to datastore
func SetCurrentContext(input ICurrentContextInput, da ICurrentContextDataAccess) {
	da.AddCurrentContext(input.GetName())
}
