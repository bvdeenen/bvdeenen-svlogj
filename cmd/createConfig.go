// Package cmd 
package cmd

import (
	"svlogj/pkg/config"

	"github.com/spf13/cobra"
)

// createConfigCmd represents the createConfig command
var createConfigCmd = &cobra.Command{
	Use:   "create-config",
	Short: "create a configuration file ~/.config/svlogj.json",
	Long: ` `,
	Run: func(cmd *cobra.Command, args []string) {
		config.ParseAndStoreConfig()
	},
}

func init() {
	rootCmd.AddCommand(createConfigCmd)
}
