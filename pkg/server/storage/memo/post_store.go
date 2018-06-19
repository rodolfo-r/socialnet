package memo

import (
	"fmt"

	"github.com/techmexdev/socialnet"
)

// PostStorage is an in-memory user storage.
type PostStorage struct {
	posts []socialnet.Post
}

// NewPostStorage returns an in memory socialnet.PostStorage.
func NewPostStorage() *PostStorage {
	return &PostStorage{}
}

// Create adds a Post to the in memory array in PostStorage.
func (ps *PostStorage) Create(post socialnet.Post) (socialnet.Post, error) {
	ps.posts = append(ps.posts, post)
	return post, nil
}

// Read retrieves a Post from the in memory array in PostStorage.
func (ps *PostStorage) Read(id string) (socialnet.Post, error) {
	for _, p := range ps.posts {
		if p.ID == id {
			return p, nil
		}
	}
	return socialnet.Post{}, fmt.Errorf(
		"could not find post with ID %s",
		id,
	)
}

// Update replaces a Post from the in memory array in PostStorage.
func (ps *PostStorage) Update(id string, post socialnet.Post) (socialnet.Post, error) {
	for i := range ps.posts {
		if ps.posts[i].ID == id {
			ps.posts[i] = post
			return ps.posts[i], nil
		}
	}
	return socialnet.Post{}, fmt.Errorf(
		"could not find post with ID %s",
		id,
	)
}

// Delete removes a Post from the in memory array in PostStorage.
func (ps *PostStorage) Delete(id string) error {
	for i := range ps.posts {
		if ps.posts[i].ID == id {
			ps.posts = append(ps.posts[:i], ps.posts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf(
		"could not find post with ID %s",
		id,
	)
}

// List retrieves a user's posts from the in memory array in PostStorage.
func (ps *PostStorage) List(username string) ([]socialnet.Post, error) {
	var pp []socialnet.Post

	for _, p := range ps.posts {
		if p.Author == username {
			pp = append(pp, p)
		}
	}

	return pp, nil
}
