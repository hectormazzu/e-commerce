APP_NAME=orders-service
PORT=8081

run:
	go run ./cmd/main.go

build:
	go build -o bin/$(APP_NAME) ./cmd/main.go

test:
	go test ./...

lint:
	golangci-lint run

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run --rm -p $(PORT):$(PORT) $(APP_NAME)

