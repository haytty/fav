package commands

import (
	"github.com/haytty/fav/cli/cli"
	"github.com/haytty/fav/cli/commands/browser"
	"github.com/spf13/cobra"
)

func BrowserCommand(c cli.Cli) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "browser",
		Short: "browser is a favorite browser manipulator.",
		Long: `This program opens your favorite sites.
               Your favorite browser, 
               Your favorite site from the CLI`,
		RunE: ShowHelp(c.Err()),
	}
	rootCmd.AddCommand(
		browser.AddCommand(c),
		browser.RemoveCommand(c),
		browser.ListCommand(c),
	)
	return rootCmd
}
