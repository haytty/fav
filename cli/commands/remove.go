package commands

import (
	"fmt"

	"github.com/haytty/fav/cli/cli"
	fav "github.com/haytty/fav/internal/handler/fav/remove"
	"github.com/spf13/cobra"
)

func RemoveCommand(cli cli.Cli) *cobra.Command {
	rmCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove from favorite sites",
		Long:  `Removes favorite site information from the data store.`,
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			if err := fav.Apply(name); err != nil {
				return fmt.Errorf("%w", err)
			}

			return nil
		},
	}

	return rmCmd
}
