package commands

import (
	"github.com/haytty/fav/cli/cli"
	"github.com/haytty/fav/cli/commands/browser"
	"github.com/spf13/cobra"
)

func BrowserCommand(cli cli.Cli) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "browser",
		Short: "Manage your browser information",
		Long:  "Manage your browser information",
		RunE:  ShowHelp(cli.Err()),
	}
	rootCmd.AddCommand(
		browser.AddCommand(cli),
		browser.RemoveCommand(cli),
		browser.ListCommand(cli),
		browser.EditCommand(cli),
	)

	return rootCmd
}
