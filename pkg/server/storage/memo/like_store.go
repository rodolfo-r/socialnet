package memo

import (
	"github.com/techmexdev/socialnet"
)

// LikeStore is an implementation of socialnet.LikeStorage.
type LikeStore struct {
	likes []socialnet.Like
}

// NewLikeStorage creates an in-memory socialnet.LikeStorage.
func NewLikeStorage() *LikeStore {
	return &LikeStore{}
}

// Create adds a like to the likeStore.
func (ls *LikeStore) Create(username, postID string) error {
	ui := socialnet.UserItem{Username: username}
	ls.likes = append(ls.likes, socialnet.Like{PostID: postID, UserItem: ui})

	return nil
}

// Delete removes a users' like from a post.
func (ls *LikeStore) Delete(username, postID string) error {
	for i := range ls.likes {
		if ls.likes[i].PostID == postID && ls.likes[i].Username == username {
			ls.likes = append(ls.likes[:i], ls.likes[i+1:]...)
		}
	}
	return nil
}

// List retrieves all of a post's likes.
func (ls *LikeStore) List(postID string) ([]socialnet.Like, error) {
	var likes []socialnet.Like
	for _, l := range ls.likes {
		if l.PostID == postID {
			likes = append(likes, l)
		}
	}
	return likes, nil
}
