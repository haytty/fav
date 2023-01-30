package config

import (
	"bytes"
	"github.com/haytty/fav/internal/util"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	DataStore string
}

const (
	configFileBaseName = "config.json"
)

func NewConfig(dataStore string) *Config {
	return &Config{
		DataStore: dataStore,
	}
}

func (c *Config) Save() error {
	configFileName := filepath.Join(RootDir.Path, configFileBaseName)

	f, err := os.Create(configFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := util.PrettyJson(c)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(b)
	_, err = io.Copy(f, buf)
	return err
}
