package database

import (
	"github.com/yuichi10/jiractl/entity"
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

func NewConfig(ds IDataStore) *Config {
	config := new(Config)
	config.DataStore = ds
	config.Credentials = make(map[string]*Credential)
	config.JiraURLs = make(map[string]string)
	config.Context = make(map[string]*Context)
	return config
}

func (c *Config) AddCredential(credentName, userID, basicAuth string) entity.Credential {
	credent := &Credential{CredentialName: credentName, UserID: userID, BasicAuth: basicAuth}
	c.setAllConfigData()
	c.Credentials[credentName] = credent
	c.DataStore.Create(c)
	return entity.Credential{UserID: credent.UserID, BasicAuth: credent.BasicAuth}
}

func (c *Config) AddJiraURL(name, url string) *entity.JiraURL {
	c.setAllConfigData()
	// TODO: すでにある名前のものは上書きされる
	c.JiraURLs[name] = url
	c.DataStore.Create(c)
	// TODO: 実際に帰ってきた値を返す必要がある
	return &entity.JiraURL{Name: name, URL: url}
}

func (c *Config) AddContext(name, user, jiraURL string) *entity.Context {
	context := &Context{Name: name, User: user, JiraUrl: jiraURL}
	c.setAllConfigData()
	c.Context[name] = context
	c.DataStore.Create(c)
	return &entity.Context{Name: name, UserID: user, URL: jiraURL}
}

func (c *Config) AddCurrentContext(name string) *entity.Current {
	c.setAllConfigData()
	c.CurrentContext = name
	c.DataStore.Create(c)
	context := c.Context[c.CurrentContext]
	credential := c.Credentials[context.User]
	return &entity.Current{ContextName: c.CurrentContext, UserID: context.User, URL: context.JiraUrl, BasicAuth: credential.BasicAuth}
}

func (c *Config) setAllConfigData() {
	data, err := c.DataStore.Read("")
	if err != nil {
		zap.S().Errorf("failed to read datastore: %v", err)
		// TODO: エラーを表示するpresenterを作成する
		panic(err)
	}
	err = yaml.Unmarshal([]byte(data), c)
	if err != nil {
		// TODO: エラーを表示するpresenterを作成する
		zap.S().Errorf("failed to unmarchal config yaml: %v", err)
		panic(err)
	}
}
