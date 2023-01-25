package commands

import (
	"github.com/haytty/fav/cli/commands/browser"
	"github.com/spf13/cobra"
)

func BrowserCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "browser",
		Short: "browser is a favorite browser manipulator.",
		Long: `This program opens your favorite sites.
               Your favorite browser, 
               Your favorite site from the CLI`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	rootCmd.AddCommand(
		browser.AddCommand(),
		browser.RemoveCommand(),
	)
	return rootCmd

}
