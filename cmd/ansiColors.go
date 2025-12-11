package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createConfigCmd represents the createConfig command
var ansiColorsCmd = &cobra.Command{
	Use:   "ansi-colors",
	Short: "shows ansi colors for use with the --ansi-color flag",
	Long: `Shows a table of ansi colors for use with the --ansi-color flag.

You can pick the combination of foreground and background colors from the table below
For instance '1;33;41' would give you a bold orange text mon a red background. Why not :-)
`,
	Example: `svlogj --ansi-color '1;33;41' --entity NetworkManager`,
	Run: func(cmd *cobra.Command, args []string) {
		gYw := "gYw"
		bgColors := []string{"40", "41", "42", "43", "44", "45", "46", "47"}
		fgColors := []string{"", "1", "30", "1;30", "31", "1;31", "32", "1;32", "33", "1;33", "34", "1;34", "35", "1;35", "36", "1;36", "37", "1;37"}
		fmt.Printf("                 ")
		for _, c := range bgColors {
			fmt.Printf("   %s   ", c+"m")
		}
		fmt.Printf("\n")
		for _, f := range fgColors {
			fmt.Printf(" %6s ", f+"m")
			fmt.Printf(" \033[%sm  %s  \033[0m ", f, gYw)
			for _, b := range bgColors {
				fmt.Printf(" \033[%s;%sm  %s  \033[0m ", f, b, gYw)
			}
			fmt.Printf("\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(ansiColorsCmd)
}
