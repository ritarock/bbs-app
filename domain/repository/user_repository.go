package repository

import (
	"context"

	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/valueobject"
)

type UserRepository interface {
	Save(ctx context.Context, user *entity.User) (valueobject.UserID, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByID(ctx context.Context, id valueobject.UserID) (*entity.User, error)
}
