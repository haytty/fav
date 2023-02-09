package cli

import (
	"fmt"
	"github.com/haytty/fav/cli/cli"
	"github.com/haytty/fav/cli/commands"
	"github.com/haytty/fav/cli/flags"
	"github.com/haytty/fav/cli/logger"
	"github.com/haytty/fav/cli/version"
	"github.com/haytty/fav/internal/config"
	"github.com/haytty/fav/internal/handler/fav"
	"github.com/haytty/fav/internal/util"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func NewFavCommand(c cli.Cli) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "fav",
		Short: "Fav is a favorite site opener.",
		Long: fmt.Sprintln(
			"This program opens your favorite sites.\n" +
				"Your favorite browser.\n" +
				"Your favorite site from the CLI"),
		Version:       version.CurrentVersion(),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return fav.Apply()
		},
		PersistentPreRunE: initialize(c),
	}
	rootCmd.AddCommand(
		commands.AddCommand(c),
		commands.RemoveCommand(c),
		commands.ListCommand(c),
		commands.BrowserCommand(c),
		commands.InitCommand(c),
		commands.EditCommand(c),
	)

	opts := flags.NewGlobalOption()
	flagName := "base-dir"
	defaultDir := filepath.Join(os.Getenv("HOME"), ".config", "fav")
	rootCmd.PersistentFlags().StringVarP(
		&opts.BaseDir,
		flagName,
		"d",
		defaultDir,
		"base directory",
	)
	if err := rootCmd.RegisterFlagCompletionFunc(
		flagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{defaultDir}, cobra.ShellCompDirectiveFilterFileExt
		}); err != nil {
		os.Exit(1)
	}

	return rootCmd
}

func initialize(c cli.Cli) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		opts := flags.NewGlobalOption()
		if err := logger.SetupLogger(c); err != nil {
			return err
		}

		config.SetBaseDir(opts.BaseDir)
		if util.IsDirectoryExist(opts.BaseDir) {
			err := config.LoadConfig()
			if err != nil {
				return err
			}
		}

		return nil
	}
}
