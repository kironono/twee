.PHONY: all
all:

.PHONY: lint
lint:
	@go vet `go list ./...`

.PHONY: test
test:
	@go test -v -cover ./...
