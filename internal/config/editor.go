package config

import (
	"os"
)

func Editor() string {
	var e string
	if e = os.Getenv("VISUAL"); e != "" {
		return e
	}
	if e = os.Getenv("EDITOR"); e != "" {
		return e
	}
	return "vim"
}
