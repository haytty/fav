package commands

import (
	"github.com/haytty/fav/cli/cli"
	"github.com/haytty/fav/cli/flags"
	fav "github.com/haytty/fav/internal/handler/fav/add"
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
			name := args[0]
			url := args[1]
			return fav.Apply(name, url, opts)
		},
	}
	return addCmd
}
