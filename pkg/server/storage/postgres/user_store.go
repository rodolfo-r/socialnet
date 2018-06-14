package postgres

import (
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/techmexdev/socialnet"
)

// UserStorage is an in-memory socialnet.UserStorage.
type UserStorage struct {
	*sqlx.DB
}

// NewUserStorage returns an in-memory socialnet.UserStorage.
func NewUserStorage(dsn string) *UserStorage {
	return &UserStorage{sqlx.MustOpen("postgres", dsn)}
}

// Create adds a User to the in memory array in UserStorage.
func (db *UserStorage) Create(usr socialnet.User) (socialnet.User, error) {
	q := "INSERT INTO users(id, username, image_url, email, password, first_name, last_name, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	id, err := uuid.NewV4()
	if err != nil {
		return socialnet.User{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 12)
	if err != nil {
		return socialnet.User{}, err
	}

	usr.Password = string(hash)

	createdAt := time.Now().Format(time.RFC3339)

	_, err = db.Exec(q, id, usr.Username, usr.ImageURL, usr.Email, usr.Password, usr.FirstName, usr.LastName, createdAt, createdAt)
	if err != nil {
		return socialnet.User{}, err
	}

	return usr, nil
}

// Read retrieves a User from the in memory array in UserStorage.
func (db *UserStorage) Read(username string) (socialnet.User, error) {
	q := "SELECT * FROM users WHERE username = $1"

	var usr socialnet.User

	err := db.Get(&usr, q, username)
	if err != nil {
		return socialnet.User{}, err
	}

	q = "SELECT * FROM posts WHERE users_id = $1"

	var pp []socialnet.Post

	err = db.Select(&pp, q, usr.ID)
	if err != nil {
		return socialnet.User{}, err
	}

	for i := range pp {
		pp[i].Author = username
	}

	usr.Posts = pp

	return usr, nil
}

// Update uses the non-nil values from the usr to replace the values in the database.
func (db *UserStorage) Update(username string, usr socialnet.User) (socialnet.User, error) {
	params, vals, args := getParamsValsArgsFromUser(usr)
	q := "UPDATE users SET " + params + " = " + vals + " WHERE username = $" + strconv.Itoa(len(args)+1)
	_, err := db.Exec(q, append(args, username)...)
	if err != nil {
		return socialnet.User{}, err
	}

	return usr, nil
}

// Delete removes a User from the in memory array in UserStorage.
func (db *UserStorage) Delete(username string) error {
	q := "DELETE FROM users WHERE username = $1"
	_, err := db.Exec(q, username)
	return err
}

// List retrieves all Users from the in memory array in UserStorage.
func (db *UserStorage) List() ([]socialnet.User, error) {
	q := "SELECT * FROM users"
	var uu []socialnet.User

	err := db.Select(&uu, q)
	if err != nil {
		return []socialnet.User{}, err
	}

	return uu, nil
}

func getParamsValsArgsFromUser(usr socialnet.User) (params, vals string, args []interface{}) {
	if usr.Username != "" {
		params, vals, args = appendParamsAndArgs("username", usr.FirstName, params, vals, args)
	}

	if usr.FirstName != "" {
		params, vals, args = appendParamsAndArgs("first_name", usr.FirstName, params, vals, args)
	}

	if usr.LastName != "" {
		params, vals, args = appendParamsAndArgs("last_name", usr.LastName, params, vals, args)
	}

	if usr.Email != "" {
		params, vals, args = appendParamsAndArgs("email", usr.Email, params, vals, args)
	}

	if usr.ImageURL != "" {
		params, vals, args = appendParamsAndArgs("image_url", usr.ImageURL, params, vals, args)
	}
	if len(args) > 1 {
		params = "(" + params + ")"
		vals = "(" + vals + ")"
	}
	return params, vals, args
}

func appendParamsAndArgs(col, val, params, vals string, args []interface{}) (newParams, newVals string, newArgs []interface{}) {
	args = append(args, val)
	if len(args) > 1 {
		params += ", "
		vals += ", "
	}
	params += col
	vals += "$" + strconv.Itoa(len(args))
	return params, vals, args
}
