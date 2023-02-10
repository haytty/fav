package list

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/haytty/fav/internal/model"
)

func Apply() error {
	fav, err := model.LoadFav()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	color.Green("The following sites are registered as favourites:")
	fmt.Println(strings.Join(fav.Selection(), "\n"))

	return nil
}
