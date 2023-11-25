package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/ritarock/bbs-app/domain"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(repo domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       repo,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) SignUp(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	_, err := u.userRepo.FindUser(ctx, user.Name, user.Password)
	switch err {
	case domain.ErrNotFound:
		return u.userRepo.Create(ctx, user)
	case errors.New("system error"), nil:
		return errors.New("already exists")
	default:
		return err
	}
}

func (u *userUsecase) Login(ctx context.Context, user *domain.User) (bool, *domain.User) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	findUser, err := u.userRepo.FindUser(ctx, user.Name, user.Password)
	return err == nil, findUser
}

func (u *userUsecase) SetToken(ctx context.Context, userId int, token string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.SetToken(ctx, userId, token)
}

func (u *userUsecase) ValidateToken(ctx context.Context, token string) bool {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.ExistToken(ctx, token)
}
