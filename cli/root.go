package cli

import (
	"github.com/haytty/fav/cli/cli"
	"github.com/haytty/fav/cli/commands"
	"github.com/haytty/fav/cli/flags"
	"github.com/haytty/fav/cli/logger"
	"github.com/haytty/fav/cli/version"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func NewFavCommand(c cli.Cli) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "fav",
		Short: "Fav is a favorite site opener.",
		Long: `This program opens your favorite sites.
               Your favorite browser, 
               Your favorite site from the CLI`,
		Version:           version.CurrentVersion(),
		RunE:              commands.ShowHelp(c.Err()),
		PersistentPreRunE: initializeCommand(c),
	}
	rootCmd.AddCommand(
		commands.AddCommand(c),
		commands.RemoveCommand(c),
		commands.BrowserCommand(c),
		commands.InitCommand(c),
	)

	opts := flags.NewGlobalOption()
	rootCmd.PersistentFlags().StringVarP(
		&opts.BaseDir,
		"base-dir",
		"d",
		filepath.Join(os.Getenv("HOME"), ".config", "fav"),
		"base directory",
	)

	return rootCmd
}

func initializeCommand(c cli.Cli) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return logger.SetupLogger(c)
	}
}
