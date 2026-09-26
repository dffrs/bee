.DEFAULT_GOAL := build

VERSION ?= dev
LDFLAGS := -X bee/internal/commands.Version=$(VERSION)

.PHONY: clean fmt vet build

clean:
				go clean
fmt: clean
				go fmt ./...
vet: fmt
				go vet ./...
build: vet
				go build -ldflags "$(LDFLAGS)" -o bin/bee cmd/main.go
