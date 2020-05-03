SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.DEFAULT_GOAL := help

TARGET=qn
BUILD_FLAGS=

GRAY=\033[1;90m
MAGENTA=\033[1;35m
RESET_COLOR=\033[0m

## build: build the application
build: clean fmt lint mod-tidy
	@echo -e "${GRAY}>> ⚙️\t${MAGENTA}Building...${RESET_COLOR}"
	@go build ${BUILD_FLAGS} -o ${TARGET} -v -trimpath .
.PHONY: build

## test: execute tests of all packages
test: lint
	@echo -e "${GRAY}>> 🧪\t${MAGENTA}Testing...${RESET_COLOR}"
	@go test -v -count=1 -race -trimpath ./...
.PHONY: test

## fmt: format all Go source files
fmt:
	@echo -e "${GRAY}>> ✏️\t${MAGENTA}Formatting...${RESET_COLOR}"
	@go fmt ./...
.PHONY: fmt

## lint: examine Go source code and report suspicious constructs
lint:
	@echo -e "${GRAY}>> 🔎\t${MAGENTA}Linting...${RESET_COLOR}"
	@go vet ./...
.PHONY: lint

## clean: clean the binary
clean:
	@echo -e "${GRAY}>> 🧹\t${MAGENTA}Cleaning...${RESET_COLOR}"
	@if [ -f ${TARGET} ] ; then rm -v ${TARGET} ; fi
.PHONY: clean

## run: run main.go
run: build
	@echo -e "${GRAY}>> 👟\t${MAGENTA}Running...${RESET_COLOR}"
	@go run -race -trimpath .
.PHONY: run

## mod-download: download modules to local cache
mod-download: mod-tidy
	@echo -e "${GRAY}>> ⬇️\t${MAGENTA}Downloading modules...${RESET_COLOR}"
	@go mod download
.PHONY: mod-download

## mod-update: update all modules
mod-update: mod-tidy
	@echo -e "${GRAY}>> ⬇️\t${MAGENTA}Updating modules...${RESET_COLOR}"
	@go get -u ./...
.PHONY: mod-update

## mod-tidy: Makes sure go.mod matches the source code in the module.
mod-tidy:
	@echo -e "${GRAY}>> ⬇️\t${MAGENTA}Tidying up modules...${RESET_COLOR}"
	@go mod tidy -v
.PHONY: mod-tidy

## help: prints this help message
help:
	@echo -e "Usage: make [target] ...\n"
	@echo -e "Targets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
.PHONY: help
