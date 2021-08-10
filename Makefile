SERVICE_NAME = $(shell basename "$(PWD)")

ROOT = $(shell pwd)
GO ?= go
OS = $(shell uname -s | tr A-Z a-z)
export GOBIN = ${ROOT}/bin

PATH := $(PATH):$(GOBIN)

LINT = ${GOBIN}/golangci-lint
LINT_DOWNLOAD = curl --progress-bar -SfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

MIGRATE = ${GOBIN}/migrate
MIGRATE_DOWNLOAD = (curl --progress-bar -fL -o $(MIGRATE).tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.$(OS)-amd64.tar.gz; tar -xzvf $(MIGRATE).tar.gz -C $(GOBIN); mv $(MIGRATE).$(OS)-amd64 $(MIGRATE); rm $(MIGRATE).tar.gz)
MIGRATE_CONFIG = -source file://internal/infrastructure/db/migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL}"

TPARSE = $(GOBIN)/tparse
TPARSE_DOWNLOAD = $(GO) get github.com/mfridman/tparse

COMPILEDEAMON = $(GOBIN)/CompileDaemon
COMPILEDEAMON_DOWNLOAD = $(GO) get github.com/githubnemo/CompileDaemon

SWAG = ${GOBIN}/swag
SWAG_DOWNLOAD = $(GO) get -u github.com/swaggo/swag/cmd/swag


.PHONY: help
help: ## Display this help message
	@ cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: release
release: ## Build production binary file
	@ $(GO) build -a -o ./bin/${SERVICE_NAME} ./cmd/...


.PHONY: build
build: ## Build development binary file
	@ $(GO) build -o ./bin/${SERVICE_NAME} ./cmd/...


.PHONY: run
run: ## run as development reload if code changes
	@ test -e $(COMPILEDEAMON) || $(COMPILEDEAMON_DOWNLOAD)
	@ $(COMPILEDEAMON) --build="make build" --command="$(GOBIN)/$(SERVICE_NAME)"


.PHONY: mod
mod: ## Get dependency packages
	@ $(GO) mod tidy


.PHONY: test
test: ## Run unit tests
	echo $(TPARSE)
	@ test -e $(TPARSE) || $(TPARSE_DOWNLOAD)
	@ $(GO) test -failfast -count=1 ./... -json -cover | $(TPARSE) -all -smallscreen


.PHONY: race
race: ## Run data race detector
	@ test -e $(TPARSE) || $(TPARSE_DOWNLOAD)
	@ $(GO) test -short -race ./... -json -cover | $(TPARSE) -all -smallscreen


.PHONY: coverage
coverage: ## check coverage test code of sample https://penkovski.com/post/gitlab-golang-test-coverage/
	@ $(GO) test ./... -coverprofile=coverage.out
	@ $(GO) tool cover -func=coverage.out
	@ $(GO) tool cover -html=coverage.out -o coverage.html;


.PHONY: lint
lint: ## Lint the files
	@ test -e $(LINT) || $(LINT_DOWNLOAD)
	@ $(LINT) version
	@ $(LINT) --timeout 10m run


.PHONY: migrate
migrate: ## base migrate
	@ test -e $(MIGRATE) || $(MIGRATE_DOWNLOAD)
	@ $(MIGRATE) --version


.PHONY: migrate-up
migrate-up:migrate ## Apply all up migrations
	@ $(MIGRATE) $(MIGRATE_CONFIG) up


.PHONY:	migrate-down
migrate-down:migrate ## Apply all down migrations
	@ $(MIGRATE) $(MIGRATE_CONFIG) down


.PHONY: migrate-drop
migrate-drop:migrate ## Apply all down migrations
	@ $(MIGRATE) $(MIGRATE_CONFIG) drop

.PHONY: docs
docs: ## Create/Update documents using swagger tool
	@ test -e  $(SWAG) || $(SWAG_DOWNLOAD)
	@ swag init -g ./cmd/main.go -o ./docs --parseDependency


.PHONY: docker-build
docker-build: ## Build docker-compose
	@ docker-compose build


.PHONY: docker-up
docker-up: ## Run with docker-compose auto reload
	@ docker-compose up -d


.PHONY: docker-down
docker-down: ## Stop docker-compose
	@ docker-compose down


.PHONY: docker-log
docker-log: ## Print docker log
	@ docker-compose logs --tail=300 -f