ifeq ($(GOPATH),)
	GOPATH := $(HOME)/go
endif

#### VARIABLES ####
SHELL := /bin/bash


# Go parameters
GOCMD = go
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOMETALINTER := $(BIN_DIR)/gometalinter
RICHGO := $(BIN_DIR)/richgo
CILINT := $(BIN_DIR)/golangci-lint
DEP := $(BIN_DIR)/dep
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -name '*test.go' )

#### RULES ####
#GNU meaning of all is to compile application and be the default. Prob shouldn't run
.PHONY: all
all: lint test

.PHONY: test
test: $(RICHGO)
	TZ="Australia/Brisbane" RICHGO_FORCE_COLOR=1 RICHGO_LOCAL=1 richgo test -v -cover

$(RICHGO):
	 go get -u github.com/kyoh86/richgo
	 go install github.com/kyoh86/richgo

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

$(CILINT):
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
lint: $(CILINT)
	golangci-lint run ./...

$(DEP):
	go get -u github.com/golang/dep/cmd/dep

.PHONY: dep
dep: $(DEP)
	dep ensure