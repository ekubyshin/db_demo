LOCAL_BIN := $(CURDIR)/bin

GOENV:=GOPRIVATE="gitlab.ae-rus.net/*"

GOLANGCI_BIN=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG=v1.61.0
$(GOLANGCI_BIN):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCAL_BIN) $(GOLANGCI_TAG)

GOOSE_BIN=$(LOCAL_BIN)/goose
$(GOOSE_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: db-create
db-create: $(GOOSE_BIN)
	goose -dir migrations create $(NAME) sql

.PHONY: db-up
db-up: $(GOOSE_BIN)
	goose -dir migrations up

.PHONY: lint
lint: $(GOLANGCI_BIN)
	$(GOLANGCI_BIN) run --fix ./...

.PHONY: test
test:
	go test -v -cover ./...