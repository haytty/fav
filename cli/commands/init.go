package commands

import (
	"github.com/haytty/fav/cli/cli"
	fav "github.com/haytty/fav/internal/handler/fav/init"
	"github.com/spf13/cobra"
)

func InitCommand(c cli.Cli) *cobra.Command {
	var dataStore string
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize config file",
		Long:  `Initialize config file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			params := &fav.Params{DataStoreType: dataStore}
			return fav.Apply(params)
		},
	}

	flagName := "data-store"
	initCmd.Flags().StringVarP(&dataStore, flagName, "s", "file", "")
	initCmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"file"}, cobra.ShellCompDirectiveFilterFileExt
	})

	return initCmd
}
