package postgres_test

import (
	"log"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/server/storage/postgres"
)

var dsn string

func init() {
	dsn = "postgres://socialnettest:socialnettest@localhost/socialnettest?sslmode=disable"
	postgres.MigrateUp("file://migrations", dsn)

	ringo := socialnet.User{Username: "rstarr"}
	userStore := postgres.NewUserStorage(dsn)
	_, err := userStore.Create(ringo)
	if err != nil {
		log.Fatal("could not create user ringo: ", err)
	}
}

func TestPostStore(t *testing.T) {
	defer postgres.MigrateDown("file://migrations", dsn)

	postStore := postgres.NewPostStorage(dsn)
	octopus := socialnet.Post{
		CreatedAt: time.Now(), Author: "rstarr", Title: "Octopus's Garden",
		Body: "I'd like to be. Under the sea. In an octopus' garden. In the shade.",
	}

	storedOcto, err := postStore.Create(octopus)
	if err != nil {
		t.Errorf("could not create %v: %s", octopus, err)
	}

	if storedOcto.Title != octopus.Title {
		t.Errorf("username should be stored. have %s, want %s", storedOcto.Title, octopus.Title)
	}

	storedOcto, err = postStore.Read("rstarr", octopus.Title)
	if err != nil {
		t.Errorf("could not read %s from %s: %s", octopus.Title, "rstarr", err)
	}

	if storedOcto.Body != octopus.Body {
		t.Errorf("post should be stored. have %#v, want %#v", storedOcto, octopus)
	}

	octopus.Body = "He'd let us in. Knows where we've been. In his octopus' garden. In the shade."
	newOcto, err := postStore.Update("rstarr", octopus.Title, octopus)
	if err != nil {
		t.Errorf("could not update %s from %s: %s", octopus.Title, "rstarr", err)
	}

	if newOcto.Body != octopus.Body {
		t.Errorf("post not updated. have %#v, want %#v", newOcto, octopus)
	}

	err = postStore.Delete("rstarr", octopus.Title)
	if err != nil {
		t.Errorf("could not delete %s from %s: %s", octopus.Title, "rstarr", err)
	}

	posts, err := postStore.List()
	if err != nil {
		t.Error("could not list all posts: ", err)
	}

	if len(posts) > 0 {
		t.Errorf("should have deleted all posts. have %#v, want none", posts)
	}
}

