package config

import (
	"github.com/haytty/fav/internal/util"
	"os"
)

type BaseDir struct {
	Path string
}

var (
	baseDir *BaseDir
)

func SetBaseDir(path string) {
	baseDir = &BaseDir{Path: path}
}

func BaseDirData() *BaseDir {
	return baseDir
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
