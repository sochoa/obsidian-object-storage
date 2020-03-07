.DEFAULT_GOAL := default
BINARY := obsidian
PROJECT_DIR :=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
PROJECT_URL := github.com/sochoa/obsidian

LD_FLAGS := 
LD_FLAGS += -X '$(PROJECT_URL).version=$(shell git rev-parse --short HEAD)'
LD_FLAGS += -X '$(PROJECT_URL).binary=$(BINARY)'

default: clean build

clean:
	rm -f $(PROJECT_DIR)/build/$(BINARY)

build:
	mkdir -p $(PROJECT_DIR)/build
	go build -ldflags="$(LD_FLAGS)" $(PROJECT_DIR) -o $(BINARY)

mod_init:
	go mod init $(PROJECT_URL)


