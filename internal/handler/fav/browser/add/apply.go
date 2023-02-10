package add

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/haytty/fav/internal/model"
)

func Apply(name, appPath string) error {
	browser, err := model.LoadBrowser()
	if err != nil {
		return fmt.Errorf("browser load: %w", err)
	}

	browserName := model.BrowserName(name)
	if err := browser.Add(browserName, model.NewBrowserInfo(appPath)); err != nil {
		return fmt.Errorf("browser add: %w", err)
	}

	if err := browser.Save(); err != nil {
		return fmt.Errorf("browser save: %w", err)
	}

	info := browser.Fetch(browserName)
	fmt.Printf(
		color.GreenString("You have registered the following information:\n")+
			color.GreenString("Name: %s\n")+
			color.GreenString("AppPath: %s\n"),
		browserName,
		info.AppPath,
	)

	return nil
}
