package commands

import (
	"github.com/haytty/fav/cli/cli"
	"github.com/haytty/fav/cli/flags"
	fav "github.com/haytty/fav/internal/cmd/fav/add"
	"github.com/spf13/cobra"
)

func AddCommand(c cli.Cli) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add to favorites",
		Long: `To the data store
               Add favorite information.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := flags.NewGlobalOption()
			url := args[0]
			return fav.Apply(url, opts)
		},
	}
	return addCmd
}
