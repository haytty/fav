package browser

import (
	"github.com/haytty/fav/cli/cli"
	browser "github.com/haytty/fav/internal/handler/fav/browser/list"
	"github.com/spf13/cobra"
)

func ListCommand(c cli.Cli) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "list",
		Short: "",
		Long: `
               `,
		Args: cobra.MatchAll(cobra.ExactArgs(0)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return browser.Apply()
		},
	}
	return addCmd
}
