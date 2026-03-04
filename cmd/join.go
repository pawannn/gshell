package cmd

import (
	"fmt"

	"github.com/pawannn/gshell/internal/transport"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:   "join [address:port]",
	Short: "Join a gshell session",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]
		fmt.Println("Joining gshell session at", address)
		if err := transport.Connect(address); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(joinCmd)
}
