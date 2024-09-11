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

## lint/ts: lints typescript code
.PHONY: lint/ts
lint/ts:
	@yarn lint

## test: run all tests
.PHONY: test
test:
	npx jest --passWithNoTests

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	npx jest --passWithNoTests --coverage

## generate/client: generate typescript-fetch for service-b from its OpenAPI documentation
.PHONY: generate/client
generate/client:
	@! [ -x "$(command -v redocly)" ] || npm i -g @redocly/cli@latest
	@redocly bundle ./../service-b-go/docs/openapi/api.yaml -o ./docs/openapi/api-service-b.bundled.yaml
	@! [ -x "$(command -v openapi-generator-cli)" ] || npm install @openapitools/openapi-generator-cli -g
	openapi-generator-cli generate -i ./docs/openapi/api-service-b.bundled.yaml -g typescript-fetch -o ./clients/service-b --skip-validate-spec --additional-properties=modelPropertyNaming=original
