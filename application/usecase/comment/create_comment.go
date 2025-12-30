package comment

import (
	"context"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/repository"
)

type CreateCommentUsecase struct {
	commentRepo repository.CommentRepository
}

func NewCreateCommentUsecase(commentRepo repository.CommentRepository) *CreateCommentUsecase {
	return &CreateCommentUsecase{
		commentRepo: commentRepo,
	}
}

func (u *CreateCommentUsecase) Execute(ctx context.Context, input dto.CreateCommentInput) (*dto.CreateCommentOutput, error) {
	comment, err := entity.NewComment(input.PostID, input.Body)
	if err != nil {
		return nil, err
	}

	commentID, err := u.commentRepo.Save(ctx, comment)
	if err != nil {
		return nil, err
	}

	return &dto.CreateCommentOutput{
		ID:          commentID.Int(),
		PostID:      comment.PostID().Int(),
		Body:        comment.Body().String(),
		CommentedAt: comment.CommentedAt(),
	}, nil
}
