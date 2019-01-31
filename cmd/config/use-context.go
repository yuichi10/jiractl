package config

import (
	"github.com/spf13/cobra"
	"github.com/yuichi10/jiractl/interface/controller"
	"github.com/yuichi10/jiractl/interface/database"
	"go.uber.org/zap"
)

func NewUseContextCmd(ds database.IDataStore) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "use-context",
		Short: "use context",
		Long:  "Able to set using context by using this command",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			zap.S().Info(args)
			controller.SetCurrentContext(args[0], ds)
			return nil
		},
	}
	return cmd
}
