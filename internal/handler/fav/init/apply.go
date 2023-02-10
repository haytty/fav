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

func Apply(params *Params) error {
	baseDir := config.BaseDirData()
	if err := baseDir.Create(); err != nil {
		return fmt.Errorf("basedir create: %w", err)
	}

	c, err := config.NewConfigWithError(params.DataStoreType, baseDir.Path)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := c.Save(); err != nil {
		return fmt.Errorf("config save: %w", err)
	}

	fmt.Printf(color.GreenString("%s init complete!\n")+
		color.GreenString("config dir path is %s\n"),
		version.Name,
		baseDir.Path,
	)

	return nil
}
