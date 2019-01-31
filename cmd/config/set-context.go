package config

import (
	"github.com/spf13/cobra"
	"github.com/yuichi10/jiractl/interface/controller"
	"github.com/yuichi10/jiractl/interface/database"
	"go.uber.org/zap"
)

var (
	user string
	url  string
)

// NewSetContextCmd return set-context command
func NewSetContextCmd(ds database.IDataStore) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-context",
		Short: "set context info",
		Long:  "Able to set context info which include credential and jira url info",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			zap.S().Info(args)
			controller.SetContext(args[0], user, url, ds)
			return nil
		},
	}
	cmd.Flags().StringVar(&user, "user", "", "set credential")
	cmd.Flags().StringVar(&url, "url", "", "set url name (set-url cmd)")
	return cmd
}
