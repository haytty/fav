package datastore

import "io"

type DataStore interface {
	io.ReadWriteCloser
}
