package cli

import (
	"github.com/haytty/fav/internal/util"
	"os"
)

type cliOption func(*FavCli) error

func SetApplicationMode(cli *FavCli) error {
	cli.appMode = os.Getenv("APP_ENV")
	return nil
}

func SetLogLevel(cli *FavCli) error {
	validLogLevels := []string{
		"panic",
		"fatal",
		"error",
		"warn",
		"warning",
		"info",
		"debug",
		"trace",
	}
	loglevel := os.Getenv("FAV_LOG_LEVEL")
	cli.loglevel = loglevel

	if !util.SliceContains[string](validLogLevels, loglevel) {
		cli.loglevel = "info"
	}
	return nil
}

func SetConfigDir(cli *FavCli) error {
	cli.configDir = os.Getenv("FAV_CONFIG_DIR")
	return nil
}
