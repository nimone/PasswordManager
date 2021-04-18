.PHONY: build

dev:
	@echo "Building a developer binary for current OS/ARCH"
	go build -o ./build/PasswordManager main.go

build:
	@echo "Building for current OS/ARCH"
	go build -ldflags="-s -w" -o ./build/PasswordManager main.go

linux32:
	@echo "Building for Linux i386"
	GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o ./build/PasswordManager_linux32 main.go

linux64:
	@echo "Building for Linux amd64"
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./build/PasswordManager_linux64 main.go

win32:
	@echo "Building for Windows i386"
	GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o ./build/PasswordManager_win32.exe main.go

win64:
	@echo "Building for Windows amd64"
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./build/PasswordManager_win64.exe main.go

macos:
	@echo "Building for macOS amd64"
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./build/PasswordManager_mac main.go

release: linux32 linux64 win32 win64 macos
	mkdir ./release
	@echo "Packing release files"
	mv ./build/PasswordManager_linux32 ./release/PasswordManager
	tar -czf ./release/passwordManager_linux_386.tar.gz ./release/PasswordManager
	mv ./build/PasswordManager_linux64 ./release/PasswordManager
	tar -czf ./release/passwordManager_linux_amd64.tar.gz ./release/PasswordManager
	mv ./build/PasswordManager_win32.exe ./release/PasswordManager.exe
	zip ./release/passwordManager_windows_386.zip ./release/PasswordManager.exe
	mv ./build/PasswordManager_win64.exe ./release/PasswordManager.exe
	zip ./release/passwordManager_windows_amd64.zip ./release/PasswordManager.exe
	mv ./build/PasswordManager_mac ./release/PasswordManager
	tar -czf ./release/passwordManager_macOS_amd64.tar.gz ./release/PasswordManager
	rm ./release/{PasswordManager.exe,PasswordManager}
	@echo "Packed successfully"
