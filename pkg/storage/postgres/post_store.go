package postgres

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/techmexdev/socialnet"
)

// PostStorage is a postgres user storage.
type PostStorage struct {
	*sqlx.DB
}

// NewPostStorage returns a postgres implementation of socialnet.PostStorage.
func NewPostStorage(dsn string) *PostStorage {
	return &PostStorage{sqlx.MustOpen("postgres", dsn)}
}

// Create adds a Post to the database.
func (db *PostStorage) Create(post socialnet.Post) (socialnet.Post, error) {
	q := `INSERT INTO posts(id, created_at, updated_at, users_id, title, body)
		VALUES ($1, $2, $3, $4, $5, $6)`

	id, err := uuid.NewV4()
	if err != nil {
		return socialnet.Post{}, err
	}

	post.ID = id.String()
	createdAt := time.Now().Format(time.RFC3339)

	userStore := UserStorage{db.DB}
	usr, err := userStore.Read(post.Author)
	if err != nil {
		return socialnet.Post{}, errors.New("could not find user with username: " + post.Author)
	}

	_, err = db.Exec(q, post.ID, createdAt, createdAt, usr.ID, post.Title, post.Body)
	if err != nil {
		return socialnet.Post{}, err
	}

	return post, nil
}

// Read retrieves a Post from the database.
func (db *PostStorage) Read(author, title string) (socialnet.Post, error) {
	q := "SELECT * FROM posts WHERE title = $1 AND users_id = (SELECT id FROM users WHERE username = $2)"
	var post socialnet.Post

	err := db.Get(&post, q, title, author)
	if err != nil {
		return socialnet.Post{}, err
	}

	return post, nil
}

// Update replaces a Post from the database.
func (db *PostStorage) Update(author, title string, post socialnet.Post) (socialnet.Post, error) {
	params, vals, args := getParamsValsArgsFromPost(post)
	q := "UPDATE posts SET (" + params + ") = (" + vals + ") WHERE title = '$1'" +
		"AND users_id = (SELECT id FROM users WHERE username = '$2')"
	_, err := db.Exec(q, args...)
	if err != nil {
		return socialnet.Post{}, err
	}

	return post, nil
}

// Delete removes a Post from the database.
func (db *PostStorage) Delete(author, title string) error {
	q := `DELETE FROM posts WHERE title = $1
		AND users_id = (SELECT id FROM users WHERE username = $2)`

	_, err := db.Exec(q, title, author)
	if err != nil {
		return err
	}

	return nil
}

// List retrieves all Posts from the database.
func (db *PostStorage) List() ([]socialnet.Post, error) {
	q := "SELECT * FROM posts"
	var pp []socialnet.Post

	err := db.Select(&pp, q)
	if err != nil {
		return []socialnet.Post{}, err
	}

	return pp, nil
}

func getParamsValsArgsFromPost(post socialnet.Post) (params, vals string, args []interface{}) {
	if post.Title != "" {
		params, vals, args = appendParamsAndArgs("title", post.Title, params, vals, args)
	}

	if post.Body != "" {
		params, vals, args = appendParamsAndArgs("body", post.Body, params, vals, args)
	}

	return params, vals, args
}
