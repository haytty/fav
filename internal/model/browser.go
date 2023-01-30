package model

import "sync"

type BrowserName string
type BrowserInfo struct {
	AppPath string
}

type Browser struct {
	M  *BrowserMap
	mu sync.Mutex
}

type BrowserMap map[BrowserName]BrowserInfo
