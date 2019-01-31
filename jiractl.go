package main

import (
	"fmt"
	"os"

	"github.com/yuichi10/jiractl/cmd"
	"github.com/yuichi10/jiractl/infrastructure"
	_ "github.com/yuichi10/jiractl/logger"
)

func main() {
	ds, err := infrastructure.NewYamlHandler()
	if err != nil {
		fmt.Println("failed to open config file")
		os.Exit(1)
	}
	defer ds.Close()
	cmd.Execute(ds)
}
