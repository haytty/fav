package remove

import "github.com/haytty/fav/internal/model"

func Apply(name string) error {
	f, err := model.LoadFav()
	if err != nil {
		return err
	}
	if err := f.Remove(model.FavName(name)); err != nil {
		return nil
	}
	if err := f.Save(); err != nil {
		return err
	}
	return nil
}
