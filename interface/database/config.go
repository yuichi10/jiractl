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

type Context struct {
	DataStore IDataStore `yaml:"-"`
	Name      string     `yaml:"-"`
	User      string     `yaml:"user"`
	JiraUrl   string     `yaml:"url"`
}

type Config struct {
	DataStore      IDataStore             `yaml:"-"`
	Credentials    map[string]*Credential `yaml:"credentials"`
	CurrentContext string                 `yaml:"currentContext"`
	JiraURLs       map[string]string      `yaml:"jiraURL"`
	Context        map[string]*Context    `yaml:"context"`
}

// TODO: NewConfigをdatasotreだけを持つConfigを、newConfigを作ってそっちがデータ全てを持つようにする
func NewConfig() *Config {
	config := new(Config)
	config.Credentials = make(map[string]*Credential)
	config.JiraURLs = make(map[string]string)
	config.Context = make(map[string]*Context)
	return config
}

func NewCredential(ds IDataStore) usecase.ICredentialDataAccess {
	return &Credential{DataStore: ds}
}

func (c *Config) AddCredential(credentName, userID, basicAuth string) entity.Credential {
	credent := &Credential{CredentialName: credentName, UserID: userID, BasicAuth: basicAuth}
	c = readAllConfig(c.DataStore)
	c.Credentials[credentName] = credent
	c.DataStore.Create(c)
	return entity.Credential{UserID: credent.UserID, BasicAuth: credent.BasicAuth}
}

func (c *Credential) AddCredential(credentName, userID, basicAuth string) entity.Credential {
	c.CredentialName = credentName
	c.UserID = userID
	c.BasicAuth = basicAuth
	conf := readAllConfig(c.DataStore)
	conf.Credentials[credentName] = c
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

func (c *Config) AddJiraURL(name, url string) *entity.JiraURL {
	c = readAllConfig(c.DataStore)
	c.JiraURLs[name]
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

func NewContext(ds IDataStore) *Context {
	return &Context{DataStore: ds}
}

func (c *Context) AddContext(name, user, jiraURL string) *entity.Context {
	c.Name = name
	c.User = user
	c.JiraUrl = jiraURL
	conf := readAllConfig(c.DataStore)
	conf.Context[c.Name] = c
	c.DataStore.Create(conf)
	return &entity.Context{Name: name, UserID: user, URL: jiraURL}
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
