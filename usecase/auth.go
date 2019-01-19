package usecase

import (
	"encoding/base64"
	"fmt"

	"github.com/yuichi10/jiractl/entity"
)

type IAuth interface {
	GetCredentialName() string
	GetUserName() string
	GetPassword() string
}

func basicToken(cred entity.Credential) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", cred.UserID, cred.Password)))
}

func SetCredential(ia IAuth) {
	fmt.Println(ia.GetCredentialName())
	fmt.Println(ia.GetUserName())
	fmt.Println(ia.GetPassword())
}
