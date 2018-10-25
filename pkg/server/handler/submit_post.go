package handler

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/rodolfo-r/socialnet"
)

func (h *handler) SubmitPost(w http.ResponseWriter, r *http.Request) {
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	username, err := h.userSvc.Auth.ValidateToken(token)
	if err != nil {
		http.Error(w, "Invalid token. Header 'Authorization' must have value: 'Bearer <token>'", http.StatusUnauthorized)
		return
	}

	r.ParseMultipartForm(32 << 20)

	title := r.FormValue("title")
	body := r.FormValue("body")

	post := socialnet.Post{Author: username, Body: body, Title: title}
	post, err = h.postSvc.Store.Create(post)
	if err != nil {
		serverError(w, err)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		h.r.JSON(w, http.StatusCreated, post)
		return
	}
	defer file.Close()

	h.createPostImage(w, username, post.ID, file)
}

func (h *handler) createPostImage(w http.ResponseWriter, username, postID string, file multipart.File) {
	err := os.MkdirAll("./files/user/"+username+"/posts/", os.ModePerm)
	if err != nil && err != os.ErrExist {
		err = fmt.Errorf("error creating static user directory: %s", err)
		serverError(w, err)
		return
	}

	newFile, err := os.OpenFile("./files/user/"+username+"/posts/"+postID+".jpg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		serverError(w, err)
		return
	}
	defer newFile.Close()

	io.Copy(newFile, file)

	// images will be accessed from the client fileserver.
	imgURL := "http://localhost:3001/files/user/" + username + "/posts/" + postID + ".jpg"
	newPost := socialnet.Post{ImageURL: imgURL}
	_, err = h.postSvc.Store.Update(postID, newPost)
	if err != nil {
		serverError(w, err)
	}
	h.r.JSON(w, http.StatusCreated, http.StatusText(http.StatusCreated))

}
