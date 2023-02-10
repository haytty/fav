package browser

import (
	"fmt"

	"github.com/haytty/fav/cli/cli"
	browser "github.com/haytty/fav/internal/handler/fav/browser/remove"
	"github.com/spf13/cobra"
)

func RemoveCommand(cli cli.Cli) *cobra.Command {
	rmCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove from favorite browsers",
		Long:  `Removes favorite browser information from the data store.`,
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			if err := browser.Apply(name); err != nil {
				return fmt.Errorf("%w", err)
			}

			return nil
		},
	}

	return rmCmd
}
