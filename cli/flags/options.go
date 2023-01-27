package flags

type GlobalOption struct {
	ConfigDir string
}

var globalOption *GlobalOption

func NewGlobalOption() *GlobalOption {
	if globalOption == nil {
		globalOption = &GlobalOption{}
	}
	return globalOption
}
