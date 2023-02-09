package add

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/haytty/fav/internal/model"
)

func Apply(name, url string) error {
	f, err := model.LoadFav()
	if err != nil {
		return err
	}
	favName := model.FavName(name)
	if err := f.Add(favName, model.NewFavData(url)); err != nil {
		return err
	}
	if err := f.Save(); err != nil {
		return err
	}

	data := f.Fetch(favName)
	fmt.Printf(
		color.GreenString("You have registered the following information:\n")+
			color.GreenString("Name: %s\n")+
			color.GreenString("URL: %s\n"),
		favName,
		data.Url,
	)
	return nil
}
