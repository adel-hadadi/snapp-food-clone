up-build:
	@docker compose down
	docker compose up -d --build

logs:
	docker compose logs -f

