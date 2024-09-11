main_package_path = ./cmd/server
binary_name = service-b

include .env

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


## bundle/openapi: bundles openapi api files into one
.PHONE: bundle/openapi
bundle/openapi:
	@! [ -x "$(command -v redocly)" ] || npm i -g @redocly/cli@latest
	redocly bundle docs/openapi/api.yaml -o docs/openapi/api.bundled.yaml

## lint/openapi: lints openapi api files
.PHONE: lint/openapi
lint/openapi:
	@! [ -x "$(command -v redocly)" ] || npm i -g @redocly/cli@latest
	redocly lint docs/openapi/api.yaml

## lint/go: lints go code
.PHONY: lint/go
lint/go:
	@! [ -x "$(command -v golangci-lint)" ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.0
	@golangci-lint run --fix

## test: run all tests
.PHONY: test
test:
	go test -v -race ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	@mkdir -p .coverage && go test ./... -race -coverprofile=./.coverage/coverage.out -covermode=atomic -coverpkg=./...
	@go tool cover -html=./.coverage/coverage.out


## install/goose: installs goose for go migrations
.PHONE: install/goose
install/goose:
	@! [ -x "$(command -v goose)" ] || curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    sh


## migration/up: runs migration up
.PHONY: migration/up
migration/up:
	@make install/goose
	@GOOSE_DRIVER=${DATABASE_DRIVER} GOOSE_DBSTRING=${DATABASE_DSN} goose up -dir=database/migrations

## migration/down: runs migration down
.PHONY: migration/down
migration/down:
	@make install/goose
	@GOOSE_DRIVER=${DATABASE_DRIVER} GOOSE_DBSTRING=${DATABASE_DSN} goose down -dir=database/migrations

## migration/status: runs migration status
.PHONY: migration/status
migration/status:
	@make install/goose
	@GOOSE_DRIVER=${DATABASE_DRIVER} GOOSE_DBSTRING=${DATABASE_DSN} goose status -dir=database/migrations

## seed/up: runs database seeding up
.PHONY: seed/up
seed/up:
	@make install/goose
	@GOOSE_DRIVER=${DATABASE_DRIVER} GOOSE_DBSTRING=${DATABASE_DSN} goose -no-versioning up -dir=database/seed

## seed/down: runs database seeding down
.PHONY: seed/down
seed/down:
	@make install/goose
	@GOOSE_DRIVER=${DATABASE_DRIVER} GOOSE_DBSTRING=${DATABASE_DSN} goose -no-versioning down -dir=database/seed

## generate/models: generate Go models from OpenAPI documentation
.PHONY: generate/models
generate/models:
	@! [ -x "$(command -v oapi-codegen)" ] || go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	oapi-codegen --config=oapi.yaml ./docs/openapi/api.bundled.yaml

.PHONY: build
build:
    # Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	go build -o=/tmp/bin/${binary_name} ${main_package_path}
