package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/haytty/fav/internal/command"
	"github.com/haytty/fav/internal/datastore/file"
	"github.com/haytty/fav/internal/util"
)

type Config struct {
	dataStore             *file.File
	DataStore             string `json:"dataStore"`
	ConfigRootPath        string `json:"configRootPath"`
	ConfigFilePath        string `json:"configFilePath"`
	FavConfigFileName     string `json:"favConfigFileName"`
	BrowserConfigFileName string `json:"browserConfigFileName"`
}

func (c *Config) FavEdit() error {
	cmd := command.NewCommand(Editor(), c.FavConfigFileName)
	if err := cmd.Execute(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (c *Config) BrowserEdit() error {
	cmd := command.NewCommand(Editor(), c.BrowserConfigFileName)
	if err := cmd.Execute(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

const (
	configFileBaseName = "config.json"
	favDataFile        = "fav.db"
	browserDataFile    = "browser.db"
)

var cachedData *Config

func newFileConfigWithError(dataStore string, configRootPath string) (*Config, error) {
	datastore, err := file.NewFileWithError(configRootPath, filepath.Join(configRootPath, configFileBaseName))
	if err != nil {
		return nil, fmt.Errorf("%w", err)
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

var errUndefinedDatastoreType = fmt.Errorf("new config error. check your datastore type")

func NewConfigWithError(dataStore string, configRootPath string) (*Config, error) {
	switch dataStore {
	case "file":
		return newFileConfigWithError(dataStore, configRootPath)
	}

	return nil, errUndefinedDatastoreType
}

func LoadConfig() error {
	dir := BaseDirData()

	datastore, err := file.NewFileWithError(dir.Path, filepath.Join(dir.Path, configFileBaseName))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	conf := &Config{
		dataStore: datastore,
	}

	b, err := io.ReadAll(conf.dataStore)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(b, conf); err != nil {
		return fmt.Errorf("%w", err)
	}

	cachedData = conf

	return nil
}

func GetConfig() *Config {
	if cachedData == nil {
		err := LoadConfig()
		if err != nil {
			os.Exit(1)
		}
	}

	return cachedData
}

func (c *Config) Save() error {
	b, err := util.PrettyJSON(c)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	_, err = c.dataStore.WriteWithIdempotency(b)

	return fmt.Errorf("%w", err)
}
