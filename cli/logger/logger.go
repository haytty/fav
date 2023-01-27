package logger

import (
	"github.com/haytty/fav/cli/cli"
	"github.com/sirupsen/logrus"
)

func SetupLogger(c cli.Cli) error {
	logrus.SetOutput(c.Out())
	err := setLogLevel(c)
	return err
}

func setLogLevel(c cli.Cli) error {
	lvl, err := logrus.ParseLevel(c.LogLevel())
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)
	return nil
}
