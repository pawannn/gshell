package cmd

import (
	"github.com/pawannn/gshell/internal/app"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:   "join [address:port]",
	Short: "Join a gshell session",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := args[0]
		return app.Join(token)
	},
}

func init() {
	rootCmd.AddCommand(joinCmd)
}
