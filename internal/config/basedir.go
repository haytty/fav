package config

import (
	"fmt"
	"os"

	"github.com/haytty/fav/internal/util"
)

type BaseDir struct {
	Path string
}

var baseDir *BaseDir

func SetBaseDir(path string) {
	baseDir = &BaseDir{Path: path}
}

func BaseDirData() *BaseDir {
	return baseDir
}

var configDirPerm = 0o700

func (d *BaseDir) Create() error {
	if d.Exist() {
		return nil
	}

	if err := os.MkdirAll(d.Path, os.FileMode(configDirPerm)); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (d *BaseDir) Exist() bool {
	return util.IsDirectoryExist(d.Path)
}
