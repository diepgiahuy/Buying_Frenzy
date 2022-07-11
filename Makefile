.PHONY: build run deploy test test-report swagger migrate-up migrate-down wire load-data
APP_NAME=buying-frenzy
POSTGRESQL_URL=postgresql://postgres:secret@localhost:5432/buying_frenzy?sslmode=disable
build:
	go build -v -o $(APP_NAME)
load-data:
	go run main.go load
run:
	go run main.go serve
deploy:
	docker-compose up	
test:
	go test -v -cover ./...
test-report:
	o test -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
swagger:
	swag init -g /pkg/api/gin.go
wire:
	wire ./cmd
migrate-up:
	migrate -path  migration  -database ${POSTGRESQL_URL} -verbose up
migrate-down:
	migrate -path  migration  -database ${POSTGRESQL_URL} -verbose down
all: build deploy test