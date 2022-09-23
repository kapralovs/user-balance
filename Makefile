run:
	go run ./cmd/main.go
build:
	go build -o user-balance ./cmd/main.go
compose-build:
	docker-compose build
compose-up:
	docker-compose up