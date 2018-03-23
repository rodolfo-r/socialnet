package storage

import "github.com/techmexdev/the_social_network/pkg/model"

// Storage in an interface to db or mock db
type Storage interface {
	GetUser(model.User) (model.User, error)
	CreateUser(model.User) (model.User, error)
	ValidateUserCreds(username string, password string) error
	GetProfile(username string) (model.Profile, error)
	CreatePost(model.Post) (model.Post, error)
	GetUserSettings(username string) (model.Settings, error)
}
