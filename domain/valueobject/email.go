package valueobject

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewEmail(email string) (Email, error) {
	if len(email) == 0 {
		return Email{}, errors.New("email is required")
	}
	if len(email) > 255 {
		return Email{}, errors.New("email must be 255 characters or less")
	}
	if !emailRegex.MatchString(email) {
		return Email{}, errors.New("invalid email format")
	}
	return Email{value: email}, nil
}

func ReconstructEmail(email string) Email {
	return Email{value: email}
}

func (e Email) String() string {
	return e.value
}
