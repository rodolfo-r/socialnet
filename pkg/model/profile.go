package model

// Profile is a user's public information
type Profile struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Posts     []Post `json:"posts"`
}
