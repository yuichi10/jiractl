package sprint

import (
	"github.com/spf13/cobra"
	"github.com/yuichi10/jiractl/interface/api"
	"github.com/yuichi10/jiractl/interface/database"
	"github.com/yuichi10/jiractl/interface/presenter"
)

var format string

func NewSprintCmd(iapi api.IAPI, ds database.IDataStore, viewer presenter.Viewer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sprint",
		Short: "do sprint related thing",
		Long:  "get sprint infomation",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&format, "format", "markdown", "add format for output")
	cmd.AddCommand(NewIssueCmd(format, iapi, ds, viewer))
	return cmd
}
