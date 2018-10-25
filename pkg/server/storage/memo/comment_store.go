package memo

import (
	"github.com/rodolfo-r/socialnet"
)

// CommentStore is an implementation of socialnet.CommentStorage.
type CommentStore struct {
	comments []socialnet.Comment
}

// NewCommentStorage creates an in-memory socialnet.CommentStorage.
func NewCommentStorage() *CommentStore {
	return &CommentStore{}
}

// Create adds a post to the CommentStore.
func (cs *CommentStore) Create(username, postID, text string) error {
	ui := socialnet.UserItem{Username: username}
	c := socialnet.Comment{PostID: postID, UserItem: ui, Text: text}
	cs.comments = append(cs.comments, c)

	return nil
}

// Delete removes a users' comment from a post.
func (cs *CommentStore) Delete(username, postID string) error {
	for i := 0; i < len(cs.comments); i++ {
		if cs.comments[i].PostID == postID && cs.comments[i].Username == username {
			if i+1 >= len(cs.comments) {
				cs.comments = cs.comments[:i]
			} else {
				cs.comments = append(cs.comments[:i], cs.comments[i+1:]...)
			}
			i--
		}
	}
	return nil
}

// List retrieves all of a post's cc.
func (cs *CommentStore) List(postID string) ([]socialnet.Comment, error) {
	var cc []socialnet.Comment
	for _, l := range cs.comments {
		if l.PostID == postID {
			cc = append(cc, l)
		}
	}
	return cc, nil
}
