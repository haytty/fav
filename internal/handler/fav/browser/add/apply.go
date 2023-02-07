package add

import "github.com/haytty/fav/internal/model"

func Apply(name, appPath string) error {
	f, err := model.LoadBrowser()
	if err != nil {
		return err
	}
	if err := f.Add(model.BrowserName(name), model.NewBrowserInfo(appPath)); err != nil {
		return nil
	}
	if err := f.Save(); err != nil {
		return err
	}
	return nil
}
