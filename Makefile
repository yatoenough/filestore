BINARY_NAME=filestorage

build:
	go build -o build/${BINARY_NAME} cmd/filestorage/main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm build/${BINARY_NAME}