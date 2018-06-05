package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/auth"
	"github.com/techmexdev/socialnet/pkg/handler"
	"github.com/techmexdev/socialnet/pkg/storage/postgres"
)

func main() {
	addr := os.Getenv("ADDRESS")
	sign := os.Getenv("JWT_SIGNATURE")
	dsn := os.Getenv("PG_DSN")

	if len(addr) == 0 || len(sign) == 0 {
		log.Fatal("PG_DSN, ADDRESS, and JWT_SIGNATURE env vars must be set")
	}

	postgres.MigrateUp("file://pkg/storage/postgres/migrations", dsn)

	usrStore := postgres.NewUserStorage(dsn)
	usrSvc := socialnet.UserService{
		Store: usrStore, Auth: auth.New(usrStore),
	}
	postSvc := socialnet.PostService{
		Store: postgres.NewPostStorage(dsn),
	}
	router := handler.New(usrSvc, postSvc, handler.Options{
		Log: true, Address: addr, Signature: sign,
	})
	log.Println("Starting server at: " + addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
