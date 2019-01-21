package config

import (
	"github.com/spf13/cobra"
	"github.com/yuichi10/jiractl/interface/database"
)

func NewConfigCmd(ds database.IDataStore) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "treat config settings",
		Long:  "Able to change config settings by using this command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(NewSetContextCmd(ds))
	cmd.AddCommand(NewSetURL(ds))

	return cmd
}
