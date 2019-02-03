package sprint

import (
	"github.com/spf13/cobra"
)

func NewSprintCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sprint",
		Short: "do sprint related thing",
		Long:  "get sprint infomation",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}
