package config

import (
	"fmt"
	"os"
	"path"

	"github.com/ghodss/yaml"
	"github.com/spf13/viper"

	homedir "github.com/mitchellh/go-homedir"
)

type Config struct {
	Context map[string]struct {
		UserID    string `yaml:"userID"`
		BasicAuth string `yaml:"basicAuth"`
	} `yaml:"context"`
	CurrentContext string `yaml:"currentContext"`
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
	viper.SetConfigType("yaml")
	configFile, err := filePath()
	if err != nil {
		return err
	}
	err = initFile(configFile)
	if err != nil {
		return err
	}
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config by viper %v: ", err)
	}
	return nil
}

func Preserve() error {
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
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

func initFile(configFile string) error {
	if _, err := os.Stat(configFile); err != nil {
		// os.OpenFile(configFile, os.O_WRONLY, 0666)
		f, err := os.Create(configFile)
		if err != nil {
			return fmt.Errorf("failed to crate config file at %v: %v", configFile, err)
		}
		defer f.Close()
	}
	return nil
}
