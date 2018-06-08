package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/techmexdev/socialnet"
)

// ProfilePicture saves jpg files to the file system.
func (h *handler) ProfilePicture(w http.ResponseWriter, r *http.Request) {
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	username, err := h.userSvc.Auth.ValidateToken(token)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	r.ParseMultipartForm(32 << 20)

	err = os.MkdirAll("./files/user/"+username+"/", os.ModePerm)
	if err != nil && err != os.ErrExist {
		serverError(w, fmt.Errorf("error creating static user directory: %s", err))
		return
	}

	newFile, err := os.OpenFile("./files/user/"+username+"/profile.jpg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newFile.Close()

	file, _, err := r.FormFile("image")
	if err != nil {
		log.Println("error reading file from form: ", err)
		http.Error(w, "expecting a multipart/form-data image with name 'image' in request.", http.StatusUnprocessableEntity)
		return
	}
	defer file.Close()

	io.Copy(newFile, file)

	// images will be accessed from the client fileserver.
	imgURL := "http://localhost:3001/files/user/" + username + "/profile.jpg"
	_, err = h.userSvc.Store.Update(username, socialnet.User{ImageURL: imgURL})
	if err != nil {
		serverError(w, err)
	}
}
