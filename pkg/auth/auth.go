package auth

import (
	"github.com/techmexdev/socialnet"
	"golang.org/x/crypto/bcrypt"
)

// UserAuth requires a UserStorage to retrieve stored passwords.
type UserAuth struct {
	store socialnet.UserStorage
}

// New returns a UserAuth.
func New(store socialnet.UserStorage) *UserAuth {
	return &UserAuth{store}
}

// Validate compares the given and stored passwords using bcrypt.
func (userAuth UserAuth) Validate(username, password string) error {
	usr, err := userAuth.store.Read(username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	return err
}
