package usecase

import (
	"context"
	"time"

	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/port"
)

type userUsecase struct {
	userRepo       port.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(repo port.UserRepository, timeout time.Duration) port.UserUsecase {
	return &userUsecase{
		userRepo:       repo,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) Create(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.Create(ctx, user)
}

func (u *userUsecase) Find(ctx context.Context, name string, pasword string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.GetByNameAndPasswd(ctx, name, pasword)
}
