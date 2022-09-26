package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kapralovs/user-balance/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port := ":8080"
	server := server.New(port)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
