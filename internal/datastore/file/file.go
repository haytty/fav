package file

import (
	"os"
)

type File struct {
	ConfigRootPath string
	DataFile       *os.File
}

func (f *File) WriteWithIdempotency(p []byte) (n int, err error) {
	_, err = f.DataFile.Seek(0, 0)
	if err != nil {
		return 0, err
	}
	return f.Write(p)
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
	f, err := os.OpenFile(configFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return &File{
		ConfigRootPath: configRootPath,
		DataFile:       f,
	}, nil
}
