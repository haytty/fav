package browser

import (
	"github.com/haytty/fav/cli/cli"
	browser "github.com/haytty/fav/internal/handler/fav/browser/edit"
	"github.com/spf13/cobra"
)

func EditCommand(c cli.Cli) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit fav browser config",
		Long:  `Manually correct your favorite browser information. Register in JSON format`,
		Args:  cobra.MatchAll(cobra.ExactArgs(0)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return browser.Apply()
		},
	}
	return addCmd
}
