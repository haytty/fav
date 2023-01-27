package browser

import (
	"fmt"
	"github.com/haytty/fav/cli/cli"
	"github.com/spf13/cobra"
)

func RemoveCommand(c cli.Cli) *cobra.Command {
	rmCmd := &cobra.Command{
		Use:   "remove",
		Short: "",
		Long: `
               
               `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("browser? remove?")
		},
	}
	return rmCmd
}
