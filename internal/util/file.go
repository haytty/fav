package util

import (
	"os"
)

func IsFileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func IsDirectoryExist(path string) bool {
	return IsFileExist(path)
}
