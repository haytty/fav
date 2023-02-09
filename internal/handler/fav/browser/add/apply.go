package add

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/haytty/fav/internal/model"
)

func Apply(name, appPath string) error {
	f, err := model.LoadBrowser()
	if err != nil {
		return err
	}

	browserName := model.BrowserName(name)
	if err := f.Add(browserName, model.NewBrowserInfo(appPath)); err != nil {
		return err
	}
	if err := f.Save(); err != nil {
		return err
	}

	info := f.Fetch(browserName)
	fmt.Printf(
		color.GreenString("You have registered the following information:\n")+
			color.GreenString("Name: %s\n")+
			color.GreenString("AppPath: %s\n"),
		browserName,
		info.AppPath,
	)
	return nil
}
