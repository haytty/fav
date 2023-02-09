package edit

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/haytty/fav/internal/config"
)

func Apply() error {
	conf := config.ConfigData()
	err := conf.FavEdit()
	if err != nil {
		return err
	}

	fmt.Printf(
		color.GreenString("Fixed the config file.\n")+
			color.GreenString("configPath: %s\n"),
		conf.FavConfigFileName,
	)
	return nil
}
