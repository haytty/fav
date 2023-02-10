package fav

import (
	"fmt"

	"github.com/haytty/fav/internal/model"
	"github.com/haytty/fav/internal/prompt"
)

func Apply() error {
	fav, err := model.LoadFav()
	if err != nil {
		return fmt.Errorf("load fav: %w", err)
	}

	browser, err := model.LoadBrowser()
	if err != nil {
		return fmt.Errorf("load browser: %w", err)
	}

	results, err := prompt.Start(fav, browser)
	if err != nil {
		return fmt.Errorf("prompt start: processing interrupted %w", err)
	}

	favName, browserName := model.FavName(results[0]), model.BrowserName(results[1])
	data, browserInfo := fav.Fetch(favName), browser.Fetch(browserName)

	err = browserInfo.Open(data)
	if err != nil {
		return fmt.Errorf("open browser: %w", err)
	}

	return nil
}
