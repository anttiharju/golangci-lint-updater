SHELL := bash
.ONESHELL:
.SHELLFLAGS := -euo pipefail -c
MAKEFLAGS += --warn-undefined-variables

GOLANGCI_LINT_VERSION=$(shell cat .golangci.version)
GOLANGCI_LINT_INSTALL_DIR=$(shell go env GOPATH)/bin

PHONY: setup
setup: install-lint install-pre-commit-hook
	bun install

.PHONY: install-pre-commit-hook
install-pre-commit-hook:
	rm -f .git/hooks/pre-commit
	cp scripts/pre-commit.sh .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit

.PHONY: install-lint
install-lint:
	VERSION=$(GOLANGCI_LINT_VERSION) INSTALL_DIR=$(GOLANGCI_LINT_INSTALL_DIR) scripts/install-lint.sh

.PHONY: lint
lint:
	$(GOLANGCI_LINT_INSTALL_DIR)/golangci-lint run

.PHONY: lint-fix
lint-fix:
	$(GOLANGCI_LINT_INSTALL_DIR)/golangci-lint run --fix

.PHONY: shellcheck
shellcheck:
	scripts/shellcheck.sh

.PHONY: ci
ci: shellcheck lint
	bun run typecheck
