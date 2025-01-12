# required environments
include .env

down:
	docker compose down

up-build:
	docker compose down
	docker compose up -d --build
	@docker compose logs -f

up-front:
	@cd ./web/ && npm run dev

logs:
	docker compose logs -f

ps:
	docker compose ps


migrate: ## run migrations on app database with docker
	@echo "running app migrations:"
	GOOSE_VERBOSE=true GOOSE_COMMAND="up" GOOSE_COMMAND_ARG="" docker compose run --rm migrate
	@echo "done!"

migration: ## create a new migration for app database with docker
	GOOSE_VERBOSE=true GOOSE_COMMAND="create" GOOSE_COMMAND_ARG="$(filter-out $@,$(MAKECMDGOALS)) sql" docker compose run --user ${UID}:${GID} --rm migrate
%:
	@:

