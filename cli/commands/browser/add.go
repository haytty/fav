package browser

import (
	"github.com/haytty/fav/cli/cli"
	browser "github.com/haytty/fav/internal/handler/fav/browser/add"
	"github.com/spf13/cobra"
)

func AddCommand(c cli.Cli) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add to browsers",
		Long:  "To the data store\nAdd favorite browser.",
		Args:  cobra.MatchAll(cobra.ExactArgs(2)),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			url := args[1]
			return browser.Apply(name, url)
		},
	}
	return addCmd
}
