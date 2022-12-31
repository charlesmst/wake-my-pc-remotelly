
build:
	mkdir -p bin/
	go build cmd/main.go -o bin/wakepc
run:
	go run cmd/main.go
test:
	go test ./...

