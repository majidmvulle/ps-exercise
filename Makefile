.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


## build/service-a: builds service-a Docker image and loads it locally
.PHONE: build/service-a
build/service-a:
	docker buildx build --load --tag service-a -f Dockerfile.service-a .

## build/service-b: builds service-b Docker image (for amd64) and loads it locally
.PHONE: build/service-b-amd64
build/service-b-amd64:
	docker buildx build --load --platform linux/amd64 --tag service-b -f Dockerfile.service-b .

## build/service-b: builds service-b Docker image (for arm64) and loads it locally
.PHONE: build/service-b-arm64
build/service-b-arm64:
	docker buildx build --load --platform linux/arm64 --tag service-b -f Dockerfile.service-b .


## run/service-a: runs service-a Docker image
.PHONE: run/service-a
run/service-a:
	docker run --env-file ./service-a-expressjs/.env -p 3000:3000 service-a:latest

## run/service-b: runs service-b Docker image
.PHONE: run/service-b
run/service-b:
	docker run --env-file ./service-b-go/.env -p 8080:8080 service-b:latest
