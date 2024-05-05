package port

import (
	"context"

	"github.com/ritarock/bbs-app/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByNameAndPassword(ctx context.Context,
		name, password string) (*domain.User, error)
	UpdateToken(ctx context.Context, userID int, token string) error
	ExistToken(ctx context.Context, token string) (bool, error)
}

type UserUsecase interface {
	Signup(ctx context.Context, user *domain.User) error
	Login(ctx context.Context, name, password string) (*domain.User, error)
	UpdateToken(ctx context.Context, userID int, token string) error
	IsTokenAvailable(ctx context.Context, token string) bool
}
