package version

var (
	Name         = "fav"
	BuildVersion = ""
	Version      = "1.0.0"
)

func CurrentVersion() string {
	return Version
}
