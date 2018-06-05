package socialnet

import "time"

// Post is a post submission.
type Post struct {
	CreatedAt time.Time `json:"createdAt"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
}

// PostService stores and retrieves posts.
type PostService struct {
	Store PostStorage
}

// PostStorage is an interface for
// storing and retrieving posts.
type PostStorage interface {
	Create(Post) (Post, error)
	Read(author, title string) (Post, error)
	Update(author, title string, newPost Post) (Post, error)
	Delete(author, title string) error
	List() ([]Post, error)
}

// User is a social network user.
type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Posts     []Post `json:"posts"`
	Password  string `json:"password"`
}

// UserService stores, and authenticates
// a user.
type UserService struct {
	Store UserStorage
	Auth  UserAuth
}

// UserStorage is an interface for
// storing and retrieving posts.
type UserStorage interface {
	Create(User) (User, error)
	Read(username string) (User, error)
	Update(username string, newUsr User) (User, error)
	Delete(username string) error
	List() ([]User, error)
}

// UserAuth validates a user's credentials
type UserAuth interface {
	Validate(username, password string) error
}

// Profile is a public profile.
type Profile struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Posts     []Post `json:"posts"`
}

// Settings contain personal details.
type Settings struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
