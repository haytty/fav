package config

import "os"

type BaseDir struct {
	Path string
}

var RootDir *BaseDir

func NewBaseDir(path string) *BaseDir {
	if RootDir != nil {
		return RootDir
	}
	RootDir = &BaseDir{Path: path}
	return RootDir
}

func (d *BaseDir) Create() error {
	if ok, _ := d.Exist(); ok {
		return nil
	}
	return os.MkdirAll(d.Path, 0700)
}

func (d *BaseDir) Exist() (bool, error) {
	if _, err := os.Stat(d.Path); err != nil {
		return false, err
	}
	return true, nil
}
