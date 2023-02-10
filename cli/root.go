package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/haytty/fav/cli/cli"
	"github.com/haytty/fav/cli/commands"
	"github.com/haytty/fav/cli/flags"
	"github.com/haytty/fav/cli/logger"
	"github.com/haytty/fav/cli/version"
	"github.com/haytty/fav/internal/config"
	"github.com/haytty/fav/internal/handler/fav"
	"github.com/haytty/fav/internal/util"
	"github.com/spf13/cobra"
)

func NewFavCommand(cli cli.Cli) *cobra.Command {
	rootCmd := &cobra.Command{ //nolint:exhaustivestruct
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
			return fmt.Errorf("%w", fav.Apply())
		},
		PersistentPreRunE: initialize(cli),
	}
	rootCmd.AddCommand(
		commands.AddCommand(cli),
		commands.RemoveCommand(cli),
		commands.ListCommand(cli),
		commands.BrowserCommand(cli),
		commands.InitCommand(cli),
		commands.EditCommand(cli),
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
			return fmt.Errorf("setup logger: %w", err)
		}

		config.SetBaseDir(opts.BaseDir)

		if util.IsDirectoryExist(opts.BaseDir) {
			err := config.LoadConfig()
			if err != nil {
				return fmt.Errorf("load config: %w", err)
			}
		}

		return nil
	}
}
