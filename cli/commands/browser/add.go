package browser

import (
	"fmt"

	"github.com/haytty/fav/cli/cli"
	browser "github.com/haytty/fav/internal/handler/fav/browser/add"
	"github.com/spf13/cobra"
)

func AddCommand(cli cli.Cli) *cobra.Command {
	acceptArglength := 2

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add to browsers",
		Long:  "To the data store\nAdd favorite browser.",
		Args:  cobra.MatchAll(cobra.ExactArgs(acceptArglength)),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			url := args[1]
			if err := browser.Apply(name, url); err != nil {
				return fmt.Errorf("%w", err)
			}

			return nil
		},
	}

	return addCmd
}
