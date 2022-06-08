ROOT = $(shell pwd)
SERVICE_NAME = wallet_service
GO ?= go
export GOBIN = ${ROOT}/bin

LINT = ${GOBIN}/golangci-lint
LINT_DOWNLOAD = curl --progress-bar -SfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
SWAG = ${GOBIN}/swag
SWAG_DOWNLOAD = $(GO) get -u github.com/swaggo/swag/cmd/swag

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

.PHONY: docs
docs: ## Create/Update documents using swagger tool
	@ test -e  $(SWAG) || $(SWAG_DOWNLOAD)
	@ swag init -g ./cmd/main.go -o ./docs --parseDependency
