package remove

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/haytty/fav/internal/model"
)

func Apply(name string) error {
	f, err := model.LoadBrowser()
	if err != nil {
		return err
	}

	browserName := model.BrowserName(name)
	info := f.Fetch(browserName)

	if err := f.Remove(browserName); err != nil {
		return err
	}
	if err := f.Save(); err != nil {
		return err
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
