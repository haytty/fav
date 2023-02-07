package datastore

import (
	"fmt"
	"github.com/haytty/fav/internal/config"
	"github.com/haytty/fav/internal/datastore/file"
	"io"
)

type DataStore interface {
	io.ReadWriteCloser
	WriteWithIdempotency(p []byte) (n int, err error)
}

func NewDataStoreWithError(c *config.Config, registerType string) (DataStore, error) {
	switch c.DataStore {
	case "file":
		datastore, err := newFileWithError(c, registerType)
		return datastore, err
	}
	err := fmt.Errorf("undefined datastore. please check your fav config.")
	return nil, err
}

func newFileWithError(c *config.Config, registerType string) (DataStore, error) {
	switch registerType {
	case "fav":
		return file.NewFileWithError(c.ConfigRootPath, c.FavConfigFileName)
	case "browser":
		return file.NewFileWithError(c.ConfigRootPath, c.FavConfigFileName)
	}
	err := fmt.Errorf("undefined type")
	return nil, err
}
