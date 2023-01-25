package main

import (
	"fmt"
	"github.com/haytty/fav/cli"
	"os"
)

func main() {
	c := cli.NewDefaultCli()
	if err := cli.NewFavCommand(c).Execute(); err != nil {
		fmt.Fprintln(c.Out(), err)
		os.Exit(1)
	}
}
