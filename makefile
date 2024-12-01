up-build:
	cd ./cmd/api && env GOOS=linux CGO_ENABLED=0 go build -o ../../webApp ./main.go
	docker compose down
	docker compose up -d --build

logs:
	docker compose logs -f

