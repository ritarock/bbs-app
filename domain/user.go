package domain

import (
	"context"
	"errors"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserUsecase interface {
	SignUp(ctx context.Context, user *User) error
	Login(ctx context.Context, user *User) (bool, *User)
	SetToken(ctx context.Context, userId int, token string) error
	ValidateToken(ctx context.Context, token string) bool
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindUser(ctx context.Context, name string, password string) (*User, error)
	SetToken(ctx context.Context, userId int, token string) error
	ExistToken(ctx context.Context, token string) bool
}

func (u *User) Validate() error {
	if len(u.Name) > 30 {
		return errors.New("name length must be less than equal to 30")
	}
	if len(u.Password) < 8 {
		return errors.New("password length must be more than equal to 8")
	}
	return nil
}
