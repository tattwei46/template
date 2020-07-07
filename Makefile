PROJECT_NAME := "template"
TIME := $(shell date +'%F-%T.o')
TIMESTAMP := $(shell date +'%s')
GIT_COMMIT := $(shell git rev-parse --short HEAD)

CURR_DIR := $(shell pwd)

.PHONY: all build production-build help

all: build

build: ## Build the binary file
	@go build -o $(CURR_DIR)/bin/$(PROJECT_NAME).o

run: ## Build the binary file
	@go build -o $(CURR_DIR)/bin/$(PROJECT_NAME).o
	./bin/$(PROJECT_NAME).o
production-build: ## Build in the Production Mode for Linux
	@GOOS=linux GOARCH=amd64 go build  -o $(CURR_DIR)/bin/$(PROJECT_NAME)-$(TIME)
clean: ## Remove previous build
	@rm -f $(CURR_DIR)/bin/*.o