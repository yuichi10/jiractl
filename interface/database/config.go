package database

import (
	"github.com/yuichi10/jiractl/entity"
	"github.com/yuichi10/jiractl/usecase"
)

type Credential struct {
	DataStore      IDataStore
	CredentialName string
	UserID         string `yaml:"userID"`
	BasicAuth      string `yaml:"basicAuth"`
}

type Config struct {
	Credentials    map[string]Credential `yaml:"credentials"`
	CurrentContext string                `yaml:"currentContext"`
}

func NewCredential(ds IDataStore) usecase.IConfigDataAcess {
	return &Credential{DataStore: ds}
}

func (c *Credential) AddCredential(credentName, userID, basicAuth string) entity.Credential {
	c.CredentialName = credentName
	c.UserID = userID
	c.BasicAuth = basicAuth
	c.DataStore.Create(c)
	return entity.Credential{UserID: c.UserID, BasicAuth: c.BasicAuth}
}
