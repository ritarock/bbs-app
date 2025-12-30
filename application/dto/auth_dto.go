package dto

import "time"

type SignUpInput struct {
	Email    string
	Password string
}

type SignUpOutput struct {
	ID        int
	Email     string
	Token     string
	CreatedAt time.Time
}

type SignInInput struct {
	Email    string
	Password string
}

type SignInOutput struct {
	ID        int
	Email     string
	Token     string
	CreatedAt time.Time
}

type GetCurrentUserInput struct {
	UserID int
}

type GetCurrentUserOutput struct {
	ID        int
	Email     string
	CreatedAt time.Time
}
