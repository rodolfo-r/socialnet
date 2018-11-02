package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/rodolfo-r/socialnet"
	"github.com/rodolfo-r/socialnet/pkg/server/auth"
	"github.com/rodolfo-r/socialnet/pkg/server/handler"
	"github.com/rodolfo-r/socialnet/pkg/server/storage/postgres"
)

func main() {
	dsn := os.Getenv("PG_DSN")
	sign := os.Getenv("JWT_SIGNATURE")
	port := os.Getenv("SERVER_PORT")
	addr := os.Getenv("SERVER_ADDRESS")

	if len(dsn) == 0 || len(sign) == 0 || len(port) == 0 {
		log.Fatal("PG_DSN, SERVER_PORT, and JWT_SIGNATURE env vars must be set")
	}

	postgres.MigrateUp("file://pkg/server/storage/postgres/migrations", dsn)

	userStore := postgres.NewUserStorage(dsn)
	userFollow := postgres.NewUserFollow(dsn)

	usrSvc := socialnet.UserService{
		Store: userStore, Auth: auth.New(userStore, addr, sign), Follow: userFollow,
	}
	postSvc := socialnet.PostService{
		Store: postgres.NewPostStorage(dsn), Like: postgres.NewLikeStorage(dsn), Comment: postgres.NewCommentStorage(dsn),
	}
	router := handler.New(usrSvc, postSvc, handler.Options{
		Log: true, Signature: sign, Address: addr,
	})
	log.Println("Starting server at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
