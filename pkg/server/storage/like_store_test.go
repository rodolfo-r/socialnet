package storage_test

import (
	"log"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/rodolfo-r/socialnet"
	"github.com/rodolfo-r/socialnet/pkg/server/storage/memo"
	"github.com/rodolfo-r/socialnet/pkg/server/storage/postgres"
)

var likeStore socialnet.LikeStorage
var likePostStore socialnet.PostStorage
var likePost socialnet.Post
var liker, unliker socialnet.User

func TestLikeStore(t *testing.T) {
	if testMemo {
		likeStore = memo.NewLikeStorage()
		likePostStore = memo.NewPostStorage()
	} else {
		postgres.MigrateDown("file://postgres/migrations", dsn)
		postgres.MigrateUp("file://postgres/migrations", dsn)

		likeStore = postgres.NewLikeStorage(dsn)
		likePostStore = postgres.NewPostStorage(dsn)
		userStore = postgres.NewUserStorage(dsn)
	}

	poster := socialnet.User{Username: "ringo"}
	liker = socialnet.User{Username: "paul"}
	unliker = socialnet.User{Username: "john"}

	_, err := userStore.Create(liker)
	if err != nil {
		t.Fatal(err)
	}

	_, err = userStore.Create(unliker)
	if err != nil {
		t.Fatal(err)
	}

	_, err = userStore.Create(poster)
	if err != nil {
		t.Fatal(err)
	}

	p := socialnet.Post{
		CreatedAt: time.Now(), Author: poster.Username, Title: "Octopus's Garden",
		Body: "I'd like to be. Under the sea. In an post' garden. In the shade.",
	}

	likePost, err = likePostStore.Create(p)
	if err != nil {
		log.Fatal(err)
	}
}

func TestLikeCreate(t *testing.T) {
	err := likeStore.Create(liker.Username, likePost.ID)
	if err != nil {
		t.Error("error creating like: ", err)
	}

	err = likeStore.Create(unliker.Username, likePost.ID)
	if err != nil {
		t.Error("error creating like: ", err)
	}
}

func TestLikeDelete(t *testing.T) {
	err := likeStore.Delete(unliker.Username, likePost.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLikeList(t *testing.T) {
	likes, err := likeStore.List(likePost.ID)
	if err != nil {
		t.Fatal(err)
	}

	if len(likes) != 1 {
		t.Errorf("error listing likes. have %#v, want 1", likes)
	}

	if likes[0].Username != liker.Username {
		t.Errorf("error listing likes. have liker %s, want %s", likes[0].Username, liker.Username)
	}
}
