.PHONY: build clean

APP_NAME := fav

build: fav_build

fav_build:
	go build -ldflags="-s -w" -trimpath -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)
clean:
	rm -fr ./bin/*
