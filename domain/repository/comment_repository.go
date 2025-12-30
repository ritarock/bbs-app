package repository

import (
	"context"

	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/valueobject"
)

type CommentRepository interface {
	Save(ctx context.Context, comment *entity.Comment) (valueobject.CommentID, error)
	FindByID(ctx context.Context, id valueobject.CommentID) (*entity.Comment, error)
	FindAll(ctx context.Context) ([]*entity.Comment, error)
	FindByPostID(ctx context.Context, postID valueobject.PostID) ([]*entity.Comment, error)
	Update(ctx context.Context, comment *entity.Comment) error
	Delete(ctx context.Context, id valueobject.CommentID) error
}
