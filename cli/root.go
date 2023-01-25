package cli

import (
	"github.com/haytty/fav/cli/commands"
	"github.com/haytty/fav/cli/commands/browser"
	"github.com/spf13/cobra"
)

func NewFavCommand(c Cli) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "fav",
		Short: "Fav is a favorite site opener.",
		Long: `This program opens your favorite sites.
               Your favorite browser, 
               Your favorite site from the CLI`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return SetupLogger(c)
		},
	}
	rootCmd.AddCommand(
		// fav commands
		commands.NewAddCommand(),
		commands.NewRemoveCommand(),

		// fav browser commands
		browser.NewAddCommand(),
		browser.NewRemoveCommand(),
	)
	return rootCmd
}
