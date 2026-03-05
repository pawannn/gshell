package cmd

import (
	"github.com/pawannn/gshell/internal/app"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:   "join [token]",
	Short: "Join a gshell session",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := args[0]
		return app.Join(token, password)
	},
}

func init() {
	joinCmd.Flags().StringVarP(&password, "password", "s", "", "session password")
	rootCmd.AddCommand(joinCmd)
}
