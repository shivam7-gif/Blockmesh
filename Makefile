
.PHONY: build run test clean

build:
	@go build -o build/fs.exe .

run: build
	@.\build\fs.exe

test:
	@go test ./... -v

clean:
	@go clean
