package init

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/haytty/fav/cli/version"
	"github.com/haytty/fav/internal/config"
)

type Params struct {
	DataStoreType string
}

// TODO
// 　Prompt whether to overwrite the Config file if it already exists
func Apply(params *Params) error {
	d := config.BaseDirData()
	if err := d.Create(); err != nil {
		return err
	}
	c, err := config.NewConfigWithError(params.DataStoreType, d.Path)
	if err != nil {
		return err
	}
	if err := c.Save(); err != nil {
		return err
	}
	fmt.Printf(
		color.GreenString("%s init complete!\n")+
			color.GreenString("config dir path is %s\n"),
		version.Name,
		d.Path,
	)
	return nil
}
