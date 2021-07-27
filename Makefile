-include .makerc

.PHONY: help build run clean ci release publish

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: ## Build application
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/service/app cmd/service/main.go

run: ## Run application on local
	@go run cmd/service/main.go

docker-up: ## Run application with grafana on local
	@docker-compose up -d

docker-down: ## Down application with grafana on local
	@docker-compose down

clean:
	@rm -rf bin
