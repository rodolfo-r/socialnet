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
func (db *PostStorage) Create(post socialnet.Post) (socialnet.Post, error) {
	db.posts = append(db.posts, post)
	return post, nil
}

// Read retrieves a Post from the in memory array in PostStorage.
func (db *PostStorage) Read(author, title string) (socialnet.Post, error) {
	for _, p := range db.posts {
		if p.Author == author && p.Title == title {
			return p, nil
		}
	}
	return socialnet.Post{}, fmt.Errorf(
		"could not find post titled: %s from %s",
		title,
		author,
	)
}

// Update replaces a Post from the in memory array in PostStorage.
func (db *PostStorage) Update(author, title string, post socialnet.Post) (socialnet.Post, error) {
	for i := range db.posts {
		if db.posts[i].Author == author && db.posts[i].Title == title {
			db.posts[i] = post
			return db.posts[i], nil
		}
	}
	return socialnet.Post{}, fmt.Errorf(
		"could not find post titled: %s from: %s",
		title,
		author,
	)
}

// Delete removes a Post from the in memory array in PostStorage.
func (db *PostStorage) Delete(author, title string) error {
	for i := range db.posts {
		if db.posts[i].Title == title && db.posts[i].Author == author {
			db.posts = append(db.posts[:i], db.posts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf(
		"could not find post titled: %s from: %s",
		title,
		author,
	)
}

// List retrieves all Posts from the in memory array in PostStorage.
func (db *PostStorage) List() ([]socialnet.Post, error) {
	return db.posts, nil
}
