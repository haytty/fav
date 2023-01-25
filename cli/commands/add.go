package commands

import "github.com/spf13/cobra"

func NewAddCommand() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add to favorites",
		Long: `To the data store
               Add favorite information.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	return addCmd
}
