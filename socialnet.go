package socialnet

import "time"

// Post is a post submission.
type Post struct {
	ID        string    `json:"-" db:"id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	Author    string    `json:"author"`
	UsersID   string    `json:"-" db:"users_id"`
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
	ID        string `json:"-"`
	Username  string `json:"username"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email"`
	Posts     []Post `json:"posts"`
	Password  string `json:"password"`
	CreatedAt string `json:"-" db:"created_at"`
	UpdatedAt string `json:"-" db:"updated_at"`
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
