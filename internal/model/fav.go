package model

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/haytty/fav/internal/config"
	"github.com/haytty/fav/internal/datastore"
	"github.com/haytty/fav/internal/util"
	"github.com/manifoldco/promptui"
)

type FavName string

func (f FavName) String() string {
	return string(f)
}

type FavData struct {
	URL string `json:"url"`
}

func NewFavData(url string) *FavData {
	return &FavData{URL: url}
}

type FavMap map[FavName]*FavData

type Fav struct {
	M     FavMap `json:"favMap"`
	store datastore.DataStore
	mu    sync.Mutex
}

func (f *Fav) Fetch(name FavName) *FavData {
	return f.M[name]
}

func (f *Fav) Selection() []string {
	keys := util.MapKeys(f.M)
	castedKeys := util.ConvertToStrings(keys)

	return castedKeys
}

var errAlreadyFavKey = fmt.Errorf("is already exist. please check your fav list")

func (f *Fav) Add(name FavName, data *FavData) error {
	if f.hasName(name) {
		return fmt.Errorf("%s %w", name, errAlreadyFavKey)
	}

	f.mu.Lock()
	f.M[name] = data
	f.mu.Unlock()

	return nil
}

var errFavKeyIsNotFound = fmt.Errorf("is not found. please check your fav data")

func (f *Fav) Remove(name FavName) error {
	if !f.hasName(name) {
		return fmt.Errorf("%s %w", name, errFavKeyIsNotFound)
	}

	f.mu.Lock()
	delete(f.M, name)
	f.mu.Unlock()

	return nil
}

func (f *Fav) Save() error {
	m, err := util.PrettyJSON(f)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	_, err = f.store.WriteWithIdempotency(m)
	if err != nil {
		return fmt.Errorf("%w", err)
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

	i, s, err := prompt.Run()
	if err != nil {
		return i, s, fmt.Errorf("%w", err)
	}

	return i, s, nil
}

func LoadFav() (*Fav, error) {
	registerType := "fav"

	store, err := datastore.NewDataStoreWithError(config.GetConfig(), registerType)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	loadData, err := io.ReadAll(store)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	fav := &Fav{
		store: store,
	}

	// After initialize config file is 0 bite
	if len(loadData) == 0 {
		fm := make(FavMap)
		fav.M = fm

		return fav, nil
	}

	if err := json.Unmarshal(loadData, fav); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return fav, nil
}
