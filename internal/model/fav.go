package model

import (
	"encoding/json"
	"fmt"
	"github.com/haytty/fav/internal/config"
	"github.com/haytty/fav/internal/datastore"
	"github.com/haytty/fav/internal/util"
	"github.com/manifoldco/promptui"
	"io"
	"sync"
)

type FavName string

func (f FavName) String() string {
	return string(f)
}

type FavData struct {
	Url string `json:"url"`
}

func NewFavData(url string) *FavData {
	return &FavData{Url: url}
}

type FavMap map[FavName]*FavData

type Fav struct {
	M     FavMap `json:"fav_map"`
	store datastore.DataStore
	mu    sync.Mutex
}

func (f *Fav) Fetch(name FavName) *FavData {
	return f.M[name]
}

func (f *Fav) Selection() []string {
	keys := util.MapKeys(f.M)
	casted_keys := util.ConvertToStrings(keys)
	return casted_keys
}

func (f *Fav) Add(name FavName, data *FavData) error {
	if f.hasName(name) {
		return fmt.Errorf("%s is already exist. please check your fav data.", name)
	}
	f.mu.Lock()
	f.M[name] = data
	f.mu.Unlock()
	return nil
}

func (f *Fav) Remove(name FavName) error {
	if !f.hasName(name) {
		return fmt.Errorf("%s is not found. please check your fav data.", name)
	}
	f.mu.Lock()
	delete(f.M, name)
	f.mu.Unlock()
	return nil
}

func (f *Fav) Save() error {
	m, err := util.PrettyJson(f)
	if err != nil {
		return err
	}
	_, err = f.store.WriteWithIdempotency(m)
	if err != nil {
		return err
	}
	return nil
}

func (f *Fav) hasName(name FavName) bool {
	_, ok := f.M[name]
	return ok
}

func (f *Fav) Tui() (int, string, error) {
	prompt := promptui.Select{
		Label: "my favorite sites",
		Items: f.Selection(),
	}
	return prompt.Run()
}

func LoadFav() (*Fav, error) {
	registerType := "fav"
	store, err := datastore.NewDataStoreWithError(config.ConfigData(), registerType)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(store)
	if err != nil {
		return nil, err
	}

	f := &Fav{
		store: store,
	}
	// After initalize config file is 0 bite
	if len(b) <= 0 {
		fm := make(FavMap)
		f.M = fm
		return f, nil
	}
	if err := json.Unmarshal(b, f); err != nil {
		return nil, err
	}
	return f, nil
}
