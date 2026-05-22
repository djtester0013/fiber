.PHONY: help build run test clean install-deps

help:
	@echo "Available commands:"
	@echo "  make install-deps  - Download and tidy dependencies"
	@echo "  make build         - Build the server binary"
	@echo "  make run           - Run the server"
	@echo "  make test          - Run tests"
	@echo "  make clean         - Clean build artifacts"

install-deps:
	go mod download
	go mod tidy

build: install-deps
	go build -o autotechno-server main.go

run:
	go run main.go

test:
	go test -v ./...

clean:
	rm -f autotechno-server
