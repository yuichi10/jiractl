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
	file *os.File
}

func NewYamlHandelr() (database.IDataStore, error) {
	configFile, err := filePath()
	if err != nil {
		return nil, err
	}
	f, err := openConfig(configFile)
	if err != nil {
		return nil, err
	}
	yh := &YamlHandler{file: f}
	return yh, nil
}

func filePath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %v", err)
	}
	return path.Join(home, ".jiractl.yaml"), nil
}

func openConfig(configFile string) (*os.File, error) {
	if _, err := os.Stat(configFile); err != nil {
		// os.OpenFile(configFile, os.O_WRONLY, 0666)
		f, err := os.Create(configFile)
		if err != nil {
			return nil, fmt.Errorf("failed to crate config file at %v: %v", configFile, err)
		}

		return f, nil
	}
	return os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0666)
}

func (y YamlHandler) Create(data interface{}) (string, error) {
	fmt.Println("every data are store to yaml")
	fmt.Println("create data %+v", data)
	b, err := yaml.Marshal(data)
	if err != nil {
		// TODO: エラーを表示するpresenterを利用する
		zap.S().Errorf("failed to marshal data for create %v", err)
		panic(err)
	}
	// make file empty
	y.file.Truncate(0)
	_, err = y.file.Write(b)
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
	b, err := ioutil.ReadAll(y.file)
	if err != nil {
		return "", fmt.Errorf("failed to read config file: %v", err)
	}
	return string(b), nil
}

func (y YamlHandler) Close() {
	y.file.Close()
}
