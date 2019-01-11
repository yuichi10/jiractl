package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
	cmd.Flags().StringP("user", "u", "", "set login user name for jira")
	viper.BindPFlag("login_user", cmd.Flags().Lookup("user"))
	cmd.Flags().StringP("password", "p", "", "set login user password for jira")
	viper.BindPFlag("login_password", cmd.Flags().Lookup("password"))
	return cmd
}
