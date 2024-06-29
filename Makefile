setup:
	go install github.com/swaggo/swag/cmd/swag@latest
	~/go/bin/swag init -g cmd/server/main.go
	go build -o bin/server cmd/server/main.go

build:
	docker compose build

dev:
	go run github.com/air-verse/air

test:
	docker compose down
	docker compose up --build -d
	go test ./pkg/api/admin -v
	docker compose down

up:
	docker compose up

down:
	docker compose down

restart:
	docker compose restart

clean:
	docker stop go-rest-api-template
	docker stop dockerPostgres
	docker rm go-rest-api-template
	docker rm dockerPostgres
	docker rm dockerRedis
	docker image rm bookstore-api-go-backend
	rm -rf .dbdata
