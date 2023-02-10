package datastore

import (
	"errors"
	"fmt"
	"io"

	"github.com/haytty/fav/internal/config"
	"github.com/haytty/fav/internal/datastore/file"
)

type DataStore interface {
	io.ReadWriteCloser
	WriteWithIdempotency(p []byte) (n int, err error)
}

var errUndefinedDatastoreType = fmt.Errorf("undefined datastore. please check your fav config")

func NewDataStoreWithError(c *config.Config, registerType string) (DataStore, error) {
	switch c.DataStore {
	case "file":
		datastore, err := newFileWithError(c, registerType)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return datastore, nil
	}

	return nil, errUndefinedDatastoreType
}

var errUndefinedType = errors.New("undefined type")

func newFileWithError(c *config.Config, registerType string) (DataStore, error) {
	switch registerType {
	case "fav":
		return file.NewFileWithError(c.ConfigRootPath, c.FavConfigFileName) //nolint:wrapcheck
	case "browser":
		return file.NewFileWithError(c.ConfigRootPath, c.BrowserConfigFileName)
	default:
		return nil, errUndefinedType
	}
}
