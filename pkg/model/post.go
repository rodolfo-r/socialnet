package model

import "time"

// Post is user subitted
type Post struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	Caption   string    `json:"caption"`
}
