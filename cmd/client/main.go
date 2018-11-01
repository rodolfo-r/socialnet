package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rodolfo-r/socialnet/pkg/client/handler"
)

func main() {
	port := os.Getenv("CLIENT_PORT")
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if len(port) == 0 || len(serverAddress) == 0 {
		log.Fatal("env vars CLIENT_PORT and SERVER_ADDRESS must be set")
	}

	log.Println("Starting server at: port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, handler.New(serverAddress)))
}
