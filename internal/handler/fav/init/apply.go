package init

import (
	"fmt"
	"github.com/haytty/fav/cli/flags"
	"github.com/haytty/fav/cli/version"
	"github.com/haytty/fav/internal/config"
)

type Params struct {
	DataStore string
}

func Apply(params *Params, opts *flags.GlobalOption) error {
	d := config.NewBaseDir(opts.BaseDir)
	if err := d.Create(); err != nil {
		return err
	}
	c, err := config.NewConfigWithError(params.DataStore, d.Path)
	if err != nil {
		return err
	}
	if err := c.Save(); err != nil {
		return err
	}
	fmt.Printf(
		"%s init complete!\n"+
			"config dir path is %s\n",
		version.Name,
		d.Path,
	)
	return nil
}
