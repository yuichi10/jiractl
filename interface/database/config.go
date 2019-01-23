package database

import (
	"fmt"

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
	JiraURLs       map[string]string     `yaml:"jiraURL"`
}

func NewConfig() *Config {
	config := new(Config)
	config.Credentials = make(map[string]Credential)
	config.JiraURLs = make(map[string]string)
	return config
}

func NewCredential(ds IDataStore) usecase.ICredentialDataAccess {
	return &Credential{DataStore: ds}
}

func (c *Credential) AddCredential(credentName, userID, basicAuth string) entity.Credential {
	c.CredentialName = credentName
	c.UserID = userID
	c.BasicAuth = basicAuth
	conf := readAllConfig(c.DataStore)
	conf.Credentials[credentName] = *c
	c.DataStore.Create(conf)
	// TODO: 取ってきた値をきちんと変換して返すようにする
	return entity.Credential{UserID: c.UserID, BasicAuth: c.BasicAuth}
}

type JiraURL struct {
	DataStore IDataStore
	Name      string
	URL       string
}

func NewJiraURL(ds IDataStore) *JiraURL {
	return &JiraURL{DataStore: ds}
}

func (j *JiraURL) AddJiraURL(name, url string) *entity.JiraURL {
	j.URL = url
	j.Name = name
	c := readAllConfig(j.DataStore)
	// すでに同じ名前のものが存在した場合上書きになる
	c.JiraURLs[name] = url
	j.DataStore.Create(c)
	// TODO: 実際に帰ってきた値を返すようにする必要がある
	return &entity.JiraURL{Name: name, URL: url}
}

func readAllConfig(ds IDataStore) *Config {
	data, err := ds.Read("")
	if err != nil {
		zap.S().Errorf("failed to read datastore: %v", err)
		// TODO: エラーを表示するpresenterを作成する
		panic(err)
	}
	conf := NewConfig()
	fmt.Println("try to open file")
	fmt.Printf("%+v", string(data))
	fmt.Println(len(data))
	err = yaml.Unmarshal([]byte(data), conf)
	if err != nil {
		// TODO: エラーを表示するpresenterを作成する
		zap.S().Errorf("failed to unmarchal config yaml: %v", err)
		panic(err)
	}
	return conf
}
