package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/rodolfo-r/socialnet"
	uuid "github.com/satori/go.uuid"
)

// UserFollow is a postgres follow storage.
type UserFollow struct {
	*sqlx.DB
}

// NewUserFollow returns a postgres socialnet.UserFollow.
func NewUserFollow(dsn string) *UserFollow {
	return &UserFollow{sqlx.MustOpen("postgres", dsn)}
}

// Follow adds a follow relationship to the UserFollow.
func (db *UserFollow) Follow(followerUsername, followeeUsername string) error {
	q := "INSERT INTO follows (id, follower_id, followee_id) VALUES ($1, $2, $3)"

	id := uuid.NewV4()

	userStore := UserStorage{db.DB}

	follower, err := userStore.Read(followerUsername)
	if err != nil {
		return err
	}

	followee, err := userStore.Read(followeeUsername)
	if err != nil {
		return err
	}

	_, err = db.Exec(q, id, follower.ID, followee.ID)
	if err != nil {
		return err
	}

	return nil
}

// Followers returns a user's followers
func (db *UserFollow) Followers(username string) ([]socialnet.UserItem, error) {
	q := `SELECT username, first_name, last_name, image_url FROM users WHERE id IN(
		SELECT follower_id FROM follows WHERE followee_id = (
		SELECT id FROM users WHERE username = $1))`

	var uu []socialnet.UserItem
	err := db.Select(&uu, q, username)
	if err != nil {
		return []socialnet.UserItem{}, err
	}

	return uu, nil
}

// Following returns a the user's a user is following
func (db *UserFollow) Following(username string) ([]socialnet.UserItem, error) {
	q := `SELECT username, first_name, last_name, image_url FROM users WHERE id IN(
		SELECT followee_id FROM follows WHERE follower_id = (
		SELECT id FROM users WHERE username = $1))`

	var uu []socialnet.UserItem
	err := db.Select(&uu, q, username)
	if err != nil {
		return []socialnet.UserItem{}, err
	}

	return uu, nil
}

// Unfollow removes a relationship from the UserFollow.
func (db *UserFollow) Unfollow(follower, followee string) error {
	q := `DELETE FROM follows WHERE follower_id = (
		SELECT id FROM users WHERE username = $1)
		AND followee_id =(SELECT id FROM users WHERE username = $2)`
	_, err := db.Exec(q, follower, followee)
	return err
}
