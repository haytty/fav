package model

import (
	"encoding/json"
	"fmt"
	"github.com/haytty/fav/internal/datastore"
	"path/filepath"
	"sync"
)

const (
	favDataFile = "fav.db"
)

type FavName string
type FavData struct {
	Url string
}

func NewFavData(url string) *FavData {
	return &FavData{Url: url}
}

type FavMap map[FavName]*FavData

type Fav struct {
	M     FavMap
	store datastore.DataStore
	mu    sync.Mutex
}

func (f *Fav) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, f.M)
}

func (f *Fav) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(f.M, "", "    ")
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
	m, err := json.Marshal(f)
	if err != nil {
		return err
	}
	_, err = f.store.Write(m)
	if err != nil {
		return err
	}

	return nil
}

func (f *Fav) hasName(name FavName) bool {
	_, ok := f.M[name]
	return ok
}

func NewFav() *Fav {
	return &Fav{}
}

func LoadFav(configDir string, store datastore.DataStore) (*Fav, error) {
	_ := filepath.Join(configDir, favDataFile)
	var b []byte
	if _, err := store.Read(b); err != nil {
		return nil, err
	}

	f := NewFav()
	if err := json.Unmarshal(b, f); err != nil {
		return nil, err
	}

	f.store = store
	return f, nil
}
