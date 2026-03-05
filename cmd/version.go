package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print gshell version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gshell version:", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
