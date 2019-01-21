package usecase

import (
	"encoding/base64"
	"fmt"
)

type IAuthInput interface {
	GetUserName() string
	GetPassword() string
	GetCredentialName() string
}

type IJiraURLInput interface {
	GetJiraURL() string
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
