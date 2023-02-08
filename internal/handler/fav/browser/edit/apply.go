package edit

import (
	"github.com/haytty/fav/internal/config"
)

func Apply() error {
	conf := config.ConfigData()
	err := conf.BrowserEdit()
	if err != nil {
		return err
	}
	return nil
}
