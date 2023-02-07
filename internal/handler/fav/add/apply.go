package add

import (
	"github.com/haytty/fav/internal/model"
)

func Apply(name, url string) error {
	f, err := model.LoadFav()
	if err != nil {
		return err
	}
	if err := f.Add(model.FavName(name), model.NewFavData(url)); err != nil {
		return nil
	}
	if err := f.Save(); err != nil {
		return err
	}
	return nil
}
