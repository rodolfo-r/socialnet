package memo

import (
	"fmt"
	"strings"

	"github.com/rodolfo-r/socialnet"
	"golang.org/x/crypto/bcrypt"
)

// UserStorage is an in-memory socialnet.UserStorage.
type UserStorage struct {
	users []socialnet.User
}

// NewUserStorage returns an in-memory socialnet.UserStorage.
func NewUserStorage() *UserStorage {
	return &UserStorage{}
}

// Create adds a User to the in memory array in UserStorage.
func (db *UserStorage) Create(usr socialnet.User) (socialnet.User, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 12)
	if err != nil {
		return socialnet.User{}, err
	}

	usr.Password = string(b)

	db.users = append(db.users, usr)
	return db.users[len(db.users)-1], nil
}

// Read retrieves a User from the in memory array in UserStorage.
func (db *UserStorage) Read(username string) (socialnet.User, error) {
	for _, u := range db.users {
		if strings.ToLower(username) == strings.ToLower(u.Username) {
			return u, nil
		}
	}
	return socialnet.User{}, fmt.Errorf("Requested User %s not found", username)
}

// Update replaces a User from the in memory array in UserStorage.
func (db *UserStorage) Update(username string, usr socialnet.User) (socialnet.User, error) {
	for i := range db.users {
		if db.users[i].Username == username {
			db.users[i] = usr
			return db.users[i], nil
		}
	}
	return socialnet.User{}, fmt.Errorf(
		"could not find user with username: %s",
		username,
	)
}

// Delete removes a User from the in memory array in UserStorage.
func (db *UserStorage) Delete(username string) error {
	for i := range db.users {
		if db.users[i].Username == username {
			db.users = append(db.users[:i], db.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf(
		"could not find user with username: %s",
		username,
	)
}

// List retrieves all Users from the in memory array in UserStorage.
func (db *UserStorage) List() ([]socialnet.User, error) {
	return db.users, nil
}
