package commands

import (
	"github.com/spf13/cobra"
	"io"
)

func ShowHelp(err io.Writer) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SetOut(err)
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}
