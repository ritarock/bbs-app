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

func (u *userUsecase) Signup(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	_, err := u.userRepo.GetByNameAndPassword(ctx, user.Name, user.Password)
	if err == nil {
		return domain.AlreadyUserExists
	}
	if err != domain.ErrNotFound {
		return err
	}

	return u.userRepo.Create(ctx, user)
}

func (u *userUsecase) Login(ctx context.Context, name string, password string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.GetByNameAndPassword(ctx, name, password)
}

func (u *userUsecase) UpdateToken(ctx context.Context, userID int, token string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.UpdateToken(ctx, userID, token)
}

func (u *userUsecase) IsTokenAvailable(ctx context.Context, token string) bool {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	existToken, err := u.userRepo.ExistToken(ctx, token)
	if err != nil {
		return false
	}
	return existToken
}
