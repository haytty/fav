package config

import (
	"encoding/json"
	"fmt"
	"github.com/haytty/fav/internal/datastore/file"
	"github.com/haytty/fav/internal/util"
	"io"
	"path/filepath"
)

type Config struct {
	dataStore             *file.File
	DataStore             string `json:"dataStore"`
	ConfigRootPath        string `json:"configRootPath"`
	ConfigFilePath        string `json:"configFilePath"`
	FavConfigFileName     string `json:"favConfigFileName"`
	BrowserConfigFileName string `json:"browserConfigFileName"`
}

const (
	configFileBaseName = "config.json"
	favDataFile        = "fav.db"
	browserDataFile    = "browser.db"
)

func newFileConfigWithError(dataStore string, configRootPath string) (*Config, error) {
	datastore, err := file.NewFileWithError(configRootPath, filepath.Join(configRootPath, configFileBaseName))
	if err != nil {
		return nil, err
	}
	return &Config{
		dataStore:             datastore,
		DataStore:             dataStore,
		ConfigRootPath:        configRootPath,
		ConfigFilePath:        filepath.Join(configRootPath, configFileBaseName),
		FavConfigFileName:     filepath.Join(configRootPath, favDataFile),
		BrowserConfigFileName: filepath.Join(configRootPath, browserDataFile),
	}, nil
}

func NewConfigWithError(dataStore string, configRootPath string) (*Config, error) {
	switch dataStore {
	case "file":
		return newFileConfigWithError(dataStore, configRootPath)
	}
	err := fmt.Errorf("new config error. check your datastore type.")
	return nil, err
}

func LoadConfig(dir *BaseDir) (*Config, error) {
	datastore, err := file.NewFileWithError(dir.Path, filepath.Join(dir.Path, configFileBaseName))
	if err != nil {
		return nil, err
	}
	conf := &Config{
		dataStore: datastore,
	}
	b, err := io.ReadAll(conf.dataStore)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func (c *Config) Save() error {
	b, err := util.PrettyJson(c)
	if err != nil {
		return err
	}
	_, err = c.dataStore.WriteWithIdempotency(b)
	return err
}
