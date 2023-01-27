package commands

import (
	"fmt"
	"github.com/haytty/fav/cli/cli"
	"github.com/spf13/cobra"
)

func RemoveCommand(cli cli.Cli) *cobra.Command {
	rmCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove from favorites",
		Long:  ` Removes favorite information from the data store.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("remove?")
		},
	}
	return rmCmd
}
