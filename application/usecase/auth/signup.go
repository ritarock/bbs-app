package auth

import (
	"context"
	"errors"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/application/service"
	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/repository"
	"github.com/ritarock/bbs-app/domain/valueobject"
)

type SignUpUsecase struct {
	userRepo        repository.UserRepository
	passwordService service.PasswordService
	tokenService    service.TokenService
}

func NewSignUpUsecase(
	userRepo repository.UserRepository,
	passwordService service.PasswordService,
	tokenService service.TokenService,
) *SignUpUsecase {
	return &SignUpUsecase{
		userRepo:        userRepo,
		passwordService: passwordService,
		tokenService:    tokenService,
	}
}

func (u *SignUpUsecase) Execute(ctx context.Context, input dto.SignUpInput) (*dto.SignUpOutput, error) {
	_, err := valueobject.NewPassword(input.Password)
	if err != nil {
		return nil, err
	}

	existingUser, _ := u.userRepo.FindByEmail(ctx, input.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	passwordHash, err := u.passwordService.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	user, err := entity.NewUser(input.Email, passwordHash)
	if err != nil {
		return nil, err
	}

	userID, err := u.userRepo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := u.tokenService.Generate(userID.Int())
	if err != nil {
		return nil, err
	}

	return &dto.SignUpOutput{
		ID:        userID.Int(),
		Email:     user.Email().String(),
		Token:     token,
		CreatedAt: user.CreatedAt(),
	}, nil
}
