package model

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/haytty/fav/internal/command"
	"github.com/haytty/fav/internal/config"
	"github.com/haytty/fav/internal/datastore"
	"github.com/haytty/fav/internal/util"
	"github.com/manifoldco/promptui"
)

type BrowserName string

func (f BrowserName) String() string {
	return string(f)
}

type BrowserInfo struct {
	AppPath string `json:"appPath"`
}

func (f *BrowserInfo) Open(data *FavData) error {
	cmd := command.NewCommand("open", "-a", f.AppPath, data.URL)
	if err := cmd.Execute(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

type Browser struct {
	M     BrowserMap `json:"browserMap"`
	store datastore.DataStore
	mu    sync.Mutex
}

func (f *Browser) Fetch(name BrowserName) *BrowserInfo {
	return f.M[name]
}

func (f *Browser) Selection() []string {
	keys := util.MapKeys(f.M)
	castedKeys := util.ConvertToStrings(keys)

	return castedKeys
}

func NewBrowserInfo(appPath string) *BrowserInfo {
	return &BrowserInfo{AppPath: appPath}
}

type BrowserMap map[BrowserName]*BrowserInfo

var errAlreadyBrowserKey = fmt.Errorf("is already exist. please check your browser list")

func (f *Browser) Add(name BrowserName, data *BrowserInfo) error {
	if f.hasName(name) {
		return fmt.Errorf("%s %w", name, errAlreadyBrowserKey)
	}

	f.mu.Lock()
	f.M[name] = data
	f.mu.Unlock()

	return nil
}

var errBrowserKeyIsNotFound = fmt.Errorf("is not found. please check your browser info")

func (f *Browser) Remove(name BrowserName) error {
	if !f.hasName(name) {
		return fmt.Errorf("%s %w", name, errBrowserKeyIsNotFound)
	}

	f.mu.Lock()
	delete(f.M, name)
	f.mu.Unlock()

	return nil
}

func (f *Browser) Save() error {
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

func (f *Browser) hasName(name BrowserName) bool {
	_, ok := f.M[name]

	return ok
}

func (f *Browser) Tui() (int, string, error) {
	prompt := promptui.Select{
		Label: "which browser is?",
		Items: f.Selection(),
	}

	i, s, err := prompt.Run()
	if err != nil {
		return i, s, fmt.Errorf("%w", err)
	}

	return i, s, nil
}

func LoadBrowser() (*Browser, error) {
	registerType := "browser"

	store, err := datastore.NewDataStoreWithError(config.GetConfig(), registerType)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	loadData, err := io.ReadAll(store)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	browser := &Browser{
		store: store,
	}

	// After initialize config file is 0 bite
	if len(loadData) == 0 {
		fm := make(BrowserMap)
		browser.M = fm

		return browser, nil
	}

	if err := json.Unmarshal(loadData, browser); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return browser, nil
}
