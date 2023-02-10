package commands

import (
	"fmt"

	"github.com/haytty/fav/cli/cli"
	fav "github.com/haytty/fav/internal/handler/fav/edit"
	"github.com/spf13/cobra"
)

func EditCommand(cli cli.Cli) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit fav config",
		Long:  `Manually correct your favorite sites information. Register in JSON format`,
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
