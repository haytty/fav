package commands

import (
	"github.com/haytty/fav/cli/cli"
	"github.com/haytty/fav/cli/commands/browser"
	"github.com/spf13/cobra"
)

func BrowserCommand(c cli.Cli) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "browser",
		Short: "Manage your browser information",
		Long:  "Manage your browser information",
		RunE:  ShowHelp(c.Err()),
	}
	rootCmd.AddCommand(
		browser.AddCommand(c),
		browser.RemoveCommand(c),
		browser.ListCommand(c),
		browser.EditCommand(c),
	)
	return rootCmd
}
