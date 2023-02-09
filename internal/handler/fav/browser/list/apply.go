package list

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/haytty/fav/internal/model"
	"strings"
)

func Apply() error {
	f, err := model.LoadBrowser()
	if err != nil {
		return err
	}
	color.Green("The following browsers are registered as favourites:")
	fmt.Println(strings.Join(f.Selection(), "\n"))
	return nil
}
