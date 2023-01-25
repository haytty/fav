package commands

import "github.com/spf13/cobra"

func RemoveCommand() *cobra.Command {
	rmCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove from favorites",
		Long: ` Removes favorite information from the data store.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	return rmCmd
}
