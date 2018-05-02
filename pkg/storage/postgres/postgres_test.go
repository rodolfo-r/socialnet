package postgres_test

import (
	"log"
	"testing"

	"github.com/golang-migrate/migrate"
	migpg "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/lib/pq"
	"github.com/techmexdev/the_social_network/pkg/model"
	"github.com/techmexdev/the_social_network/pkg/storage/postgres"
)

var db *postgres.Postgres

func init() {
	db = postgres.New("postgres://the_social_network_test:n3tw0rk50ci4l@localhost/tsnw_test?sslmode=disable")
	driver, err := migpg.WithInstance(db.DB, &migpg.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	m.Up()
}

func TestGetUser(t *testing.T) {
	_, err := db.CreateUser(model.User{}, "password123")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateUser(t *testing.T) {

}

func TestValidateUserCreds(t *testing.T) {

}

func TestGetProfile(t *testing.T) {

}

func TestCreatePost(t *testing.T) {

}

func TestGetUserSettings(t *testing.T) {

}

/*
	GetUser(model.User) (model.User, error)
	CreateUser(model.User) (model.User, error)
	ValidateUserCreds(username string, password string) error
	GetProfile(username string) (model.Profile, error)
	CreatePost(model.Post) (model.Post, error)
	GetUserSettings(username string) (model.Settings, error)
*/
