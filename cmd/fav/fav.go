package main

import (
	"fmt"
	"github.com/haytty/fav/cli"
	clistruct "github.com/haytty/fav/cli/cli"
	"os"
)

func main() {
	c := clistruct.NewFavCli()
	if err := cli.NewFavCommand(c).Execute(); err != nil {
		fmt.Fprint(c.Err(), err)
		os.Exit(1)
	}
}
