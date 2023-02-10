package commands

import (
	"fmt"
	"os"

	"github.com/haytty/fav/cli/cli"
	fav "github.com/haytty/fav/internal/handler/fav/init"
	"github.com/spf13/cobra"
)

func InitCommand(cli cli.Cli) *cobra.Command {
	var dataStore string

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize config file",
		Long:  `Generates configuration information needed to use the fav command`,
		RunE: func(cmd *cobra.Command, args []string) error {
			params := &fav.Params{DataStoreType: dataStore}

			if err := fav.Apply(params); err != nil {
				return fmt.Errorf("%w", err)
			}

			return nil
		},
	}

	flagName := "data-store"
	initCmd.Flags().StringVarP(&dataStore, flagName, "s", "file", "")

	if err := initCmd.RegisterFlagCompletionFunc(
		flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{"file"}, cobra.ShellCompDirectiveFilterFileExt
		}); err != nil {
		fmt.Fprintln(cli.Err(), err)
		os.Exit(1)
	}

	return initCmd
}
