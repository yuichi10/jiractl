package config

import (
	"github.com/yuichi10/jiractl/interface/controller"
	"github.com/yuichi10/jiractl/interface/database"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var username string
var password string

func NewSetContextCmd(ds database.IDataStore) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-credentials",
		Short: "set credential info",
		Long:  "Able to set credentials info by using this command",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			zap.S().Info(args)
			// setLoginInfo(args[0])
			// config.Preserve()
			controller.SetCredentialController(args[0], username, password, ds)
			return nil
		},
	}
	cmd.Flags().StringVarP(&username, "user", "u", "", "set login user name for jira")
	cmd.Flags().StringVarP(&password, "password", "p", "", "set login user password for jira")
	return cmd
}

// func setLoginInfo(credentialID string) error {
// 	if username == "" {
// 		fmt.Print("login user: ")
// 		stdin := bufio.NewScanner(os.Stdin)
// 		stdin.Scan()
// 		username = stdin.Text()
// 	}
// 	if password == "" {
// 		fmt.Print("login password: ")
// 		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
// 		if err != nil {
// 			return fmt.Errorf("failed to get password: %v", err)
// 		}
// 		password = string(bytePassword)
// 	}
// 	basic := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
// 	config.AddCredential(credentialID, username, basic)

// 	return nil
// }
