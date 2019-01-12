package cmd

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var username string
var password string

func NewSetContextCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-context",
		Short: "set context info",
		Long:  "Able to set context info by using this command",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			zap.S().Info(args)
			loginInfo(args[0])
			return nil
		},
	}
	cmd.Flags().StringVarP(&username, "user", "u", "", "set login user name for jira")
	cmd.Flags().StringVarP(&password, "password", "p", "", "set login user password for jira")
	return cmd
}

func loginInfo(context string) error {
	if username == "" {
		fmt.Print("login user: ")
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		username = stdin.Text()
	}
	if password == "" {
		fmt.Print("login password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return fmt.Errorf("failed to get password: %v", err)
		}
		password = string(bytePassword)
	}
	return nil
}
