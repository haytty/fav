package logger

import (
	"fmt"

	"github.com/haytty/fav/cli/cli"
	"github.com/sirupsen/logrus"
)

func SetupLogger(c cli.Cli) error {
	logrus.SetOutput(c.Out())
	logrus.SetFormatter(&logrus.TextFormatter{})

	err := setLogLevel(c)

	return fmt.Errorf("%w", err)
}

func setLogLevel(c cli.Cli) error {
	lvl, err := logrus.ParseLevel(c.LogLevel())
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	logrus.SetLevel(lvl)

	return nil
}
