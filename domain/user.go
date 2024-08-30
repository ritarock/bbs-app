package domain

import "errors"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *User) Validate() error {
	if u.Name == "" || len(u.Name) > 30 {
		return errors.New("length of name must range from 1 to 30 characters")
	}

	if len(u.Password) < 8 || len(u.Password) > 30 {
		return errors.New("length of password must range from 8 to 30 characters")
	}

	return nil
}
