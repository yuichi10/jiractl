package config

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewUseContextCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "use-context",
		Short: "use context",
		Long:  "Able to set using context by using this command",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			zap.S().Info(args)
			return nil
		},
	}
	return cmd
}
