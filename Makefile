EXECUTABLE=bin/wakemypc
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
LINUX_ARM=$(EXECUTABLE)_linux_arm64
DARWIN=$(EXECUTABLE)_darwin_amd64
DARWIN_ARM=$(EXECUTABLE)_darwin_arm64
VERSION=$(shell git describe --tags --always --long --dirty)

.PHONY: all test clean

all: test build ## Build and run tests

build: bin windows linux darwin ## Build binaries
	@echo version: $(VERSION)

bin:
	mkdir -p bin
windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) $(LINUX_ARM)## Build for Linux

darwin: $(DARWIN) $(DARWIN_ARM)## Build for Darwin (macOS)

$(WINDOWS): bin
	env GOOS=windows GOARCH=amd64 go build -i -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/service/main.go

$(LINUX): bin
	env GOOS=linux GOARCH=amd64 go build -i -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/service/main.go

$(LINUX_ARM): bin
	env GOOS=linux GOARCH=arm64 go build -i -v -o $(LINUX_ARM) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/service/main.go

$(DARWIN): bin
	env GOOS=darwin GOARCH=amd64 go build -i -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/service/main.go

$(DARWIN_ARM): bin
	env GOOS=darwin GOARCH=arm64 go build -i -v -o $(DARWIN_ARM) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/service/main.go

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)
windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -i -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/service/main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -i -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/service/main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -i -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/service/main.go

run-daemon:
	go run cmd/sevice/main.go
test:
	go test ./...

