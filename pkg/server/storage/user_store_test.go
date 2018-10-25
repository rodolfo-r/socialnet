package storage_test

import (
	"testing"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/rodolfo-r/socialnet"
	"github.com/rodolfo-r/socialnet/pkg/server/storage/memo"
	"github.com/rodolfo-r/socialnet/pkg/server/storage/postgres"
)

var userStore socialnet.UserStorage
var user socialnet.User

func TestUserStore(t *testing.T) {
	if testMemo {
		userStore = memo.NewUserStorage()
	} else {
		postgres.MigrateDown("file://postgres/migrations", dsn)
		postgres.MigrateUp("file://postgres/migrations", dsn)

		userStore = postgres.NewUserStorage(dsn)
	}

	user = socialnet.User{
		Username: "jlennon", FirstName: "John", LastName: "Lennon", Email: "jlennon@beatles.com", Password: "Strawberryfields67!",
	}
}

func TestUserStoreCreate(t *testing.T) {
	u, err := userStore.Create(user)
	if err != nil {
		t.Error("error creating user: ", err)
	}

	if u.Username != user.Username {
		t.Errorf("username should be stored. have %s, want %s", u.Username, user.Username)
	}
}

func TestUserStoreRead(t *testing.T) {
	u, err := userStore.Read(user.Username)
	if err != nil {
		t.Error("error reading user: ", err)
	}

	if u.Username != user.Username {
		t.Errorf("username should be stored. have %s, want %s", u.Username, user.Username)
	}

	if u.Password == user.Password {
		t.Errorf("user storage should hash passwords")
	}
}

func TestUserStoreUpdate(t *testing.T) {
	u, err := userStore.Update(user.Username, socialnet.User{LastName: "Lemon"})
	if err != nil {
		t.Error("error updating user: ", err)
	}

	if u.LastName != "Lemon" {
		t.Errorf("user not updated. have %#v, want %#v", u, user)
	}
}

func TestUserStoreDelete(t *testing.T) {
	err := userStore.Delete(user.Username)
	if err != nil {
		t.Error("error deleting user: ", err)
	}
}

func TestUserList(t *testing.T) {
	uu, err := userStore.List()
	if err != nil {
		t.Error("error listing users: ", err)
	}

	if len(uu) > 0 {
		t.Errorf("error listing users: have %#v, want none", uu)
	}
}
