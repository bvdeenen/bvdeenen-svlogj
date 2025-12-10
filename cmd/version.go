// Package cmd 
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev-build"

// createConfigCmd represents the createConfig command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "returns the version",
	Long: ` `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
