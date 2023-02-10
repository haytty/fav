package browser

import (
	"fmt"

	"github.com/haytty/fav/cli/cli"
	browser "github.com/haytty/fav/internal/handler/fav/browser/edit"
	"github.com/spf13/cobra"
)

func EditCommand(cli cli.Cli) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit fav browser config",
		Long:  `Manually correct your favorite browser information. Register in JSON format`,
		Args:  cobra.MatchAll(cobra.ExactArgs(0)),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := browser.Apply(); err != nil {
				return fmt.Errorf("%w", err)
			}

			return nil
		},
	}

	return addCmd
}
