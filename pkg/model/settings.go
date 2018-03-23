package model

// Settings contain personal details
type Settings struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email     string `json:"email"`
}
