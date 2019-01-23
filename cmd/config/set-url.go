package config

import (
	"github.com/spf13/cobra"
	"github.com/yuichi10/jiractl/interface/controller"
	"github.com/yuichi10/jiractl/interface/database"
	"go.uber.org/zap"
)

var jiraURL string

func NewSetURL(ds database.IDataStore) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-url",
		Short: "set jira api info",
		Long:  "Able to set ipi nfo by using this command",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			zap.S().Info(args)
			controller.SetJiraURL(args[0], jiraURL, ds)
			return nil
		},
	}
	cmd.Flags().StringVarP(&jiraURL, "url", "u", "", "set jira url")
	return cmd
}
