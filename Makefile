.DEFAULT_GOAL := default
BINARY := obsidian
PROJECT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
PROJECT_URL := github.com/sochoa/$(BINARY)
PROJECT_BUILD_OUTPUT := $(PROJECT_DIR)/build
VERSION := $(shell git rev-parse --short HEAD)
OS_FAMILY := $(shell uname -s | tr '[[:upper:]]' '[[:lower:]]')

DOCKER_PROJECT_DIR := /go/src/$(PROJECT_URL)
DOCKER_PROJECT_SRC := $(DOCKER_PROJECT_DIR)/$(BINARY)
DOCKER_PROJECT_BUILD_OUTPUT := $(DOCKER_PROJECT_DIR)/build
DOCKER_VOLUMES := -v $(PROJECT_DIR):$(DOCKER_PROJECT_DIR)
DOCKER_WORK_DIR := -w $(DOCKER_PROJECT_DIR)
DOCKER_ENV := -e GOOS=$(OS_FAMILY) -e GOARCH=amd64 
DOCKER_OPTS := -it $(DOCKER_ENV) $(DOCKER_VOLUMES) $(DOCKER_WORK_DIR)
DOCKER_RUN_GOLANG := docker run $(DOCKER_OPTS) golang
DOCKER_RUN_GO := $(DOCKER_RUN_GOLANG) go
DOCKER_RUN_BASH := $(DOCKER_RUN_GOLANG) bash

LD_FLAGS := 
LD_FLAGS += -X '$(PROJECT_URL).version=$(VERSION)'
LD_FLAGS += -X '$(PROJECT_URL).binary=$(BINARY)'

default: clean compile

clean:
	rm -rf $(PROJECT_BUILD_OUTPUT)/*

mod-vendor:
	rm -rf "$(PROJECT_DIR)/vendor/*"
	$(DOCKER_RUN_GO) mod vendor

$(PROJECT_BUILD_OUTPUT):
	mkdir -p $(PROJECT_BUILD_OUTPUT)

compile: $(PROJECT_BUILD_OUTPUT) clean 
	mkdir -p $(PROJECT_BUILD_OUTPUT)
	$(DOCKER_RUN_GO) build                        \
		-mod=vendor                                 \
		-ldflags="$(LD_FLAGS)"                      \
		-o $(DOCKER_PROJECT_BUILD_OUTPUT)/$(BINARY) \
		$(DOCKER_PROJECT_SRC)
