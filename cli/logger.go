package cli

import (
	"github.com/sirupsen/logrus"
	"os"
)

func SetupLogger(c Cli) error {
	logrus.SetOutput(c.Out())
	err := setLogLevel()
	return err
}

func setLogLevel() error {
	level := os.Getenv("FAV_LOG_LEVEL")
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)
	return nil
}
