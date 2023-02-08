package model

import (
	"encoding/json"
	"fmt"
	"github.com/haytty/fav/internal/config"
	"github.com/haytty/fav/internal/datastore"
	"github.com/haytty/fav/internal/util"
	"github.com/manifoldco/promptui"
	"io"
	"os/exec"
	"sync"
)

type BrowserName string

func (f BrowserName) String() string {
	return string(f)
}

type BrowserInfo struct {
	AppPath string `json:"appPath"`
}

func (f *BrowserInfo) Open(data *FavData) ([]byte, error) {
	cmd := exec.Command("open", "-a", f.AppPath, data.Url)
	return cmd.Output()
}

type Browser struct {
	M     BrowserMap `json:"browser_map"`
	store datastore.DataStore
	mu    sync.Mutex
}

func (f *Browser) Fetch(name BrowserName) *BrowserInfo {
	return f.M[name]
}

func (f *Browser) Selection() []string {
	keys := util.MapKeys(f.M)
	casted_keys := util.ConvertToStrings(keys)
	return casted_keys
}

func NewBrowserInfo(appPath string) *BrowserInfo {
	return &BrowserInfo{AppPath: appPath}
}

type BrowserMap map[BrowserName]*BrowserInfo

func (f *Browser) Add(name BrowserName, data *BrowserInfo) error {
	if f.hasName(name) {
		return fmt.Errorf("%s is already exist. please check your fav data.", name)
	}
	f.mu.Lock()
	f.M[name] = data
	f.mu.Unlock()
	return nil
}

func (f *Browser) Remove(name BrowserName) error {
	if !f.hasName(name) {
		return fmt.Errorf("%s is not found. please check your fav data.", name)
	}
	f.mu.Lock()
	delete(f.M, name)
	f.mu.Unlock()
	return nil
}

func (f *Browser) Save() error {
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

func (f *Browser) hasName(name BrowserName) bool {
	_, ok := f.M[name]
	return ok
}

func (f *Browser) Tui() (int, string, error) {
	prompt := promptui.Select{
		Label: "which browser is?",
		Items: f.Selection(),
	}
	return prompt.Run()
}

func LoadBrowser() (*Browser, error) {
	registerType := "browser"
	store, err := datastore.NewDataStoreWithError(config.ConfigData(), registerType)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(store)
	if err != nil {
		return nil, err
	}

	f := &Browser{
		store: store,
	}
	// After initalize config file is 0 bite
	if len(b) <= 0 {
		fm := make(BrowserMap)
		f.M = fm
		return f, nil
	}
	if err := json.Unmarshal(b, f); err != nil {
		return nil, err
	}
	return f, nil
}
