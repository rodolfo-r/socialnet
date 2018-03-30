package mock

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/techmexdev/the_social_network/pkg/model"
	"github.com/techmexdev/the_social_network/pkg/storage"
	"golang.org/x/crypto/bcrypt"
)

// Mock is a mock storage
type Mock struct {
	users []model.User
	posts []model.Post
}

// New returns a mock storage
func New() storage.Storage {
	return &Mock{}
}

// GetUser uses email and username to find a user
func (m *Mock) GetUser(usr model.User) (model.User, error) {
	for _, u := range m.users {
		if strings.ToLower(u.Email) == strings.ToLower(usr.Email) ||
			strings.ToLower(u.Username) == strings.ToLower(usr.Username) {
			return u, nil
		}
	}
	return model.User{}, fmt.Errorf("Requested User %#v not found", usr)
}

// CreateUser adds a user to the users array without erroring
func (m *Mock) CreateUser(usr model.User) (model.User, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 12)
	if err != nil {
		return model.User{}, err
	}
	usr.Password = string(b)
	m.users = append(m.users, usr)
	return m.users[len(m.users)-1], nil
}

// ValidateUserCreds compares the given passwod with the one
// stored in m.users for a given user
func (m *Mock) ValidateUserCreds(username, password string) error {
	usr, err := m.GetUser(model.User{Username: username})
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// GetProfile returns public data
func (m *Mock) GetProfile(username string) (model.Profile, error) {
	for _, u := range m.users {
		if strings.ToLower(u.Username) == strings.ToLower(username) {
			ps, err := m.getUserPosts(u.Username)
			if err != nil {
				return model.Profile{}, err
			}

			return model.Profile{
				Username: u.Username, FirstName: u.FirstName, LastName: u.LastName, Posts: ps,
			}, nil
		}
	}
	return model.Profile{}, errors.New("Requested Profile not found")
}

// CreatePost adds a post to m.users
func (m *Mock) CreatePost(p model.Post) (model.Post, error) {
	p.CreatedAt = time.Now()
	m.posts = append(m.posts, p)
	return m.posts[len(m.posts)-1], nil
}

// GetUserSettings returns settings for a user
func (m *Mock) GetUserSettings(username string) (model.Settings, error) {
	usr, err := m.GetUser(model.User{Username: username})
	if err != nil {
		return model.Settings{}, err
	}

	return model.Settings{
		Username: usr.Username, FirstName: usr.FirstName, LastName: usr.LastName, Email: usr.Email,
	}, nil
}

func (m *Mock) getUserPosts(username string) ([]model.Post, error) {
	foundPosts := []model.Post{}
	for _, p := range m.posts {
		if strings.ToLower(username) == strings.ToLower(p.Username) {
			foundPosts = append(foundPosts, p)
		}
	}
	return foundPosts, nil
}
