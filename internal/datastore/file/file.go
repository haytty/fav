package file

import (
	"fmt"
	"os"
)

type File struct {
	ConfigRootPath string
	DataFile       *os.File
}

func (f *File) WriteWithIdempotency(p []byte) (n int, err error) {
	realName := f.DataFile.Name()
	tmpName := fmt.Sprintf("%s.bk", realName)

	err = os.Rename(realName, tmpName)
	if err != nil {
		return 0, err
	}
	f2, err := fileOpen(realName)
	if err != nil {
		return 0, err
	}
	defer f2.Close()
	n, err = f2.Write(p)
	if err != nil {
		return 0, err
	}
	err = os.Remove(tmpName)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (f *File) Close() error {
	return f.DataFile.Close()
}

func (f *File) Read(p []byte) (n int, err error) {
	return f.DataFile.Read(p)
}

func (f *File) Write(p []byte) (n int, err error) {
	return f.DataFile.Write(p)
}

func NewFileWithError(configRootPath string, configFileName string) (*File, error) {
	f, err := fileOpen(configFileName)
	if err != nil {
		return nil, err
	}
	return &File{
		ConfigRootPath: configRootPath,
		DataFile:       f,
	}, nil
}

func fileOpen(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModePerm)
}
