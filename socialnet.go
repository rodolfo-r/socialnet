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
	ImageURL  string `json:"imageURL" db:"image_url"`
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
	Store  UserStorage
	Auth   UserAuth
	Follow UserFollow
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
	CreateToken(username string) (token string, err error)
	ValidateToken(token string) (username string, err error)
	ValidateCreds(username, password string) error
}

// UserFollow updates follow relationships between users.
type UserFollow interface {
	Follow(follower, followee string) error
	Unfollow(follower, followee string) error
	Followers(username string) ([]UserItem, error)
	Following(username string) ([]UserItem, error)
}

// Profile is a public profile.
type Profile struct {
	Username   string     `json:"username"`
	ImageURL   string     `json:"imageURL" db:"image_url"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Posts      []Post     `json:"posts"`
	Followers  []UserItem `json:"followers" db:"-"`
	Following  []UserItem `json:"following" db:"-"`
	IsFollower bool       `json:"isFollower" db:"-"`
}

// Settings contain personal details.
type Settings struct {
	Username  string `json:"username"`
	ImageURL  string `json:"imageURL" db:"image_url"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// UserItem is used in user list. i.e. a user's followers
type UserItem struct {
	Username  string `json:"username"`
	ImageURL  string `json:"imageURL" db:"image_url"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
}

// Feed is a collection of a user's following
// list posts
type Feed []FeedItem

// FeedItem is a user post
type FeedItem struct {
	ProfileImageURL string `json:"imageURL"`
	Post            `json:"post"`
}
