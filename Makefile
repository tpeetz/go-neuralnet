# Go parameters
GOPATH=$(HOME)/go
GOBIN=$(HOME)/go/bin
GOCMD=go
GOGET=$(GOCMD) get
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOCOVER=$(GOCMD) tool cover
COBERTURA=$(GOBIN)/gocover-cobertura
GOLINT=$(GOBIN)/golint
GOVET=$(GOCMD) vet

.PHONY: all deps clean clean-bin

all: deps check

deps:
	GOPATH=$(GOPATH) $(GOGET) -u github.com/jstemmer/go-junit-report
	GOPATH=$(GOPATH) $(GOGET) -u github.com/t-yuki/gocover-cobertura
	GOPATH=$(GOPATH) $(GOGET) -u golang.org/x/lint/golint

install:
	@echo using $(GOPATH)
	GOPATH=$(GOPATH) $(GOCMD) install -v -ldflags="-X main.version=$(shell git describe --tags)" ./...

check: test coverage lint vet

test:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -coverprofile=cover.out ./...
	$(GOCOVER) -html=cover.out -o coverage.html
	$(COBERTURA) < cover.out > coverage.xml

lint:
	$(GOLINT) ./...

vet:
	$(GOVET) -v ./...

