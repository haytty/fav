package model

import (
	"github.com/haytty/fav/internal/datastore"
	"reflect"
	"sync"
	"testing"
)

func TestBrowserInfo_Open(t *testing.T) {
	type fields struct {
		AppPath string
	}
	type args struct {
		data *FavData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &BrowserInfo{
				AppPath: tt.fields.AppPath,
			}
			if err := f.Open(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBrowserName_String(t *testing.T) {
	tests := []struct {
		name string
		f    BrowserName
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBrowser_Add(t *testing.T) {
	type fields struct {
		M     BrowserMap
		store datastore.DataStore
		mu    sync.Mutex
	}
	type args struct {
		name BrowserName
		data *BrowserInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Browser{
				M:     tt.fields.M,
				store: tt.fields.store,
				mu:    tt.fields.mu,
			}
			if err := f.Add(tt.args.name, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBrowser_Fetch(t *testing.T) {
	type fields struct {
		M     BrowserMap
		store datastore.DataStore
		mu    sync.Mutex
	}
	type args struct {
		name BrowserName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BrowserInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Browser{
				M:     tt.fields.M,
				store: tt.fields.store,
				mu:    tt.fields.mu,
			}
			if got := f.Fetch(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBrowser_Remove(t *testing.T) {
	type fields struct {
		M     BrowserMap
		store datastore.DataStore
		mu    sync.Mutex
	}
	type args struct {
		name BrowserName
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Browser{
				M:     tt.fields.M,
				store: tt.fields.store,
				mu:    tt.fields.mu,
			}
			if err := f.Remove(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBrowser_Save(t *testing.T) {
	type fields struct {
		M     BrowserMap
		store datastore.DataStore
		mu    sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Browser{
				M:     tt.fields.M,
				store: tt.fields.store,
				mu:    tt.fields.mu,
			}
			if err := f.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBrowser_Selection(t *testing.T) {
	type fields struct {
		M     BrowserMap
		store datastore.DataStore
		mu    sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Browser{
				M:     tt.fields.M,
				store: tt.fields.store,
				mu:    tt.fields.mu,
			}
			if got := f.Selection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Selection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBrowser_Tui(t *testing.T) {
	type fields struct {
		M     BrowserMap
		store datastore.DataStore
		mu    sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Browser{
				M:     tt.fields.M,
				store: tt.fields.store,
				mu:    tt.fields.mu,
			}
			got, got1, err := f.Tui()
			if (err != nil) != tt.wantErr {
				t.Errorf("Tui() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Tui() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Tui() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBrowser_hasName(t *testing.T) {
	type fields struct {
		M     BrowserMap
		store datastore.DataStore
		mu    sync.Mutex
	}
	type args struct {
		name BrowserName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Browser{
				M:     tt.fields.M,
				store: tt.fields.store,
				mu:    tt.fields.mu,
			}
			if got := f.hasName(tt.args.name); got != tt.want {
				t.Errorf("hasName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadBrowser(t *testing.T) {
	tests := []struct {
		name    string
		want    *Browser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadBrowser()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadBrowser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadBrowser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBrowserInfo(t *testing.T) {
	type args struct {
		appPath string
	}
	tests := []struct {
		name string
		args args
		want *BrowserInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBrowserInfo(tt.args.appPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBrowserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
