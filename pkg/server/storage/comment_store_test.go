package storage_test

import (
	"log"
	"testing"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/rodolfo-r/socialnet"
	"github.com/rodolfo-r/socialnet/pkg/server/storage/memo"
	"github.com/rodolfo-r/socialnet/pkg/server/storage/postgres"
)

var commentStore socialnet.CommentStorage
var commentPost socialnet.Post
var comment, comment2 socialnet.Comment

func TestCommentStore(t *testing.T) {
	var userStore socialnet.UserStorage
	var postStore socialnet.PostStorage
	if testMemo {
		commentStore = memo.NewCommentStorage()
		userStore = memo.NewUserStorage()
		postStore = memo.NewPostStorage()
	} else {
		postgres.MigrateDown("file://postgres/migrations", dsn)
		postgres.MigrateUp("file://postgres/migrations", dsn)

		commentStore = postgres.NewCommentStorage(dsn)
		userStore = postgres.NewUserStorage(dsn)
		postStore = postgres.NewPostStorage(dsn)
	}

	commentPost = socialnet.Post{
		Author: "John", Title: "You never give me your money", Body: "You only give me your funny paper",
	}

	comment = socialnet.Comment{
		UserItem: socialnet.UserItem{Username: "Paul"},
		Text:     "And in the middle of negotiations...",
	}
	comment2 = socialnet.Comment{
		UserItem: socialnet.UserItem{Username: "Ringo"},
		Text:     "You break down",
	}

	_, err := userStore.Create(socialnet.User{Username: comment.Username})
	if err != nil {
		log.Fatal(err)
	}

	_, err = userStore.Create(socialnet.User{Username: comment2.Username})
	if err != nil {
		log.Fatal(err)
	}

	_, err = userStore.Create(socialnet.User{Username: commentPost.Author})
	if err != nil {
		log.Fatal(err)
	}

	commentPost, err = postStore.Create(commentPost)
	if err != nil {
		log.Fatal(err)
	}
}

func TestCommentStoreCreate(t *testing.T) {
	err := commentStore.Create(comment.Username, commentPost.ID, comment.Text)
	if err != nil {
		t.Error("error creating comment: ", err)
	}

	err = commentStore.Create(comment2.Username, commentPost.ID, comment2.Text)
	if err != nil {
		t.Error("error creating comment: ", err)
	}
}

func TestCommentStoreDelete(t *testing.T) {
	err := commentStore.Delete(comment2.Username, commentPost.ID)
	if err != nil {
		t.Error("error deleting comment: ", err)
	}
}

func TestCommentStoreList(t *testing.T) {
	cc, err := commentStore.List(commentPost.ID)
	if err != nil {
		t.Error("error listing comments: ", err)
	}

	if len(cc) != 1 {
		t.Fatalf("expected 1 comment but got %v", cc)
	}

	if cc[0].Username != comment.Username {
		t.Errorf("error comparing listed comment. have username %s, want %s", cc[0].Username, comment.Username)
	}
}
