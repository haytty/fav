package fav

import (
	"fmt"
	"github.com/haytty/fav/internal/model"
	"github.com/haytty/fav/internal/prompt"
)

func Apply() error {
	f, err := model.LoadFav()
	if err != nil {
		return err
	}
	b, err := model.LoadBrowser()
	if err != nil {
		return err
	}

	results, err := prompt.Start(f, b)
	if err != nil {
		return fmt.Errorf("Processing Interrupted.")
	}

	favName, browserName := model.FavName(results[0]), model.BrowserName(results[1])
	data, browserInfo := f.Fetch(favName), b.Fetch(browserName)

	err = browserInfo.Open(data)
	if err != nil {
		return err
	}
	return nil
}
