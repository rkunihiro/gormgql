.PHONY: all tools clean lint fmt test build

all: clean lint fmt test build

tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	go install github.com/99designs/gqlgen@v0.17.20

codegen:
	gqlgen generate

clean:
	rm -rf ./out

lint:
	go vet ./...
	staticcheck ./...

fmt:
	find . -name '*.go' | xargs gofmt -w -l
	find . -name '*.go' | xargs goimports -w -l -local github.com/rkunihiro/gormgql

test:
	go test -v -cover ./...

build:
	go build -ldflags="-s -w" -o ./out/server ./server.go
