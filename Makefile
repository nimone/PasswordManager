.PHONY: build

dev:
	@echo "Building a developer binary for current OS/ARCH"
	go build -o ./build/gopass main.go

build:
	@echo "Building for current OS/ARCH"
	go build -ldflags="-s -w" -o ./build/gopass main.go