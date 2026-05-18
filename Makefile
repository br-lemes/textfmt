.PHONY: build linux release test version windows

BINARY_NAME := $(notdir $(shell go list -m 2>/dev/null))
ifeq ($(BINARY_NAME),)
    BINARY_NAME := $(notdir $(CURDIR))
endif

IS_NIXOS := $(shell test -f /etc/NIXOS && echo "yes" || echo "")

build: test
	@go build -ldflags "-s -w"

linux: test
	@GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
ifneq ($(IS_NIXOS),)
	@patchelf --remove-rpath --set-interpreter /lib64/ld-linux-x86-64.so.2 $(BINARY_NAME)
endif

release: version linux windows
	@go run ./tools/release/main.go

test:
	@go test ./...

version: test
	@go run ./tools/version/main.go

windows: test
	@GOOS=windows GOARCH=amd64 go build -ldflags "-s -w"
