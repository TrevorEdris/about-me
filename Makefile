# Blackbox files that need to be decrypted.
clear_files=$(shell blackbox_list_files)
encrypt_files=$(patsubst %,%.gpg,${clear_files})

.PHONY: all
all: reset run

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

.PHONY: run
run: decrypt ## Run the application
	clear
	go run main.go

.PHONY: deploy_local
deploy_local: decrypt ## Deploy the application to the local environment (go run)
	@deploy/local/deploy

.PHONY: deploy_kube
deploy_kube: decrypt ## Deploy the application to Kubernetes
	@deploy/kube/deploy

.PHONY: deploy_lambda
deploy_lambda: decrypt ## Deploy the application to AWS Lambda
	@deploy/lambda/deploy

.PHONY: test
test: ## Run all tests
	go test -p 1 ./...