package main

import (
	"log"
	"net/http"
	"os"

	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/auth"
	"github.com/techmexdev/socialnet/pkg/handler"
	"github.com/techmexdev/socialnet/pkg/storage/memo"
)

func main() {
	addr := os.Getenv("ADDRESS")
	sign := os.Getenv("JWT_SIGNATURE")

	if len(addr) == 0 || len(sign) == 0 {
		log.Fatal("ADDRESS and JWT_SIGNATURE env vars must be set")
	}

	usrStore := memo.NewUserStorage()
	usrSvc := socialnet.UserService{
		Store: usrStore, Auth: auth.New(usrStore),
	}
	postSvc := socialnet.PostService{
		Store: memo.NewPostStorage(),
	}
	router := handler.New(usrSvc, postSvc, handler.Options{
		Log: true, Address: addr, Signature: sign,
	})
	log.Println("Starting server at: " + addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
