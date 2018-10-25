package storage_test

import (
	"errors"
	"fmt"
	"log"
	"testing"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"

	"github.com/rodolfo-r/socialnet"
	"github.com/rodolfo-r/socialnet/pkg/server/storage/memo"
	"github.com/rodolfo-r/socialnet/pkg/server/storage/postgres"
)

var postStore socialnet.PostStorage
var post socialnet.Post

func TestPostStore(t *testing.T) {
	var userStore socialnet.UserStorage
	if testMemo {
		postStore = memo.NewPostStorage()
		userStore = memo.NewUserStorage()
	} else {
		postgres.MigrateDown("file://postgres/migrations", dsn)
		postgres.MigrateUp("file://postgres/migrations", dsn)
		postStore = postgres.NewPostStorage(dsn)
		userStore = postgres.NewUserStorage(dsn)
	}

	user := socialnet.User{Username: "ringo"}

	_, err := userStore.Create(user)
	if err != nil {
		log.Fatal(err)
	}
	post = socialnet.Post{
		Author: user.Username, Title: "Octopus's Garden",
		Body: "I'd like to be. Under the sea. In an post' garden. In the shade.",
	}
}

func TestPostStoreCreate(t *testing.T) {
	p, err := postStore.Create(post)
	if err != nil {
		t.Error("error creating post: ", err)
	}

	if err := validatePost(p, post); err != nil {
		t.Errorf("error validating post: %s", err)
	}
	post.ID = p.ID
}

func TestPostStoreRead(t *testing.T) {
	p, err := postStore.Read(post.ID)
	if err != nil {
		t.Error(err)
	}

	if err := validatePost(p, post); err != nil {
		t.Errorf("error validating post: %s", err)
	}
}

func TestPostStoreUpdate(t *testing.T) {
	post.Body = "He'd let us in. Knows where we've been. In his post' garden. In the shade."
	p, err := postStore.Update(post.ID, post)
	if err != nil {
		t.Error(err)
	}

	if err := validatePost(p, post); err != nil {
		t.Errorf("error validating post: %s", err)
	}
}

func TestPostStoreDelete(t *testing.T) {
	err := postStore.Delete(post.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestPostStoreList(t *testing.T) {
	pp, err := postStore.List(post.Author)
	if err != nil {
		t.Error(err)
	}

	if len(pp) > 0 {
		t.Errorf("should have deleted all posts. have %#v, want none", pp)
	}

	_, err = postStore.Create(post)
	if err != nil {
		t.Fatal("could not read create post: ", err)
	}

	pp, err = postStore.List(post.Author)
	if err != nil {
		t.Fatal("error listing ringo's posts: ", err)
	}

	if len(pp) != 1 {
		t.Fatalf("expected %s to have one post", post.Author)
	}

	if err := validatePost(pp[0], post); err != nil {
		t.Errorf("error validating post: %s", err)
	}

}

func validatePost(p, validP socialnet.Post) error {
	var errMsg string

	if p.Body != validP.Body {
		errMsg = fmt.Sprintf("%s. have body %s want %s", errMsg, p.Body, validP.Body)
	}

	if p.Title != validP.Title {
		errMsg = fmt.Sprintf("%s. have title %s want %s", errMsg, p.Title, validP.Title)
	}

	if p.Author != validP.Author {
		errMsg = fmt.Sprintf("%s. have author %s want %s", errMsg, p.Author, validP.Author)
	}

	if len(errMsg) != 0 {
		return errors.New(errMsg)
	}

	return nil
}
