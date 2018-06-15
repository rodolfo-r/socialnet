package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/techmexdev/handlertest"
	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/server/auth"
	"github.com/techmexdev/socialnet/pkg/server/handler"
	"github.com/techmexdev/socialnet/pkg/server/storage/memo"
)

func TestFeed(t *testing.T) {
	t.Parallel()

	usrStore := memo.NewUserStorage()
	usrSvc := socialnet.UserService{
		Store: usrStore, Auth: *auth.New(usrStore, "localhost", "test jwt signature"), Follow: memo.NewUserFollow(),
	}
	postSvc := socialnet.PostService{Store: memo.NewPostStorage()}

	router := handler.New(usrSvc, postSvc, handler.Options{
		Signature: "jwt test signature",
	})

	token, err := usrSvc.Auth.CreateToken("jlennon")
	if err != nil {
		t.Fatal(err)
	}

	tc := handlertest.TestCase{
		Name:       "No feed",
		StatusCode: http.StatusOK,
		Request: &http.Request{
			Method: "GET",
			URL:    &url.URL{Path: "/feed"},
			Header: map[string][]string{"Authorization": {"Bearer " + token}},
			Body:   ioutil.NopCloser(strings.NewReader("")),
		},
	}

	handlertest.Test(t, tc, router)
}

func TestFeedOnePost(t *testing.T) {
	t.Parallel()

	usrStore := memo.NewUserStorage()
	usrSvc := socialnet.UserService{
		Store: usrStore, Auth: *auth.New(usrStore, "localhost", "test jwt string"), Follow: memo.NewUserFollow(),
	}

	_, err := usrSvc.Store.Create(socialnet.User{Username: "jlennon", Password: "strawberryfields"})
	if err != nil {
		t.Fatal(err)
	}

	_, err = usrSvc.Store.Create(socialnet.User{Username: "gharrison"})
	if err != nil {
		t.Fatal(err)
	}

	err = usrSvc.Follow.Follow("jlennon", "gharrison")
	if err != nil {
		t.Fatal(err)
	}

	postSvc := socialnet.PostService{Store: memo.NewPostStorage()}
	_, err = postSvc.Store.Create(socialnet.Post{Author: "gharrison", Title: "Hello"})
	if err != nil {
		t.Fatal(err)
	}

	token, err := usrSvc.Auth.CreateToken("jlennon")
	if err != nil {
		t.Fatal(err)
	}

	router := handler.New(usrSvc, postSvc, handler.Options{
		Signature: "jwt test signature",
	})

	tc := handlertest.TestCase{
		Name:       "One post",
		StatusCode: http.StatusOK,
		Request: &http.Request{
			Method: "GET",
			URL:    &url.URL{Path: "/feed"},
			Header: map[string][]string{"Authorization": {"Bearer " + token}},
			Body:   ioutil.NopCloser(strings.NewReader("")),
		},
		BodyAssert: func(b []byte) error {
			return handlertest.Assert(strings.Contains(string(b), "gharrison"), "feed should contain a gharrison post")
		},
	}
	handlertest.Test(t, tc, router)
}
