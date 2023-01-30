package init

import (
	"github.com/haytty/fav/cli/flags"
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

	c := config.NewConfig(
		params.DataStore,
	)
	return c.Save()
}
