package storage_test

import (
	"log"
	"testing"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/server/storage/memo"
	"github.com/techmexdev/socialnet/pkg/server/storage/postgres"
)

var followStore socialnet.UserFollow
var follower, followee socialnet.User

func TestFollow(t *testing.T) {
	var userStore socialnet.UserStorage
	if testMemo {
		followStore = memo.NewUserFollow()
		userStore = memo.NewUserStorage()
	} else {
		postgres.MigrateDown("file://postgres/migrations", dsn)
		postgres.MigrateUp("file://postgres/migrations", dsn)

		followStore = postgres.NewUserFollow(dsn)
		userStore = postgres.NewUserStorage(dsn)
	}

	follower = socialnet.User{Username: "paul"}
	followee = socialnet.User{Username: "george"}

	_, err := userStore.Create(follower)
	if err != nil {
		log.Fatal(err)
	}

	_, err = userStore.Create(followee)
	if err != nil {
		log.Fatal(err)
	}
}

func TestFollowFollow(t *testing.T) {
	err := followStore.Follow(follower.Username, followee.Username)
	if err != nil {
		t.Error(err)
	}
}

func TestFollowFollowers(t *testing.T) {
	ff, err := followStore.Followers(followee.Username)
	if err != nil {
		t.Error(err)
	}

	if len(ff) != 1 || ff[0].Username != follower.Username {
		t.Errorf(
			"error reading followers. have %#v, want %#v",
			ff,
			[]socialnet.User{follower},
		)
	}
}

func TestFollowFollowing(t *testing.T) {
	ff, err := followStore.Following(follower.Username)
	if err != nil {
		t.Error(err)
	}

	if len(ff) != 1 || ff[0].Username != followee.Username {
		t.Errorf(
			"error reading following. have %#v, want %#v",
			ff,
			[]socialnet.User{followee},
		)
	}
}

func TestFollowUnfollow(t *testing.T) {
	err := followStore.Unfollow(follower.Username, followee.Username)
	if err != nil {
		t.Error(err)
	}

	ff, err := followStore.Following(follower.Username)
	if err != nil {
		t.Error(err)
	}

	if len(ff) > 0 {
		t.Errorf("error reading following. have %v, want %v", len(ff), 0)
	}
}
