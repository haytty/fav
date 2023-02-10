package edit

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/haytty/fav/internal/config"
)

func Apply() error {
	conf := config.GetConfig()
	if err := conf.FavEdit(); err != nil {
		return fmt.Errorf("fav edit: %w", err)
	}

	fmt.Printf(
		color.GreenString("Fixed the config file.\n")+
			color.GreenString("configPath: %s\n"),
		conf.FavConfigFileName,
	)

	return nil
}
