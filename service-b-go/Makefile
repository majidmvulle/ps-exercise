# Change these variables as necessary.
main_package_path = ./cmd/example
binary_name = example

include .env

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


## bundle/openapi: bundles openapi api files into one
.PHONE: bundle/openapi
bundle/openapi:
	redocly bundle docs/openapi/api.yaml -o docs/openapi/api.bundled.yaml

## lint/go: lints go code
.PHONY: lint/go
lint/go:
	@golangci-lint run --timeout=3m --issues-exit-code=0

## test: run all tests
.PHONY: test
test:
	go test -v -race ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	@mkdir -p .coverage && go test ./... -race -coverprofile=./.coverage/coverage.out -covermode=atomic -coverpkg=./...
	@go tool cover -html=./.coverage/coverage.out

## migration/up: runs migration up
.PHONY: migration/up
migration/up:
	@GOOSE_DRIVER=${DATABASE_DRIVER} GOOSE_DBSTRING=${DATABASE_DSN} goose up -dir=database/migrations

## migration/down: runs migration down
.PHONY: migration/down
migration/down:
	@GOOSE_DRIVER=${DATABASE_DRIVER} GOOSE_DBSTRING=${DATABASE_DSN} goose down -dir=database/migrations

## migration/status: runs migration status
.PHONY: migration/status
migration/status:
	@GOOSE_DRIVER=${DATABASE_DRIVER} GOOSE_DBSTRING=${DATABASE_DSN} goose status -dir=database/migrations

## seed/up: runs database seeding up
.PHONY: seed/up
seed/up:
	@GOOSE_DRIVER=${DATABASE_DRIVER} GOOSE_DBSTRING=${DATABASE_DSN} goose -no-versioning up -dir=database/seed

## seed/down: runs database seeding down
.PHONY: seed/down
seed/down:
	@GOOSE_DRIVER=${DATABASE_DRIVER} GOOSE_DBSTRING=${DATABASE_DSN} goose -no-versioning down -dir=database/seed


## generate/gin-server: generate code from OpenAPI documentation
.PHONY: generate/gin-server
generate/gin-server:
	openapi-generator-cli generate -g go-gin-server -o ./internal/api --skip-validate-spec -i ./docs/openapi/api.bundled.yaml --git-repo-id=ps-exercise/service-b-go --git-user-id=majidmvulle --additional-properties=withGoMod=false,generateInterfaces=true,packageName=api
## generate/openapi: generate code from OpenAPI documentation
.PHONY: generate/models
generate/models:
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	oapi-codegen --config=oapi.yaml ./docs/openapi/api.bundled.yaml

.PHONY: build
build:
    # Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	go build -o=/tmp/bin/${binary_name} ${main_package_path}
