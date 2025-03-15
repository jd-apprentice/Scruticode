all: lint test imports build start

### Development

start:
	./bin/ScrutiCode

dev:
	air -c .air.toml

fmt:
	@if [ -n "$$(go fmt ./src)" ]; then \
        echo "Code is not properly formatted"; \
        exit 1; \
    fi

fmt-fix:
	go fmt ./src

imports: path
	goimports -w ./src

lint: path
	golangci-lint run

lint-fix: path
	golangci-lint run --fix

clear:
	rm -rf ./bin

cache:
	go clean -modcache

## https://github.com/golang-standards/project-layout/issues/113#issuecomment-1336514449
build: clear fmt
	GOARCH=amd64 go build -o ./bin/InfobaeAPI ./src/main.go

build-arm: clear fmt
	GOARCH=arm64 go build -o ./bin/InfobaeAPI ./src/main.go

test:
	go test -v ./tests

path:
	@export PATH=$$PATH:$$HOME/go/bin;

pre-commit:
	pre-commit install