package commands

import (
	"github.com/haytty/fav/cli/cli"
	fav "github.com/haytty/fav/internal/handler/fav/remove"
	"github.com/spf13/cobra"
)

func RemoveCommand(cli cli.Cli) *cobra.Command {
	rmCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove from favorites",
		Long:  ` Removes favorite information from the data store.`,
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			return fav.Apply(name)
		},
	}
	return rmCmd
}
