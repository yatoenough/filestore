BINARY_NAME=filestorage

.PHONY: build
build:
	go build -o build/${BINARY_NAME}.exe cmd/filestorage/main.go

run:build
	./build/${BINARY_NAME}.exe

clean:
	go clean
	rd /s /q "./build"