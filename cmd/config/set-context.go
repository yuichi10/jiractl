package config

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewSetContextCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-url",
		Short: "set jira api info",
		Long:  "Able to set ipi nfo by using this command",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			zap.S().Info(args)
			return nil
		},
	}
	return cmd
}
