package auth_test

import (
	"testing"

	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/auth"
	"github.com/techmexdev/socialnet/pkg/storage/memo"
)

func TestValidate(t *testing.T) {
	usrStore := memo.NewUserStorage()
	_, err := usrStore.Create(socialnet.User{Username: "jlennon", Password: "Strawberryfields67!"})
	if err != nil {
		t.Fatal(err)
	}

	usrAuth := auth.New(usrStore)

	err = usrAuth.Validate("jlennon", "password123")
	if err == nil {
		t.Fatal("should error on incorrect password")
	}

	err = usrAuth.Validate("jlennon", "Strawberryfields67!")
	if err != nil {
		t.Fatal("failed validation on correct password", err)
	}

	err = usrAuth.Validate("jlennon", "")
	if err == nil {
		t.Fatal("should fail for blank passwords")
	}

	err = usrAuth.Validate("gharrison", "password123")
	if err == nil {
		t.Fatal("should fail for non-existent users")
	}
}
