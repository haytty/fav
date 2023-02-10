package commands

import (
	"fmt"

	"github.com/haytty/fav/cli/cli"
	fav "github.com/haytty/fav/internal/handler/fav/list"
	"github.com/spf13/cobra"
)

func ListCommand(cli cli.Cli) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "list",
		Short: "Show favorite site list",
		Long:  `Show the registered favorites site list.`,
		Args:  cobra.MatchAll(cobra.ExactArgs(0)),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := fav.Apply(); err != nil {
				return fmt.Errorf("%w", err)
			}

			return nil
		},
	}

	return addCmd
}
