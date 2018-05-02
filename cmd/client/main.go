package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	addr := os.Getenv("ADDRESS")
	res, err := http.Post("http://"+addr+"/signup", "application/json", strings.NewReader(`
		{"username": "georgeh", "email": "george@beatles.com",
		"firstName": "George", "lastName": "Harrison", "password": "h3r3c0m3sth35un"}
	`))
	if err != nil {
		fmt.Println("POST /login request failed: " + err.Error())
		return
	}

	resBody := struct{ token string }{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		log.Println("could not decode response in json: " + err.Error())
		return
	}

	fmt.Printf("response body: %v", resBody)

	res, err = http.Post("http://"+addr+"/login", "application/json", strings.NewReader(`
	{"username": "georgeh", "password": "h3r3c0m3sth35un"}
	`))

	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("response body: %s", b)
}
