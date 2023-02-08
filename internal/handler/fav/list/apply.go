package list

import (
	"fmt"
	"github.com/haytty/fav/internal/model"
	"strings"
)

func Apply() error {
	f, err := model.LoadFav()
	if err != nil {
		return err
	}
	fmt.Println(strings.Join(f.Selection(), "\n"))
	return nil
}
