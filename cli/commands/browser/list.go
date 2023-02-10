package browser

import (
	"fmt"

	"github.com/haytty/fav/cli/cli"
	browser "github.com/haytty/fav/internal/handler/fav/browser/list"
	"github.com/spf13/cobra"
)

func ListCommand(cli cli.Cli) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "list",
		Short: "Show favorite browser list",
		Long:  `Show the registered favorites browser list.`,
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
