package remove

import "github.com/haytty/fav/internal/model"

func Apply(name string) error {
	f, err := model.LoadBrowser()
	if err != nil {
		return err
	}
	if err := f.Remove(model.BrowserName(name)); err != nil {
		return nil
	}
	if err := f.Save(); err != nil {
		return err
	}
	return nil
}
