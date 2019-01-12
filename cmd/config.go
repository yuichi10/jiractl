package cmd

import (
	"github.com/spf13/cobra"
)

type Config struct {
	Context []struct {
		Name      string `yaml:"name"`
		UserID    string `yaml:"userID"`
		BasicAuth string `yaml:"basicAuth"`
	} `yaml:"context"`
	CurrentContext string `yaml:"currentContext"`
}

func NewConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "treat config settings",
		Long:  "Able to change config settings by using this command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(NewSetContextCmd())

	return cmd
}
