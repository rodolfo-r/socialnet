package postgres

import (
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/techmexdev/socialnet"
)

// LikeStore is an implementation of socialnet.LikeStorage.
type LikeStore struct {
	*sqlx.DB
}

// NewLikeStorage creates an in-memory socialnet.LikeStorage.
func NewLikeStorage(dsn string) *LikeStore {
	return &LikeStore{sqlx.MustOpen("postgres", dsn)}
}

// Create adds a like to the database.
func (db *LikeStore) Create(username, postID string) error {
	q := "SELECT id FROM users WHERE username = $1"
	var user socialnet.User

	err := db.Get(&user, q, username)
	if err != nil {
		return err
	}

	q = "INSERT INTO likes (id, post_id, liker_id) VALUES ($1, $2, $3)"

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	_, err = db.Exec(q, id, postID, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a users' like from a post.
func (db *LikeStore) Delete(username, postID string) error {
	q := `DELETE FROM likes WHERE post_id = $1 AND liker_id = (
		SELECT id FROM users WHERE username = $2)`

	_, err := db.Exec(q, postID, username)
	return err
}

// List retrieves all of a post's likes.
func (db *LikeStore) List(postID string) ([]socialnet.Like, error) {
	q := `SELECT username, image_url, first_name, last_name FROM users WHERE id IN (
		SELECT liker_id FROM likes WHERE post_id = $1)`

	var uu []socialnet.UserItem
	err := db.Select(&uu, q, postID)
	if err != nil {
		return []socialnet.Like{}, err
	}

	var likes []socialnet.Like
	for _, u := range uu {
		likes = append(likes, socialnet.Like{
			PostID: postID, UserItem: u,
		})
	}

	return likes, nil
}
