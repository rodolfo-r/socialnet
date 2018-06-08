package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"github.com/techmexdev/handlertest"
	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/server/auth"
	"github.com/techmexdev/socialnet/pkg/server/handler"
	"github.com/techmexdev/socialnet/pkg/server/storage/memo"
)

func TestUsersEmpty(t *testing.T) {
	t.Parallel()

	usrStore := memo.NewUserStorage()
	usrSvc := socialnet.UserService{Store: usrStore, Auth: *auth.New(usrStore, "localhost", "test jwt string")}
	postSvc := socialnet.PostService{Store: memo.NewPostStorage()}

	router := handler.New(usrSvc, postSvc, handler.Options{
		Signature: "jwt test signature",
	})

	tc := handlertest.TestCase{
		Name:       "No users",
		StatusCode: http.StatusOK,
		Request:    httptest.NewRequest("GET", "/users", nil),
	}

	handlertest.Test(t, tc, router)
}

func TestUsersNotEmpty(t *testing.T) {
	t.Parallel()

	usrStore := memo.NewUserStorage()
	usrSvc := socialnet.UserService{Store: usrStore, Auth: *auth.New(usrStore, "localhost", "test jwt string")}
	postSvc := socialnet.PostService{Store: memo.NewPostStorage()}

	uu := []socialnet.User{
		{Username: "sgt-pepper", FirstName: "Sargent", LastName: "Pepper", Email: "sgt@pepper.com", Password: "lonely-hearts"},
		{Username: "bshears", FirstName: "Billy", LastName: "Shears", Email: "b@shears.com", Password: "shears-billy"},
	}

	for _, u := range uu {
		_, err := usrSvc.Store.Create(u)
		if err != nil {
			t.Fatal(err)
		}
	}

	router := handler.New(usrSvc, postSvc, handler.Options{
		Signature: "jwt test signature",
	})

	tc := handlertest.TestCase{
		Name:       "No users",
		StatusCode: http.StatusOK,
		Request:    httptest.NewRequest("GET", "/users", nil),
		BodyAssert: func(b []byte) error {
			var pp []socialnet.Profile
			err := json.Unmarshal(b, &pp)
			if err != nil {
				return err
			}
			return handlertest.Assert(len(pp) == len(uu), "should respond with 2 profiles")
		},
	}

	handlertest.Test(t, tc, router)
}
