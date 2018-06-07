package postgres_test

import (
	"testing"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"

	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/server/storage/postgres"
)

var dsn string

func init() {
	dsn = "postgres://socialnettest:socialnettest@localhost/socialnettest?sslmode=disable"
}

func TestUserStore(t *testing.T) {
	postgres.MigrateUp("file://migrations", dsn)
	defer postgres.MigrateDown("file://migrations", dsn)

	userStore := postgres.NewUserStorage(dsn)
	john := socialnet.User{
		Username: "jlennon", FirstName: "John", LastName: "Lennon", Email: "jlennon@beatles.com", Password: "Strawberryfields67!",
	}

	storedJohn, err := userStore.Create(john)
	if err != nil {
		t.Error("error creating user: ", err)
	}

	if storedJohn.Username != john.Username {
		t.Errorf("username should be stored. have %s, want %s", storedJohn.Username, john.Username)
	}

	storedJohn, err = userStore.Read(john.Username)
	if err != nil {
		t.Error("error reading user: ", err)
	}

	if storedJohn.Username != john.Username {
		t.Errorf("username should be stored. have %s, want %s", storedJohn.Username, john.Username)
	}

	if storedJohn.Password == john.Password {
		t.Errorf("user storage should hash passwords")
	}

	john.LastName = "Lemon"
	newJohn, err := userStore.Update(john.Username, john)
	if err != nil {
		t.Error("error updating user: ", err)
	}

	if newJohn.LastName != "Lemon" {
		t.Errorf("user not updated. have %#v, want %#v", newJohn, john)
	}

	err = userStore.Delete(john.Username)
	if err != nil {
		t.Error("error deleting user: ", err)
	}

	beatles, err := userStore.List()
	if err != nil {
		t.Error("error listing users: ", err)
	}

	if len(beatles) > 0 {
		t.Errorf("I DON'T BELIEVE, IN BEATLES. have %#v, want none", beatles)
	}
}

func TestUserPosts(t *testing.T) {
	postgres.MigrateUp("file://migrations", dsn)
	defer postgres.MigrateDown("file://migrations", dsn)

	userStore := postgres.NewUserStorage(dsn)
	paul := socialnet.User{
		Username: "pmccartney", FirstName: "Paul", LastName: "McCartney", Email: "pmccarney@beatles.com", Password: "heyyyjuuude123",
	}

	_, err := userStore.Create(paul)
	if err != nil {
		t.Error("error creating paul: ", err)
	}

	postStore := postgres.NewPostStorage(dsn)

	yesterday := socialnet.Post{
		Author: paul.Username, Title: "Yesterday", Body: "Yesterday. All my troubles seemed so far away...",
	}

	_, err = postStore.Create(yesterday)
	if err != nil {
		t.Error("error creating yesterday post: ", err)
	}

	storedPaul, err := userStore.Read(paul.Username)
	if err != nil {
		t.Error("error reading paul: ", err)
	}

	if len(storedPaul.Posts) == 0 || storedPaul.Posts[0].Body != yesterday.Body {
		t.Errorf("error retrieving paul's posts. have %#v, want %#v", storedPaul.Posts, []socialnet.Post{yesterday})
	}
}
