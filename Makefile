run:

	go run ./cmd/bot

build:

	go build -o bin/bot ./cmd/bot

tidy:

	go mod tidy

docker-up:
	docker compose up -d