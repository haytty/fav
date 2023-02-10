package commands

import (
	"fmt"

	"github.com/haytty/fav/cli/cli"
	fav "github.com/haytty/fav/internal/handler/fav/add"
	"github.com/spf13/cobra"
)

func AddCommand(cli cli.Cli) *cobra.Command {
	acceptArglength := 2

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add to favorite sites",
		Long:  "To the data store\nAdd favorite information.",
		Args:  cobra.MatchAll(cobra.ExactArgs(acceptArglength)),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			url := args[1]
			if err := fav.Apply(name, url); err != nil {
				return fmt.Errorf("%w", err)
			}

			return nil
		},
	}

	return addCmd
}
