package auth_test

import (
	"testing"

	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/auth"
	"github.com/techmexdev/socialnet/pkg/storage/memo"
)

func TestValidateCreds(t *testing.T) {
	userStore := memo.NewUserStorage()
	_, err := userStore.Create(socialnet.User{Username: "jlennon", Password: "Strawberryfields67!"})
	if err != nil {
		t.Error(err)
	}

	userAuth := auth.New(userStore, "localhost", "test jwt signature")

	err = userAuth.ValidateCreds("jlennon", "password123")
	if err == nil {
		t.Error("should error on incorrect password")
	}

	err = userAuth.ValidateCreds("jlennon", "Strawberryfields67!")
	if err != nil {
		t.Error("failed validation on correct password", err)
	}

	err = userAuth.ValidateCreds("jlennon", "")
	if err == nil {
		t.Error("should fail for blank passwords")
	}

	err = userAuth.ValidateCreds("gharrison", "password123")
	if err == nil {
		t.Error("should fail for non-existent users")
	}
}

func TestCreateAndValidateToken(t *testing.T) {
	userStore := memo.NewUserStorage()

	userAuth := auth.New(userStore, "localhost", "test jwt signature")

	token, err := userAuth.CreateToken("gharrison")
	if err != nil {
		t.Errorf("could not create token for gharrison: %s", err)
	}

	username, err := userAuth.ValidateToken(token)
	if err != nil {
		t.Errorf("could not validate token: %s", err)
	}

	fakeUserAuth := auth.New(userStore, "localhost", "incorrect jwt signature")

	fakeToken, err := fakeUserAuth.CreateToken("gharrison")
	if err != nil {
		t.Errorf("could not create fake token for %s: %s", username, err)
	}

	_, err = userAuth.ValidateToken(fakeToken)
	if err == nil {
		t.Errorf("should error on fake tokens")
	}
}
