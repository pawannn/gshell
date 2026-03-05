package cmd

import (
	"fmt"
	"os"

	"github.com/pawannn/gshell/internal/app"
	"github.com/spf13/cobra"
)

var port string

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Start gshell in share mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		user := os.Getenv("USER")
		host, _ := os.Hostname()

		session := fmt.Sprintf("%s-%s", user, host)
		return app.Share(port, session, password)
	},
}

func init() {
	shareCmd.Flags().StringVarP(&port, "port", "p", "1337", "Port to listen on")
	shareCmd.Flags().StringVarP(&password, "password", "s", "", "Session password")
	rootCmd.AddCommand(shareCmd)
}
