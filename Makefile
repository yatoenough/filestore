.PHONY: build
build:
	go build -v -o ./build/filestore ./cmd/filestore

.PHONY: run
run: build
	./build/filestore

.PHONY: clean
clean:
	go clean
	del /q /s build
.DEFAULT_GOAL := build
