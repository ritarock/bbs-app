package valueobject

import "errors"

type Password struct {
	value string
}

func NewPassword(password string) (Password, error) {
	if len(password) < 8 {
		return Password{}, errors.New("password must be at least 8 characters")
	}
	if len(password) > 128 {
		return Password{}, errors.New("password must be 128 characters or less")
	}
	return Password{value: password}, nil
}

func (p Password) String() string {
	return p.value
}
