package handler_test

import (
	"net/http/httptest"
	"testing"

	"github.com/rodolfo-r/socialnet"
	"github.com/rodolfo-r/socialnet/pkg/server/auth"
	"github.com/rodolfo-r/socialnet/pkg/server/handler"
	"github.com/rodolfo-r/socialnet/pkg/server/storage/memo"
)

func TestLike(t *testing.T) {
	t.Parallel()

	usrStore := memo.NewUserStorage()
	usrSvc := socialnet.UserService{Store: usrStore, Auth: *auth.New(usrStore, "localhost", "test jwt string")}
	postSvc := socialnet.PostService{Store: memo.NewPostStorage(), Like: memo.NewLikeStorage()}

	router := handler.New(usrSvc, postSvc, handler.Options{
		Signature: "jwt test signature",
	})

	user := socialnet.User{
		Username: "jlennon", FirstName: "john", LastName: "lennon", Password: "strawberryfields",
	}
	_, err := usrSvc.Store.Create(user)
	if err != nil {
		t.Fatal(err)
	}

	post := socialnet.Post{
		Author: "gharrison", Title: "Something", Body: "Something in the way she moves",
	}
	_, err = postSvc.Store.Create(post)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRequest("POST", "/like", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, r)
}

func TestLikeLike(t *testing.T) {

}

func TestLikeUnlike(t *testing.T) {

}
