package cmd

import (
	"github.com/pawannn/gshell/internal/app"
	"github.com/spf13/cobra"
)

var port string

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Start gshell in share mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Share(port, "default-session")
	},
}

func init() {
	shareCmd.Flags().StringVarP(&port, "port", "p", "9000", "Port to listen on")
	rootCmd.AddCommand(shareCmd)
}
