package main

import (
	"log"

	"github.com/kapralovs/user-balance/internal/server"
)

func main() {
	port := ":8080"
	server := server.New(port)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
