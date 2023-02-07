package version

import "fmt"

var (
	Name         = "fav"
	BuildVersion = ""
	Version      = "1.0.0"
)

func CurrentVersion() string {
	return fmt.Sprintf("%s", Version)
}
