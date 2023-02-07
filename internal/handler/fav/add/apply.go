package add

import (
	"github.com/haytty/fav/cli/flags"
	"github.com/haytty/fav/internal/config"
	"github.com/haytty/fav/internal/datastore"
	"github.com/haytty/fav/internal/model"
)

func Apply(name, url string, opts *flags.GlobalOption) error {
	d := config.NewBaseDir(opts.BaseDir)
	conf, err := config.LoadConfig(d)
	if err != nil {
		return err
	}

	//TODO Fix me
	// bad string type.
	registerType := "fav"
	store, err := datastore.NewDataStoreWithError(conf, registerType)
	if err != nil {
		return err
	}

	f, err := model.LoadFav(store)
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
