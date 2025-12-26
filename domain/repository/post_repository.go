package repository

import (
	"context"

	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/valueobject"
)

type PostRepository interface {
	Save(ctx context.Context, post *entity.Post) (valueobject.PostID, error)
	FindByID(ctx context.Context, id valueobject.PostID) (*entity.Post, error)
	FindAll(ctx context.Context) ([]*entity.Post, error)
	Update(ctx context.Context, post *entity.Post) error
	Delete(ctx context.Context, id valueobject.PostID) error
}
