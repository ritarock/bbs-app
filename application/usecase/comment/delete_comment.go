package comment

import (
	"context"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/domain/repository"
	"github.com/ritarock/bbs-app/domain/valueobject"
)

type DeleteCommentUsecase struct {
	commentRepo repository.CommentRepository
}

func NewDeleteCommentUsecase(commentRepo repository.CommentRepository) *DeleteCommentUsecase {
	return &DeleteCommentUsecase{
		commentRepo: commentRepo,
	}
}

func (u *DeleteCommentUsecase) Execute(ctx context.Context, input dto.DeleteCommentInput) error {
	commentID := valueobject.NewCommentID(input.ID)
	return u.commentRepo.Delete(ctx, commentID)
}
