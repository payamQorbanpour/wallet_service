ROOT = $(shell pwd)
SERVICE_NAME = wallet_service
GO ?= go
export GOBIN = ${ROOT}/bin

LINT = ${GOBIN}/golangci-lint
LINT_DOWNLOAD = curl --progress-bar -SfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

.PHONY: mod
mod: ## Get dependency packages
	@ $(GO) mod tidy

.PHONY: test
test: ## Run unit tests
	@ $(GO) test 

.PHONY: lint
lint: ## Lint the files
	@ test -e $(LINT) || $(LINT_DOWNLOAD)
	@ $(LINT) version
	@ $(LINT) --timeout 10m run

.PHONY: build
build: ## Build development binary file
	@ $(GO) build -race -ldflags '$(VERSION)' -o ./bin/${SERVICE_NAME} ./cmd/...

.PHONY: run
run: ## run as development reload if code changes
	@ $(GOBIN)/$(SERVICE_NAME)
