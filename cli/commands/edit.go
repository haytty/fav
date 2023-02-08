package commands

import (
	"github.com/haytty/fav/cli/cli"
	fav "github.com/haytty/fav/internal/handler/fav/edit"
	"github.com/spf13/cobra"
)

func EditCommand(c cli.Cli) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "edit",
		Short: "",
		Long: `
               `,
		Args: cobra.MatchAll(cobra.ExactArgs(0)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return fav.Apply()
		},
	}
	return addCmd
}