package remove

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/haytty/fav/internal/model"
)

func Apply(name string) error {
	fav, err := model.LoadFav()
	if err != nil {
		return fmt.Errorf("fav load: %w", err)
	}

	favName := model.FavName(name)
	data := fav.Fetch(favName)

	if err := fav.Remove(favName); err != nil {
		return fmt.Errorf("fav remove: %w", err)
	}

	if err := fav.Save(); err != nil {
		return fmt.Errorf("fav save: %w", err)
	}

	fmt.Printf(
		color.YellowString("You have removed the following information:\n")+
			color.YellowString("Name: %s\n")+
			color.YellowString("URL: %s\n"),
		favName,
		data.URL,
	)

	return nil
}
