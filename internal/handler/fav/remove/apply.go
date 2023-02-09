package remove

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/haytty/fav/internal/model"
)

func Apply(name string) error {
	f, err := model.LoadFav()
	if err != nil {
		return err
	}

	favName := model.FavName(name)
	data := f.Fetch(favName)

	if err := f.Remove(favName); err != nil {
		return err
	}
	if err := f.Save(); err != nil {
		return err
	}

	fmt.Printf(
		color.YellowString("You have removed the following information:\n")+
			color.YellowString("Name: %s\n")+
			color.YellowString("URL: %s\n"),
		favName,
		data.Url,
	)
	return nil
}
