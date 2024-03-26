.DEFAULT_GOAL := run

run:
	docker-compose up -d --build

down:
	docker-compose down

build:
	go mod download && \
	go build -o ./.bin/app ./cmd/app/main.go && \
	go build -o ./.bin/publisher ./cmd/publisher/main.go
