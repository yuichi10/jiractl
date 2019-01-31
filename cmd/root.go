package cmd

import (
	"fmt"
	"os"

	"github.com/yuichi10/jiractl/interface/database"

	"github.com/spf13/cobra"
	configCmd "github.com/yuichi10/jiractl/cmd/config"
	// "github.com/yuichi10/jiractl/config"
)

var configFile string

// NewRootCmd return root comand
func NewRootCmd(ds database.IDataStore) *cobra.Command {
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
	cmd.AddCommand(configCmd.NewConfigCmd(ds))
	return cmd
}

func initConfig() {
	// err := config.LoadConfig()
	// if err != nil {
	// 	zap.S().Error(err)
	// 	os.Exit(1)
	// }
}

// Execute exec jiractl command
func Execute(ds database.IDataStore) {
	cmd := NewRootCmd(ds)
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
