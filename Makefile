SHELL ?= /bin/bash
export IMAGEORG ?= tedris
export IMAGE ?= about-me
export VERSION ?= $(shell printf "`./tools/version`${VERSION_SUFFIX}")

# Blackbox files that need to be decrypted.
clear_files=$(shell blackbox_list_files)
encrypt_files=$(patsubst %,%.gpg,${clear_files})

.PHONY: all
all: reset deploy_local

# -------------------------[ General Tools ]-------------------------

.PHONY: help
help: ## List of available commands
	@echo "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\033[36m\1\\033[m:\2/' | column -c2 -t -s :)"

.PHONY: clear
clear: ${clear_files}

${clear_files}: ${encrypt_files}
	@blackbox_decrypt_all_files

.PHONY: decrypt
decrypt: ${clear_files} ## Decrypt all .gpg files registered in .blackbox/blackbox-files.txt

.PHONY: encrypt
encrypt: ${encrypt_files} ## Encrypt all files registered in .blackbox/blackbox-files.txt
	blackbox_edit_end $^

.PHONY: version
version: tools/version ## Automatically calculate the version based on the number of commits since the last change to VERSION
	@echo ${VERSION}

# ---------------------------[ Local App ]---------------------------

.PHONY: dev
dev: ## Run the live-reloading application
	docker-compose -f docker-compose.dev.yml up -d
	make -s dev-logs

.PHONY: dev-down
dev_down: ## Bring down the live-reloading application
	docker-compose -f docker-compose.dev.yml down

.PHONY: dev-logs
dev-logs: ## Connect to the logs of the live-reloading application
	docker-compose -f docker-compose.dev.yml logs -f api

# -----------------------------[ Build ]-----------------------------

.PHONY: build
build: ## Build and tag the docker container for the API
	@docker build -f container/api.Dockerfile -t ${IMAGEORG}/${IMAGE}-build:${VERSION} --target builder .
	@docker tag ${IMAGEORG}/${IMAGE}-build:${VERSION} ${IMAGEORG}/${IMAGE}-build:latest

.PHONY: build-integration
build-integration: ## Build the integration test Docker container
	@docker build -f container/integration.Dockerfile -t ${IMAGEORG}/${IMAGE}-integration:${VERSION} .
	@docker tag ${IMAGEORG}/${IMAGE}-integration:${VERSION} ${IMAGEORG}/${IMAGE}-integration:latest

# -----------------------------[ Test ]------------------------------

.PHONY: test
test: test-unit test-integration ## Run all tests

.PHONY: test-unit
test-unit: build ## Run unit tests
	@tests/test_unit

.PHONY: test-integration
test-integration: build-integration ## Run integration tests
	@tests/test_integration

# -----------------------------[ Publish ]---------------------------

.PHONY: finalize
finalize: test ## Build, test, and tag the docker container with the finalized tag (typically, the full docker registery will be tagged here)
	@docker build -f container/api.Dockerfile -t ${IMAGEORG}/${IMAGE}:${VERSION} .
	@docker tag ${IMAGEORG}/${IMAGE}:${VERSION} ${IMAGEORG}/${IMAGE}:latest

.PHONY: publish_only
publish_only: ## Push the tagged docker image to the docker registry
	@docker push ${IMAGEORG}/${IMAGE}:${VERSION}

.PHONY: publish
publish: finalize publish_only ## Finalize and publish the docker container

# -----------------------------[ Deploy ]----------------------------

.PHONY: deploy_local
deploy_local: decrypt ## Deploy the application to the local environment (go run)
	@deploy/local/deploy

.PHONY: deploy_kube
deploy_kube: publish ## Deploy the application to Kubernetes
	@deploy/kube/deploy

.PHONY: deploy_lambda
deploy_lambda: publish ## Deploy the application to AWS Lambda
	@deploy/lambda/deploy

# ----------------------------[ Release ]----------------------------
# TODO

# -----------------------------[ Other ] ----------------------------

.PHONY: db
db: ## Connect to the primary database
	 psql postgresql://admin:admin@localhost:5432/app

.PHONY: db-test
db-test: ## Connect to the test database
	 psql postgresql://admin:admin@localhost:5432/app_test

.PHONY: cache
cache: ## Connect to the cache
	 redis-cli

.PHONY: ent-install
ent-install: ## Install Ent code-generation module
	go get -d entgo.io/ent/cmd/ent

.PHONY: ent-gen
ent-gen: ## Generate Ent code
	go generate ./ent

.PHONY: ent-new
ent-new: ## Create a new Ent entity
	go run entgo.io/ent/cmd/ent init $(name)

.PHONY: up
up: ## Start the Docker containers
	docker-compose up -d
	sleep 3

.PHONY: reset
reset: ## Rebuild Docker containers to wipe all data
	docker-compose down
	make up