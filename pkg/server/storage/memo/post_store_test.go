package memo_test

import (
	"testing"
	"time"

	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/server/storage/memo"
)

func TestPostStore(t *testing.T) {
	postStore := memo.NewPostStorage()
	octopus := socialnet.Post{
		CreatedAt: time.Now(), Author: "rstarr", Title: "Octopus's Garden",
		Body: "I'd like to be. Under the sea. In an octopus' garden. In the shade.",
	}

	storedOcto, err := postStore.Create(octopus)
	if err != nil {
		t.Error(err)
	}

	if storedOcto.Title != octopus.Title {
		t.Errorf("username should be stored. have %s, want %s", storedOcto.Title, octopus.Title)
	}

	storedOcto, err = postStore.Read("rstarr", octopus.Title)
	if err != nil {
		t.Error(err)
	}

	if storedOcto.Body != octopus.Body {
		t.Errorf("post should be stored. have %#v, want %#v", storedOcto, octopus)
	}

	octopus.Body = "He'd let us in. Knows where we've been. In his octopus' garden. In the shade."
	newOcto, err := postStore.Update("rstarr", octopus.Title, octopus)
	if err != nil {
		t.Error(err)
	}

	if newOcto.Body != octopus.Body {
		t.Errorf("post not updated. have %#v, want %#v", newOcto, octopus)
	}

	err = postStore.Delete("rstarr", octopus.Title)
	if err != nil {
		t.Error(err)
	}

	posts, err := postStore.List()
	if err != nil {
		t.Error(err)
	}

	if len(posts) > 0 {
		t.Errorf("should have deleted all posts. have %#v, want none", posts)
	}
}
