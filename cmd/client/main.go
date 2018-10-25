package main

import (
	"log"
	"net/http"

	"github.com/rodolfo-r/socialnet/pkg/client/handler"
)

func main() {
	log.Println("Starting server at: localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", handler.New()))
}
