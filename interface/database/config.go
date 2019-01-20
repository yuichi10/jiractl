package database

import (
	"github.com/yuichi10/jiractl/entity"
	"github.com/yuichi10/jiractl/usecase"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type Credential struct {
	DataStore      IDataStore `yaml:"-"`
	CredentialName string     `yaml:"-"`
	UserID         string     `yaml:"userID"`
	BasicAuth      string     `yaml:"basicAuth"`
}

type Config struct {
	Credentials    map[string]Credential `yaml:"credentials"`
	CurrentContext string                `yaml:"currentContext"`
}

func NewConfig() *Config {
	config := new(Config)
	config.Credentials = make(map[string]Credential)
	return config
}

func NewCredential(ds IDataStore) usecase.IConfigDataAcess {
	return &Credential{DataStore: ds}
}

func (c *Credential) AddCredential(credentName, userID, basicAuth string) entity.Credential {
	c.CredentialName = credentName
	c.UserID = userID
	c.BasicAuth = basicAuth
	data, err := c.DataStore.Read("")
	if err != nil {
		zap.S().Errorf("failed to read datastore: %v", err)
		// TODO: エラーを表示するpresenterを作成する
		panic(err)
	}
	conf := NewConfig()
	err = yaml.Unmarshal([]byte(data), conf)
	if err != nil {
		// TODO: エラーを表示するpresenterを作成する
		zap.S().Errorf("failed to unmarchal config yaml: %v", err)
		panic(err)
	}
	conf.Credentials[credentName] = *c
	c.DataStore.Create(conf)
	// TODO: 取ってきた値をきちんと変換して返すようにする
	return entity.Credential{UserID: c.UserID, BasicAuth: c.BasicAuth}
}
