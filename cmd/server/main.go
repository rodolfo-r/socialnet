package main

import (
	"log"
	"net/http"
	"os"

	"github.com/techmexdev/the_social_network/pkg/handler"
	"github.com/techmexdev/the_social_network/pkg/storage/mock"
)

func main() {
	addr := os.Getenv("ADDRESS")
	sign := os.Getenv("JWT_SIGNATURE")
	if len(addr) == 0 || len(sign) == 0 {
		log.Fatal("ADDRESS and JWT_SIGNATURE env vars must be set")
	}

	r := handler.New(mock.New(), handler.Options{
		Log: true, Address: addr, Signature: sign,
	})
	log.Println("Starting server at: " + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
