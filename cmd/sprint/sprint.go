package sprint

import (
	"github.com/spf13/cobra"
	"github.com/yuichi10/jiractl/interface/api"
	"github.com/yuichi10/jiractl/interface/database"
)

func NewSprintCmd(iapi api.IAPI, ds database.IDataStore) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sprint",
		Short: "do sprint related thing",
		Long:  "get sprint infomation",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(NewIssueCmd(iapi, ds))
	return cmd
}
