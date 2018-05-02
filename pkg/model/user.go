package model

import "errors"

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (u *User) Validate() error {
	fields := []string{u.FirstName, u.LastName, u.Email, u.Username,}
	emptyField := false
	emptyMsg := "user contains empty fields: "

	for _, f := range fields {
		if len(f) == 0 {
			emptyField = true
			emptyMsg += ", " + f
		}
	}

	if emptyField {
		return errors.New(emptyMsg)
	}

	return nil
}
