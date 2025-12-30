package comment

import (
	"context"
	"errors"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/domain/repository"
	"github.com/ritarock/bbs-app/domain/valueobject"
)

type GetCommentUsecase struct {
	commentRepo repository.CommentRepository
}

func NewGetCommentUsecase(commentRepo repository.CommentRepository) *GetCommentUsecase {
	return &GetCommentUsecase{
		commentRepo: commentRepo,
	}
}

func (u *GetCommentUsecase) Execute(ctx context.Context, input dto.GetCommentInput) (*dto.GetCommentOutput, error) {
	commentID := valueobject.NewCommentID(input.ID)

	comment, err := u.commentRepo.FindByID(ctx, commentID)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, errors.New("comment not found")
	}

	return &dto.GetCommentOutput{
		ID:          comment.ID().Int(),
		PostID:      comment.PostID().Int(),
		Body:        comment.Body().String(),
		CommentedAt: comment.CommentedAt(),
	}, nil
}
