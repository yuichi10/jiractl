package cmd

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/spf13/cobra"
	"github.com/yuichi10/jiractl/config"
)

var configFile string

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Do Stuff Here
			return nil
		},
	}
	cobra.OnInitialize(initConfig)
	return cmd
}

func initConfig() {
	err := config.LoadConfig()
	if err != nil {
		zap.S().Error(err)
		os.Exit(1)
	}
}

func Execute() {
	cmd := NewRootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
