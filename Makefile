SHELL := bash
CONTAINER_NAME := gopi
VERSION := $(shell git describe --tags --always || echo latest)
PODMAN ?= false
ifeq "$(PODMAN)" "true"
DOCKER := podman
else
DOCKER := docker
endif
OUTPUT_DIRECTORY ?= examples/test

.PHONY: build
build: build-release build-test # build all targets

.PHONY: build-release
build-release: ## build release target
	$(DOCKER) build \
		--target release \
		-t $(CONTAINER_NAME):$(VERSION) \
		-t $(CONTAINER_NAME):latest \
		.

.PHONY: build-test
build-test: build-release ## build test target
	$(DOCKER) build --target test -t $(CONTAINER_NAME):test .

.PHONY: test
test: build-test ## run go tests
	$(DOCKER) run -it gopi:test go test /app/test/... -test.v

.PHONY: run
run: build ## run generate in the docker container
	$(DOCKER) run \
		-v $(shell pwd):/run \
		-it \
		gopi:$(VERSION) \
		/gopi generate -o /run/$(OUTPUT_DIRECTORY)

.PHONY: deps
deps: ## install local deps
	go mod download

test-local: ## run go tests locally
	go test ./test/... -test.v

build-local: ## run the go tool locally
	go build -o gopi ./cmd/go-pi/main.go
	chmod +x gopi

run-local: build-local ## output command to run locally
	@echo "To run just use the binary from build-local and it will include help output"
	@echo "./gopi --help"

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
