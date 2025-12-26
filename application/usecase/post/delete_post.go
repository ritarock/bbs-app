package post

import (
	"context"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/domain/repository"
	"github.com/ritarock/bbs-app/domain/valueobject"
)

type DeletePostUsecase struct {
	postRepo repository.PostRepository
}

func NewDeletePostUsecase(postRepo repository.PostRepository) *DeletePostUsecase {
	return &DeletePostUsecase{
		postRepo: postRepo,
	}
}

func (u *DeletePostUsecase) Execute(ctx context.Context, input dto.DeletePostInput) error {
	postID := valueobject.NewPostID(input.ID)
	return u.postRepo.Delete(ctx, postID)
}
