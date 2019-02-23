package sprint

import (
	"github.com/spf13/cobra"
	"github.com/yuichi10/jiractl/interface/api"
	"github.com/yuichi10/jiractl/interface/controller"
	"github.com/yuichi10/jiractl/interface/database"
	"github.com/yuichi10/jiractl/interface/presenter"
)

var board string
var sprint string
var detail bool

func NewIssueCmd(format string, iapi api.IAPI, ds database.IDataStore, viewer presenter.Viewer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issues",
		Short: "list sprint related issues",
		Long:  "get sprint issues",
		RunE: func(cmd *cobra.Command, args []string) error {
			controller.GetSprintIssue(board, sprint, format, iapi, ds, viewer)
			return nil
		},
	}
	cmd.Flags().BoolVar(&detail, "detail", false, "show detail data")
	cmd.Flags().StringVarP(&board, "board", "b", "", "set board name")
	cmd.Flags().StringVarP(&sprint, "sprint", "s", "", "set sprint name. If you do not set, you got active sprint issues")
	return cmd
}
