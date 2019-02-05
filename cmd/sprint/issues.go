package sprint

import (
	"github.com/spf13/cobra"
	"github.com/yuichi10/jiractl/interface/api"
	"github.com/yuichi10/jiractl/interface/controller"
	"github.com/yuichi10/jiractl/interface/database"
)

var board string
var sprint string

func NewIssueCmd(iapi api.IAPI, ds database.IDataStore) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issues",
		Short: "list sprint related issues",
		Long:  "get sprint issues",
		RunE: func(cmd *cobra.Command, args []string) error {
			controller.GetSprintIssue(board, sprint, iapi, ds)
			return nil
		},
	}
	cmd.Flags().StringVarP(&board, "board", "b", "", "set board name")
	cmd.Flags().StringVarP(&sprint, "sprint", "s", "", "set sprint name")
	return cmd
}
