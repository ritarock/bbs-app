package entity

import (
	"time"

	"github.com/ritarock/bbs-app/domain/valueobject"
)

type User struct {
	id           valueobject.UserID
	email        valueobject.Email
	passwordHash string
	createdAt    time.Time
}

func NewUser(email, passwordHash string) (*User, error) {
	emailVO, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{
		email:        emailVO,
		passwordHash: passwordHash,
		createdAt:    time.Now(),
	}, nil
}

func (u *User) ID() valueobject.UserID {
	return u.id
}

func (u *User) Email() valueobject.Email {
	return u.email
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func ReconstructUser(id valueobject.UserID, email, passwordHash string, createdAt time.Time) *User {
	return &User{
		id:           id,
		email:        valueobject.ReconstructEmail(email),
		passwordHash: passwordHash,
		createdAt:    createdAt,
	}
}
