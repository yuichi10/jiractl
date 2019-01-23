package infrastructure

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/yuichi10/jiractl/interface/database"
)

type YamlHandler struct {
	filename string
}

func NewYamlHandelr() (database.IDataStore, error) {
	f, err := filePath()
	if err != nil {
		return nil, err
	}
	err = initConfigFile(f)
	if err != nil {
		return nil, err
	}
	yh := &YamlHandler{filename: f}
	return yh, nil
}

func filePath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %v", err)
	}
	return path.Join(home, ".jiractl.yaml"), nil
}

func initConfigFile(configFile string) error {
	if _, err := os.Stat(configFile); err != nil {
		// os.OpenFile(configFile, os.O_WRONLY, 0666)
		_, err := os.Create(configFile)
		if err != nil {
			return fmt.Errorf("failed to crate config file at %v: %v", configFile, err)
		}
		return nil
	}
	return nil
}

func (y YamlHandler) Create(data interface{}) (string, error) {
	fmt.Println("every data are store to yaml")
	b, err := yaml.Marshal(data)
	if err != nil {
		// TODO: エラーを表示するpresenterを利用する
		zap.S().Errorf("failed to marshal data for create %v", err)
		panic(err)
	}
	err = ioutil.WriteFile(y.filename, b, 0666)
	if err != nil {
		// TODO: エラーを表示するpresenterを利用する
		zap.S().Errorf("failed to create data for create %v", err)
		panic(err)
	}
	return string(b), nil
}

func (y YamlHandler) Update(data interface{}) (string, error) {
	fmt.Println("update yaml data")
	return "", nil

}

func (y YamlHandler) Read(data interface{}) (string, error) {
	b, err := ioutil.ReadFile(y.filename)
	if err != nil {
		return "", fmt.Errorf("failed to read config file: %v", err)
	}
	return string(b), nil
}

func (y YamlHandler) Close() {
	fmt.Println("close")
}
