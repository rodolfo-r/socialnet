package postgres

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/techmexdev/reddit-api/pkg/model"
)

// Postgres is an implementation of Storage
type Postgres struct {
	*sqlx.DB
}

// New returns a pointer to a pg connection
func New(dsn string) *Postgres {
	p := &Postgres{sqlx.MustConnect("postgres", dsn)}
	err := p.init()
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func (p *Postgres) init() error {

}

// InsertUser Inserts new user into users table.
func (p *Postgres) CreateUser(model.User) (model.User, error) {
	q := "INSERT INTO users (id, username, email, password, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6)"

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	createdAt := time.Now().Format(time.RFC3339)

	_, err = db.Exec(q, id, usr.Username, usr.Email, pwd, createdAt, createdAt)
	if err != nil {
		return err
	}

	return nil
}

// Get User retrieves a username
func (db *Postgres) GetUser(username string) (model.User, error) {
	q := "SELECT username, email FROM users WHERE username = $0"
	row := db.QueryRow(q, username)
	var u model.User
	err := row.Scan(&u)
	if err != nil {
		return model.User{}, err
	}

	return u, nil
}

// GetUsers Retrieves all users' username & email from users table.
func (db *Postgres) GetUsers() ([]model.User, error) {
	q := "SELECT username, email FROM users;"
	rows, err := db.Queryx(q)
	if err != nil {
		return nil, err
	}

	users := []model.User{}
	for rows.Next() {
		var u model.User
		err := rows.StructScan(&u)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
