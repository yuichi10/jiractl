package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
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
	return cmd
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		os.Exit(1)
	}
	configFile = path.Join(home, ".jiractl.yaml")
}

func Execute() {
	cmd := NewRootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
