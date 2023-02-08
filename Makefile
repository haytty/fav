.PHONY: build clean install

APP_NAME := fav

install: build copy
uninstall:
	rm -fr /usr/local/bin/$(APP_NAME)
build: fav_build

fav_build:
	go build -ldflags="-s -w" -trimpath -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)
clean:
	rm -fr ./bin/*
copy:
	cp -pr ./bin/$(APP_NAME) /usr/local/bin/$(APP_NAME)
