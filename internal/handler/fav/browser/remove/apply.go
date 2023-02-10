package remove

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/haytty/fav/internal/model"
)

func Apply(name string) error {
	browser, err := model.LoadBrowser()
	if err != nil {
		return fmt.Errorf("browser load: %w", err)
	}

	browserName := model.BrowserName(name)
	info := browser.Fetch(browserName)

	if err := browser.Remove(browserName); err != nil {
		return fmt.Errorf("browser remove: %w", err)
	}

	if err := browser.Save(); err != nil {
		return fmt.Errorf("browser save: %w", err)
	}

	fmt.Printf(
		color.YellowString("You have removed the following information:\n")+
			color.YellowString("Name: %s\n")+
			color.YellowString("AppPath: %s\n"),
		browserName,
		info.AppPath,
	)

	return nil
}
