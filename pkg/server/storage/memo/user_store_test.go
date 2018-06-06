package memo_test

import (
	"testing"

	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/server/storage/memo"
)

func TestUserStore(t *testing.T) {
	userStore := memo.NewUserStorage()
	john := socialnet.User{
		Username: "jlennon", FirstName: "John", LastName: "Lennon", Email: "jlennon@beatles.com", Password: "Strawberryfields67!",
	}

	storedJohn, err := userStore.Create(john)
	if err != nil {
		t.Error(err)
	}

	if storedJohn.Username != john.Username {
		t.Errorf("username should be stored. have %s, want %s", storedJohn.Username, john.Username)
	}

	storedJohn, err = userStore.Read(john.Username)
	if err != nil {
		t.Error(err)
	}

	if storedJohn.Username != john.Username {
		t.Errorf("username should be stored. have %s, want %s", storedJohn.Username, john.Username)
	}

	john.LastName = "Lemon"
	newJohn, err := userStore.Update(john.Username, john)
	if err != nil {
		t.Error(err)
	}

	if newJohn.LastName != "Lemon" {
		t.Errorf("user not updated. have %#v, want %#v", newJohn, john)
	}

	err = userStore.Delete(john.Username)
	if err != nil {
		t.Error(err)
	}

	beatles, err := userStore.List()
	if err != nil {
		t.Error(err)
	}

	if len(beatles) > 0 {
		t.Errorf("I DON'T BELIEVE, IN BEATLES. have %#v, want none", beatles)
	}
}
