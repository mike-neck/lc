.PHONY: build-osx
build-osx:
	GOOS=darwin GOARCH=arm64 go build -o build/lc .
