package version

import "fmt"

var (
	Version = "1.0.0"
)

func CurrentVersion() string {
	return fmt.Sprintf("%s", Version)
}
