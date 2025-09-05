hello:
	echo "Hello"

build:
	go build -o bin/caching-proxy cmd/main.go

run:
	go run cmd/main.go

compile:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/caching-proxy-linux-amd64 cmd/main.go

all: hello build
