package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/haytty/fav/cli"
	clistruct "github.com/haytty/fav/cli/cli"
)

func main() {
	c := clistruct.NewFavCli()
	if err := cli.NewFavCommand(c).Execute(); err != nil {
		fmt.Fprintln(c.Out(), color.RedString(err.Error()))
		os.Exit(1)
	}
}
