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
	postgres.MigrateUp("file://migrations", dsn)
}

func TestUserFollow(t *testing.T) {
	defer postgres.MigrateDown("file://migrations", dsn)

	ringo := socialnet.User{Username: "ringo"}
	paul := socialnet.User{Username: "paul"}
	george := socialnet.User{Username: "george"}

	createUsers(t, ringo, paul, george)

	followStore := postgres.NewUserFollow(dsn)

	err := followStore.Follow(george.Username, ringo.Username)
	if err != nil {
		t.Error(err)
	}

	err = followStore.Follow(paul.Username, ringo.Username)
	if err != nil {
		t.Error(err)
	}

	ringoFols, err := followStore.Followers(ringo.Username)
	if err != nil {
		t.Error(err)
	}

	var foundGeorge, foundPaul bool
	for _, u := range ringoFols {
		if u.Username == george.Username {
			foundGeorge = true
		} else if u.Username == paul.Username {
			foundPaul = true
		}
	}

	if !foundGeorge || !foundPaul || len(ringoFols) != 2 {
		t.Errorf(
			"error reading ringo's followers. have %#v, want %#v",
			ringoFols,
			[]socialnet.User{george, paul},
		)
	}

	georgeFollowing, err := followStore.Following(george.Username)
	if err != nil {
		t.Error(err)
	}

	if len(georgeFollowing) != 1 || georgeFollowing[0].Username != ringo.Username {
		t.Errorf(
			"error reading george's followees. have %#v, want %#v",
			georgeFollowing,
			[]socialnet.User{ringo},
		)
	}

	err = followStore.Unfollow(paul.Username, ringo.Username)
	if err != nil {
		t.Error(err)
	}

	paulFollowing, err := followStore.Following(paul.Username)
	if err != nil {
		t.Error(err)
	}

	if len(paulFollowing) > 0 {
		t.Errorf(
			"error reading george's followees. have %#v, want %#v",
			paulFollowing,
			[]socialnet.User{},
		)
	}
}

func createUsers(t *testing.T, uu ...socialnet.User) {
	userStore := postgres.NewUserStorage(dsn)

	for _, u := range uu {
		_, err := userStore.Create(u)
		if err != nil {
			t.Fatal("error creating beatles: ", err)
		}
	}
}