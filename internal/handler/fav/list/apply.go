package list

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/haytty/fav/internal/model"
	"strings"
)

func Apply() error {
	f, err := model.LoadFav()
	if err != nil {
		return err
	}

	color.Green("The following sites are registered as favourites:")
	fmt.Println(strings.Join(f.Selection(), "\n"))
	return nil
}
