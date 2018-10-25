package memo

import (
	"github.com/rodolfo-r/socialnet"
)

type follow struct {
	follower string
	followee string
}

// UserFollow is an in-memory user follow storage.
type UserFollow struct {
	follows []follow
}

// NewUserFollow returns an in memory socialnet.UserFollow.
func NewUserFollow() *UserFollow {
	return &UserFollow{}
}

// Follow adds a follow relationship to the UserFollow.
func (uf *UserFollow) Follow(follower, followee string) error {
	for _, f := range uf.follows {
		if f.follower == follower && f.followee == followee {
			return nil
		}
	}

	uf.follows = append(uf.follows, follow{follower, followee})
	return nil
}

// Followers returns a user's followers
func (uf *UserFollow) Followers(username string) ([]socialnet.UserItem, error) {
	var uu []socialnet.UserItem
	for _, f := range uf.follows {
		if f.followee == username {
			u := socialnet.UserItem{Username: f.follower}
			uu = append(uu, u)
		}
	}
	return uu, nil
}

// Following returns a the user's a user is following
func (uf *UserFollow) Following(username string) ([]socialnet.UserItem, error) {
	var uu []socialnet.UserItem
	for _, f := range uf.follows {
		if f.follower == username {
			u := socialnet.UserItem{Username: f.followee}
			uu = append(uu, u)
		}
	}
	return uu, nil
}

// Unfollow removes a relationship from the UserFollow.
func (uf *UserFollow) Unfollow(follower, followee string) error {
	for i := range uf.follows {
		if uf.follows[i].follower == follower && uf.follows[i].followee == followee {
			uf.follows = append(uf.follows[:i], uf.follows[i+1:]...)
			return nil
		}
	}

	return nil
}
