package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
	"github.com/techmexdev/socialnet"
)

// CommentStore is an implementation of socialnet.CommentStorage.
type CommentStore struct {
	*sqlx.DB
}

// NewCommentStorage creates an in-memory socialnet.CommentStorage.
func NewCommentStorage(dsn string) *CommentStore {
	return &CommentStore{sqlx.MustOpen("postgres", dsn)}
}

// Create adds a like to the cctore.
func (db *CommentStore) Create(username, postID, text string) error {
	q := "SELECT id FROM users WHERE username = $1"

	var user socialnet.User
	err := db.Get(&user, q, username)
	if err != nil {
		return err
	}

	q = "INSERT INTO comments(id, post_id, commenter_id, text) VALUES ($1, $2, $3, $4)"
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	_, err = db.Exec(q, id, postID, user.ID, text)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a users' like from a post.
func (db *CommentStore) Delete(username, postID string) error {
	q := `DELETE FROM comments WHERE post_id = $1 AND
		commenter_id = (SELECT id FROM users WHERE username = $2)`
	_, err := db.Exec(q, postID, username)
	return err
}

// List retrieves all of a post's cc.
func (db *CommentStore) List(postID string) ([]socialnet.Comment, error) {
	q := `SELECT comments.text, username, first_name, last_name, image_url
		FROM comments INNER JOIN users
		ON users.id = comments.commenter_id  
		WHERE comments.post_id = $1`

	var cc []socialnet.Comment
	err := db.Select(&cc, q, postID)
	if err != nil {
		return []socialnet.Comment{}, err
	}

	return cc, nil
}
