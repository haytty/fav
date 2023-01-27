package browser

import (
	"fmt"
	"github.com/haytty/fav/cli/cli"
	"github.com/spf13/cobra"
)

func AddCommand(c cli.Cli) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "",
		Long: `
               
               `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("browser? add?")
		},
	}
	return addCmd
}
