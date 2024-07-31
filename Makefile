CFG_PATH=./configs/local.json

.PHONY: build
build:
	go build -v -o ./build/filestore ./cmd/filestore

.PHONY: run
run: build
	./build/filestore -cfgpath=${CFG_PATH}

.PHONY: clean
clean:
	go clean
	del /q /s build
.DEFAULT_GOAL := build
