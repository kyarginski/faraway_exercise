GO=go
GO111MODULE := auto
export GO111MODULE

lint:
	golangci-lint run ./...

test:
	$(GO) test -count=1 -race ./...

build: test
	$(GO) build -mod vendor -ldflags="-w -s" -o server faraway/cmd/server

buildstatic:
	$(GO) build -tags musl -ldflags="-w -extldflags '-static' -X 'main.Version=$(VERSION)'" -o server faraway/cmd/server
