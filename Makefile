CURDIR = $(shell pwd)

.PHONY: test lint lintci

all: test lint

test:
	go test -count=1 -failfast -v ./...

lint:
	go mod tidy
	docker run --rm \
		-v $(shell go env GOPATH):/go \
		-v $(CURDIR):/app -w /app \
		-e GOLANGCI_ADDITIONAL_YML=/app/build/ci/.golangci.yml \
		quay.io/mittwald/golangci-lint:0.1.2 \
			golangci-lint run -v --fix ./...
