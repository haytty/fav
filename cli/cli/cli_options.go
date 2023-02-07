package cli

import (
	"fmt"
	"github.com/haytty/fav/cli/version"
	"github.com/haytty/fav/internal/util"
	"os"
	"strings"
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
	loglevel := os.Getenv(fmt.Sprintf("%s_LOG_LEVEL", strings.ToUpper(version.Name)))
	cli.loglevel = loglevel

	if !util.SliceContains[string](validLogLevels, loglevel) {
		cli.loglevel = "info"
	}
	return nil
}

// TODO not used
func SetConfigDir(cli *FavCli) error {
	cli.configDir = os.Getenv(fmt.Sprintf("%s_CONFIG_DIR", strings.ToUpper(version.Name)))
	return nil
}
