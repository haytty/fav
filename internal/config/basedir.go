package config

import (
	"github.com/haytty/fav/internal/util"
	"os"
)

type BaseDir struct {
	Path string
}

func NewBaseDir(path string) *BaseDir {
	return &BaseDir{Path: path}
}

func (d *BaseDir) Create() error {
	if d.Exist() {
		return nil
	}
	return os.MkdirAll(d.Path, 0700)
}

func (d *BaseDir) Exist() bool {
	return util.IsDirectoryExist(d.Path)
}
