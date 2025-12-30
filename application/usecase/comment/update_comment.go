package comment

import (
	"context"
	"errors"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/domain/repository"
	"github.com/ritarock/bbs-app/domain/valueobject"
)

type UpdateCommentUsecase struct {
	commentRepo repository.CommentRepository
}

func NewUpdateCommentUsecase(commentRepo repository.CommentRepository) *UpdateCommentUsecase {
	return &UpdateCommentUsecase{
		commentRepo: commentRepo,
	}
}

func (u *UpdateCommentUsecase) Execute(ctx context.Context, input dto.UpdateCommentInput) (*dto.UpdateCommentOutput, error) {
	commentID := valueobject.NewCommentID(input.ID)

	comment, err := u.commentRepo.FindByID(ctx, commentID)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, errors.New("comment not found")
	}

	if err := comment.Update(input.Body); err != nil {
		return nil, err
	}

	if err := u.commentRepo.Update(ctx, comment); err != nil {
		return nil, err
	}

	return &dto.UpdateCommentOutput{
		ID:          commentID.Int(),
		PostID:      comment.PostID().Int(),
		Body:        comment.Body().String(),
		CommentedAt: comment.CommentedAt(),
	}, nil
}
