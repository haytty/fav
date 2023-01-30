package commands

import (
	"github.com/haytty/fav/cli/cli"
	"github.com/haytty/fav/cli/flags"
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
			opts := flags.NewGlobalOption()
			params := &fav.Params{DataStore: dataStore}

			return fav.Apply(params, opts)
		},
	}

	initCmd.Flags().StringVarP(&dataStore, "data-store", "s", "file", "")
	return initCmd
}
