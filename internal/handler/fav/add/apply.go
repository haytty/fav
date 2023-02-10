package add

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/haytty/fav/internal/model"
)

func Apply(name, url string) error {
	fav, err := model.LoadFav()
	if err != nil {
		return fmt.Errorf("fav load: %w", err)
	}

	favName := model.FavName(name)
	if err := fav.Add(favName, model.NewFavData(url)); err != nil {
		return fmt.Errorf("fav add: %w", err)
	}

	if err := fav.Save(); err != nil {
		return fmt.Errorf("fav save: %w", err)
	}

	data := fav.Fetch(favName)
	fmt.Printf(
		color.GreenString("You have registered the following information:\n")+
			color.GreenString("Name: %s\n")+
			color.GreenString("URL: %s\n"),
		favName,
		data.URL,
	)

	return nil
}
