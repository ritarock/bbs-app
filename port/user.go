package port

import (
	"context"

	"github.com/ritarock/bbs-app/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByNameAndPasswd(ctx context.Context, name, password string) (*domain.User, error)
}

type UserUsecase interface {
	Create(ctx context.Context, user *domain.User) error
	Find(ctx context.Context, name, pasword string) (*domain.User, error)
}
