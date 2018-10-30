package postgres

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rodolfo-r/socialnet"
	uuid "github.com/satori/go.uuid"
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
	q := `INSERT INTO posts(id, created_at, updated_at, users_id, title, body, image_url)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	id := uuid.NewV4()

	post.ID = id.String()
	createdAt := time.Now().Format(time.RFC3339)

	userStore := UserStorage{db.DB}
	usr, err := userStore.Read(post.Author)
	if err != nil {
		return socialnet.Post{}, errors.New("could not find user with username: " + post.Author + ": " + err.Error())
	}

	_, err = db.Exec(q, post.ID, createdAt, createdAt, usr.ID, post.Title, post.Body, post.ImageURL)
	if err != nil {
		return socialnet.Post{}, err
	}

	return post, nil
}

// Read retrieves a Post from the database.
func (db *PostStorage) Read(id string) (socialnet.Post, error) {
	q := `SELECT posts.id, posts.created_at, posts.updated_at, title, body, username, image_url	FROM posts
		INNER JOIN users
		ON posts.users_id = users.id
		WHERE posts.id = $1`
	var post socialnet.Post

	err := db.Get(&post, q, id)
	if err != nil {
		return socialnet.Post{}, err
	}

	return post, nil
}

// Update replaces a Post from the database.
func (db *PostStorage) Update(id string, post socialnet.Post) (socialnet.Post, error) {
	params, vals, args := getParamsValsArgsFromPost(post)
	q := "UPDATE posts SET " + params + " = " + vals + " WHERE id = $" + strconv.Itoa(len(args)+1)
	log.Println("update query: ", q)

	var interID interface{} = id
	args = append(args, interID)

	_, err := db.Exec(q, args...)
	if err != nil {
		return socialnet.Post{}, err
	}

	return post, nil
}

// Delete removes a Post from the database.
func (db *PostStorage) Delete(id string) error {
	q := "DELETE FROM posts WHERE id = $1"

	_, err := db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}

// List retrieves all Posts from the database.
func (db *PostStorage) List(username string) ([]socialnet.Post, error) {
	q := "SELECT * FROM posts WHERE users_id = (SELECT id FROM users WHERE username = $1)"
	var pp []socialnet.Post

	err := db.Select(&pp, q, username)
	if err != nil {
		return []socialnet.Post{}, err
	}

	for i := range pp {
		pp[i].Author = username
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

	if post.ImageURL != "" {
		params, vals, args = appendParamsAndArgs("image_url", post.ImageURL, params, vals, args)
	}

	if len(args) > 1 {
		params = "(" + params + ")"
		vals = "(" + vals + ")"
	}

	return params, vals, args
}
