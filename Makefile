APP_NAME=hopeline

.PHONY: build run clean

build:
	go build -o bin/$(APP_NAME) cmd/main.go

run:
	go run ./cmd/main.go

clean:
	rm -f bin/$(APP_NAME)
