package config

import (
	"github.com/spf13/cobra"
	"github.com/yuichi10/jiractl/interface/database"
	"go.uber.org/zap"
)

var board string
var sprint string

// NewSetTeam return set-team command
func NewSetTeam(ds database.IDataStore) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-url",
		Short: "set jira api info",
		Long:  "Able to set ipi nfo by using this command",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			zap.S().Info(args)
			// controller.SetJiraURL(args[0], jiraURL, ds)
			return nil
		},
	}
	cmd.Flags().StringVarP(&board, "board", "b", "", "set jira board")
	cmd.Flags().StringVarP(&sprint, "sprint", "s", "", "set jira sprint info")
	return cmd
}
