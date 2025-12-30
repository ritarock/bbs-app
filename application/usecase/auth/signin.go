package auth

import (
	"context"
	"errors"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/application/service"
	"github.com/ritarock/bbs-app/domain/repository"
)

type SignInUsecase struct {
	userRepo        repository.UserRepository
	passwordService service.PasswordService
	tokenService    service.TokenService
}

func NewSignInUsecase(
	userRepo repository.UserRepository,
	passwordService service.PasswordService,
	tokenService service.TokenService,
) *SignInUsecase {
	return &SignInUsecase{
		userRepo:        userRepo,
		passwordService: passwordService,
		tokenService:    tokenService,
	}
}

func (u *SignInUsecase) Execute(ctx context.Context, input dto.SignInInput) (*dto.SignInOutput, error) {
	user, err := u.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = u.passwordService.Verify(input.Password, user.PasswordHash())
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := u.tokenService.Generate(user.ID().Int())
	if err != nil {
		return nil, err
	}

	return &dto.SignInOutput{
		ID:        user.ID().Int(),
		Email:     user.Email().String(),
		Token:     token,
		CreatedAt: user.CreatedAt(),
	}, nil
}
