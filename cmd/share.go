package cmd

import (
	"fmt"

	"github.com/pawannn/gshell/internal/transport"
	"github.com/spf13/cobra"
)

var port string

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Start gshell in share mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting gshell in share mode on port", port)
		if err := transport.StartListener(port); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	shareCmd.Flags().StringVarP(&port, "port", "p", "9000", "Port to listen on")
	rootCmd.AddCommand(shareCmd)
}
