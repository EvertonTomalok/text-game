tests:
	go test ./...

start-game:
	go run .

setup:
	go mod tidy

build:
	go build .

build-and-run:
	go build . && ./text-game
