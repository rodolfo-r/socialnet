package memo_test

import (
	"testing"
	"time"

	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/storage/memo"
)

func TestPostStore(t *testing.T) {
	postStore := memo.NewPostStorage()
	octopus := socialnet.Post{
		CreatedAt: time.Now(), Author: "rstarr", Title: "Octopus's Garden",
		Body: "I'd like to be. Under the sea. In an octopus' garden. In the shade.",
	}

	storedOcto, err := postStore.Create(octopus)
	if err != nil {
		t.Fatal(err)
	}

	if storedOcto.Title != octopus.Title {
		t.Fatalf("username should be stored. have %s, want %s", storedOcto.Title, octopus.Title)
	}

	storedOcto, err = postStore.Read("rstarr", octopus.Title)
	if err != nil {
		t.Fatal(err)
	}

	if storedOcto.Body != octopus.Body {
		t.Fatalf("post should be stored. have %#v, want %#v", storedOcto, octopus)
	}

	octopus.Body = "He'd let us in. Knows where we've been. In his octopus' garden. In the shade."
	newOcto, err := postStore.Update("rstarr", octopus.Title, octopus)
	if err != nil {
		t.Fatal(err)
	}

	if newOcto.Body != octopus.Body {
		t.Fatalf("post not updated. have %#v, want %#v", newOcto, octopus)
	}

	err = postStore.Delete("rstarr", octopus.Title)
	if err != nil {
		t.Fatal(err)
	}

	posts, err := postStore.List()
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) > 0 {
		t.Fatalf("should have deleted all posts. have %#v, want none", posts)
	}
}
