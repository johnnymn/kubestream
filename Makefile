.PHONY: vendor

default: test

release: generate
	goreleaser

snapshot: generate
	goreleaser --snapshot --skip-publish --rm-dist

build:
	go build -mod vendor

generate:
	pkger -include github.com/relingan/kubestream:/stacks/templates

vendor:
	go mod tidy && go mod download && go mod vendor

GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck'"

lint:
	golangci-lint run ./...

test: fmtcheck
	go test ./...
