GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: test

release:
	goreleaser

snapshot:
	goreleaser --snapshot --skip-publish --rm-dist

build:
	go build -mod vendor

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck'"

lint:
	golangci-lint run ./...

test: fmtcheck
	go test ./...
