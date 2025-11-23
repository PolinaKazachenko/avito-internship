LOCAL_BIN := $(CURDIR)/bin
GOLANGCI_LINT = $(CURDIR)/bin/golangci-lint

.PHONY: install-deps
install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: compose-up
compose-up:
	docker-compose -p pr-reviewer -f $(CURDIR)/build/docker-compose.yaml up

.PHONY: compose-stop
compose-stop:
	docker-compose -p pr-reviewer -f $(CURDIR)/build/docker-compose.yaml down

.PHONY: run-local
run-local:
	go run $(CURDIR)/cmd --env-path $(CURDIR)/.env

.PHONY: run-in-docker
run-in-docker:
	$(CURDIR)/main --env-path $(CURDIR)/.env

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run --config=.golangci.yaml