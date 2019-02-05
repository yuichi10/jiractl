package database

import (
	"fmt"

	"github.com/yuichi10/jiractl/entity"
	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v2"
)

// Credential is credential data in datastore
type Credential struct {
	DataStore      IDataStore `yaml:"-"`
	CredentialName string     `yaml:"-"`
	UserID         string     `yaml:"userID"`
	BasicAuth      string     `yaml:"basicAuth"`
}

// Context is context data in datastore
type Context struct {
	DataStore IDataStore `yaml:"-"`
	Name      string     `yaml:"-"`
	User      string     `yaml:"user"`
	JiraURL   string     `yaml:"url"`
}

// Config is config data in datastore
type Config struct {
	DataStore      IDataStore             `yaml:"-"`
	Credentials    map[string]*Credential `yaml:"credentials"`
	CurrentContext string                 `yaml:"currentContext"`
	JiraURLs       map[string]string      `yaml:"jiraURL"`
	Context        map[string]*Context    `yaml:"context"`
}

// NewConfig return empty config
func NewConfig(ds IDataStore) *Config {
	config := new(Config)
	config.DataStore = ds
	config.Credentials = make(map[string]*Credential)
	config.JiraURLs = make(map[string]string)
	config.Context = make(map[string]*Context)
	return config
}

// AddCredential add credential data to datastore
func (c *Config) AddCredential(credentName, userID, basicAuth string) entity.Credential {
	credent := &Credential{CredentialName: credentName, UserID: userID, BasicAuth: basicAuth}
	c.setAllConfigData()
	c.Credentials[credentName] = credent
	c.DataStore.Create(c)
	return entity.Credential{UserID: credent.UserID, BasicAuth: credent.BasicAuth}
}

// AddJiraURL add jira data to datastore
func (c *Config) AddJiraURL(name, url string) *entity.JiraURL {
	c.setAllConfigData()
	// TODO: すでにある名前のものは上書きされる
	c.JiraURLs[name] = url
	c.DataStore.Create(c)
	// TODO: 実際に帰ってきた値を返す必要がある
	return &entity.JiraURL{Name: name, URL: url}
}

// AddContext add context data to datastore
func (c *Config) AddContext(name, user, jiraURL string) *entity.Context {
	context := &Context{Name: name, User: user, JiraURL: jiraURL}
	c.setAllConfigData()
	c.Context[name] = context
	c.DataStore.Create(c)
	return &entity.Context{Name: name, UserID: user, URL: jiraURL}
}

// AddCurrentContext add current context data to datastore
func (c *Config) AddCurrentContext(name string) *entity.Current {
	c.setAllConfigData()
	c.CurrentContext = name
	c.DataStore.Create(c)
	context := c.Context[c.CurrentContext]
	credential := c.Credentials[context.User]
	return &entity.Current{ContextName: c.CurrentContext, UserID: context.User, URL: context.JiraURL, BasicAuth: credential.BasicAuth}
}

func (c *Config) GetCurrentContext() (*entity.Current, error) {
	c.setAllConfigData()
	current := &entity.Current{}
	context, ok := c.Context[c.CurrentContext]
	if !ok {
		return nil, fmt.Errorf("no context")
	}
	current.ContextName = c.CurrentContext
	user, ok := c.Credentials[context.User]
	if !ok {
		return nil, fmt.Errorf("no user")
	}
	current.BasicAuth = user.BasicAuth
	current.UserID = user.UserID
	url, ok := c.JiraURLs[context.JiraURL]
	if !ok {
		return nil, fmt.Errorf("no url")
	}
	current.JiraURL = url
	return current, nil
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
