package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	yaml "gopkg.in/yaml.v2"

	homedir "github.com/mitchellh/go-homedir"
)

type Credential struct {
	UserID    string `yaml:"userID"`
	BasicAuth string `yaml:"basicAuth"`
}

type Config struct {
	Credentials    map[string]Credential `yaml:"credentials"`
	CurrentContext string                `yaml:"currentContext"`
}

var conf *Config

func AddCredential(credentName, userID, basicAuth string) {
	if conf.Credentials == nil {
		fmt.Println("this is null")
		os.Exit(1)
	}
	conf.Credentials[credentName] = Credential{UserID: userID, BasicAuth: basicAuth}
}

func CurrentContext() string {
	return "currentContext"
}

func UserID(credentialID string) string {
	return fmt.Sprintf("credentials.%s.userID", credentialID)
}

func BasicAuth(credentialID string) string {
	return fmt.Sprintf("credentials.%s.basicAuth", credentialID)
}

// LoadConfig load config file. If there is no config file
// this create new config file
func LoadConfig() error {
	configFile, err := filePath()
	if err != nil {
		return err
	}
	conf, err = initConfig(configFile)
	if err != nil {
		return err
	}
	return nil
}

func Preserve() error {
	bs, err := yaml.Marshal(conf)
	if err != nil {
		return fmt.Errorf("failed to preserve settings: %v", err)
	}
	fpath, err := filePath()
	if err != nil {
		return fmt.Errorf("failed to get config file path: %v", err)
	}
	f, err := os.OpenFile(fpath, os.O_WRONLY, 666)
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)
	}
	f.Write(bs)
	return nil
}

func filePath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %v", err)
	}
	return path.Join(home, ".jiractl.yaml"), nil
}

func newConfig() *Config {
	c := new(Config)
	c.Credentials = make(map[string]Credential)
	return c
}

func initConfig(configFile string) (*Config, error) {
	if _, err := os.Stat(configFile); err != nil {
		// os.OpenFile(configFile, os.O_WRONLY, 0666)
		f, err := os.Create(configFile)
		if err != nil {
			return nil, fmt.Errorf("failed to crate config file at %v: %v", configFile, err)
		}
		defer f.Close()
		return newConfig(), nil
	}
	return readConfigFile(configFile)
}

func readConfigFile(filepath string) (*Config, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}
	c := newConfig()
	yaml.Unmarshal(b, c)
	return c, nil
}
