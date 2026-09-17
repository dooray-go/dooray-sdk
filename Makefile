GO ?= go
GOARCH ?= amd64

.PHONY: build-all build-linux build-windows build-darwin test clean

build-all: build-linux build-windows build-darwin

build-linux:
	GOOS=linux GOARCH=$(GOARCH) $(GO) build ./...

build-windows:
	GOOS=windows GOARCH=$(GOARCH) $(GO) build ./...

build-darwin:
	GOOS=darwin GOARCH=$(GOARCH) $(GO) build ./...

test:
	$(GO) test ./...

clean:
	$(GO) clean ./...
