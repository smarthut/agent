# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

PACKAGE := github.com/smarthut/megad328/cmd/megad328
DOCKER_IMAGE := smarthut/megad328

GOEXE ?= go

.PHONY: vendor build test help
.DEFAULT_GOAL := help

vendor: ## Install deps and sync vendored dependencies
	@echo "Installing vendored dependencies"
	@${GOEXE} get -u github.com/golang/dep/cmd/dep
	@dep ensure

build: vendor ## Build megad328 binary
	@echo "Building megad328 binary"
	@CGO_ENABLED=0 GOOS=linux ${GOEXE} build -a -installsuffix cgo ${PACKAGE}

test: ## Run tests
	@${GOEXE} test

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
