package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rodolfo-r/socialnet/pkg/client/handler"
)

func main() {
	port := os.Getenv("CLIENT_PORT")
	log.Println("Starting server at: port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, handler.New()))
}
