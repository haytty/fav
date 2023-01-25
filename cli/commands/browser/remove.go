package browser

import "github.com/spf13/cobra"

func RemoveCommand() *cobra.Command {
	rmCmd := &cobra.Command{
		Use:   "remove",
		Short: "Fav is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	return rmCmd
}
