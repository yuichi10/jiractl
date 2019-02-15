package main

import (
	"fmt"
	"os"

	"github.com/yuichi10/jiractl/cmd"
	"github.com/yuichi10/jiractl/infrastructure"
	"github.com/yuichi10/jiractl/infrastructure/view"
	_ "github.com/yuichi10/jiractl/logger"
)

func main() {
	ds, err := infrastructure.NewYamlHandler()
	if err != nil {
		fmt.Println("failed to open config file")
		os.Exit(1)
	}
	apiClient := infrastructure.NewJiraAPIClient()

	viewer := view.NewStdoutViewer()
	defer ds.Close()
	cmd.Execute(apiClient, ds, viewer)
}
