package commands

import (
	"github.com/haytty/fav/cli/cli"
	fav "github.com/haytty/fav/internal/handler/fav/add"
	"github.com/spf13/cobra"
)

func AddCommand(c cli.Cli) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add to favorite sites",
		Long:  "To the data store\nAdd favorite information.",
		Args:  cobra.MatchAll(cobra.ExactArgs(2)),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			url := args[1]
			return fav.Apply(name, url)
		},
	}
	return addCmd
}
