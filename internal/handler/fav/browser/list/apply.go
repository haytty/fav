package list

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/haytty/fav/internal/model"
)

func Apply() error {
	browser, err := model.LoadBrowser()
	if err != nil {
		return fmt.Errorf("browser load: %w", err)
	}

	color.Green("The following browsers are registered as favourites:")
	fmt.Println(strings.Join(browser.Selection(), "\n"))

	return nil
}
