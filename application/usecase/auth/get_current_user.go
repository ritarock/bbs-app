package auth

import (
	"context"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/domain/repository"
	"github.com/ritarock/bbs-app/domain/valueobject"
)

type GetCurrentUserUsecase struct {
	userRepo repository.UserRepository
}

func NewGetCurrentUserUsecase(userRepo repository.UserRepository) *GetCurrentUserUsecase {
	return &GetCurrentUserUsecase{
		userRepo: userRepo,
	}
}

func (u *GetCurrentUserUsecase) Execute(ctx context.Context, input dto.GetCurrentUserInput) (*dto.GetCurrentUserOutput, error) {
	user, err := u.userRepo.FindByID(ctx, valueobject.NewUserID(input.UserID))
	if err != nil {
		return nil, err
	}

	return &dto.GetCurrentUserOutput{
		ID:        user.ID().Int(),
		Email:     user.Email().String(),
		CreatedAt: user.CreatedAt(),
	}, nil
}
