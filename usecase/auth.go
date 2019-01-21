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

func basicToken(userID, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", userID, password)))
}

func SetCredential(ia IAuthInput, cda ICredentialDataAccess) {
	cda.AddCredential(ia.GetCredentialName(), ia.GetUserName(), basicToken(ia.GetUserName(), ia.GetPassword()))
}
