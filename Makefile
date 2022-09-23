run:
	go run ./cmd/main.go
build:
	go build -o user-balance ./cmd/main.go
run-nothing:
	touch somefile.txt