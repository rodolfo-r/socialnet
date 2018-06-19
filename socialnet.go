package socialnet

import "time"

// PostStorage is an interface for
// storing and retrieving posts.
type PostStorage interface {
	Create(Post) (Post, error)
	Read(id string) (Post, error)
	Update(id string, newPost Post) (Post, error)
	Delete(id string) error
	List(username string) ([]Post, error)
}

// Post is a post submission.
type Post struct {
	ID        string    `json:"id" db:"id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	Author    string    `json:"author" db:"username"`
	UserID    string    `json:"-" db:"users_id"`
	ImageURL  string    `json:"imageURL" db:"image_url"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Likes     []Like    `json:"likes" db:"-"`
	Comments  []Comment `json:"comments" db:"-"`
}

// PostService stores and retrieves posts.
type PostService struct {
	Store   PostStorage
	Like    LikeStorage
	Comment CommentStorage
}

// LikeStorage stores and retrieves post likes.
type LikeStorage interface {
	Create(username, postID string) error
	Delete(username, postID string) error
	List(postID string) ([]Like, error)
}

// Like is a post's like.
type Like struct {
	ID       string `json:"-" db:"id"`
	PostID   string `json:"postID" db:"post_id"`
	UserItem `json:"user"`
}

// CommentStorage stores and retrieves post comments.
type CommentStorage interface {
	Create(username, postID, text string) error
	Delete(username, postID string) error
	List(postID string) ([]Comment, error)
}

// Comment is a post comment made by a user.
type Comment struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt" db:"created_at"`
	PostID    string `json:"postID" db:"post_id"`
	UserItem  `json:"user"`
	Text      string `json:"text"`
}

// User is a social network user.
type User struct {
	ID        string    `json:"-"`
	Username  string    `json:"username"`
	ImageURL  string    `json:"imageURL" db:"image_url"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	Email     string    `json:"email"`
	Posts     []Post    `json:"posts"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
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
	Liked           bool `json:"liked" db:"-"`
}
