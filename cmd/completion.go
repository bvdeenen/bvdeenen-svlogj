// Package cmd 
package cmd

import (
	"svlogj/pkg/utils"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "generate a completion for your shell",
	Long: `
	`,
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.AddCommand(bashCompletion)
	completionCmd.AddCommand(zshCompletion)
	completionCmd.AddCommand(fishCompletion)
}

var bashCompletion = &cobra.Command{
	Use:                        "bash",
	Long:                       `Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

        source <(cobra-cli completion bash)

To load completions for every new session, execute once:

#### Linux:

        cobra-cli completion bash > /etc/bash_completion.d/cobra-cli

#### macOS:

        cobra-cli completion bash > /usr/local/etc/bash_completion.d/cobra-cli

You will need to start a new shell for this setup to take effect.
	`,
	ValidArgsFunction:          utils.NoFilesEmptyCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Root().GenBashCompletion(cmd.OutOrStdout())
	},
}
var zshCompletion = &cobra.Command{
	Use:                        "zsh",
	Long: `
	Generate the autocompletion script for the zsh shell.

	If shell completion is not already enabled in your environment you will need
	to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

	To load completions for every new session, execute once:

	#### Linux:

	cobra-cli completion zsh > "${fpath[1]}/_cobra-cli"

	#### macOS:

	cobra-cli completion zsh > /usr/local/share/zsh/site-functions/_cobra-cli

	You will need to start a new shell for this setup to take effect.`,
	ValidArgsFunction:          utils.NoFilesEmptyCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Root().GenZshCompletion(cmd.OutOrStdout())
	},
}
var fishCompletion = &cobra.Command{
	Use:                        "fish",
	Long: `
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

        cobra-cli completion fish | source

To load completions for every new session, execute once:

        cobra-cli completion fish > ~/.config/fish/completions/cobra-cli.fish

You will need to start a new shell for this setup to take effect.
	`,
	ValidArgsFunction:          utils.NoFilesEmptyCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
	},
}
