# required environments
include .env

.PHONY: down
down:
	docker compose down

.PHONY: up-build
up-build: build-app build-cron
	docker compose down
	docker compose up -d --build
	@docker compose logs -f

.PHONY: build-app
build-app:
	@echo "Building The App Binary"
	env GOOS=linux CGO_ENABLED=0 go build -o ./webApp ./cmd/api/main.go

.PHONY: build-cron
build-cron:
	@echo "Building The CronJob Binary"
	env GOOS=linux CGO_ENABLED=0 go build -o ./cronApp ./cmd/cron/main.go

.PHONY: up-front
up-front:
	@cd ./web/ && npm run dev

.PHONY: logs
logs:
	docker compose logs -f

.PHONY: ps
ps:
	docker compose ps

.PHONY: migrate
migrate: ## run migrations on app database with docker
	@echo "running app migrations:"
	GOOSE_VERBOSE=true GOOSE_COMMAND="up" GOOSE_COMMAND_ARG="" docker compose run --rm migrate
	@echo "done!"

.PHONY: migration
migration: ## create a new migration for app database with docker
	GOOSE_VERBOSE=true GOOSE_COMMAND="create" GOOSE_COMMAND_ARG="$(filter-out $@,$(MAKECMDGOALS)) sql" docker compose run --user ${UID}:${GID} --rm migrate
%:
	@:

