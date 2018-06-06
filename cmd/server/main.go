package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/server/auth"
	"github.com/techmexdev/socialnet/pkg/server/handler"
	"github.com/techmexdev/socialnet/pkg/server/storage/postgres"
)

func main() {
	addr := "localhost:3001"
	dsn := os.Getenv("PG_DSN")
	sign := os.Getenv("JWT_SIGNATURE")

	if len(dsn) == 0 || len(sign) == 0 {
		log.Fatal("PG_DSN, and JWT_SIGNATURE env vars must be set")
	}

	postgres.MigrateUp("file://pkg/server/storage/postgres/migrations", dsn)

	usrStore := postgres.NewUserStorage(dsn)
	usrSvc := socialnet.UserService{
		Store: usrStore, Auth: auth.New(usrStore, addr, sign),
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
