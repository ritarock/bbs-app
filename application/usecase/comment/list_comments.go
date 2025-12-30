package comment

import (
	"context"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/domain/repository"
	"github.com/ritarock/bbs-app/domain/valueobject"
)

type ListCommentsUsecase struct {
	commentRepo repository.CommentRepository
}

func NewListCommentsUsecase(commentRepo repository.CommentRepository) *ListCommentsUsecase {
	return &ListCommentsUsecase{
		commentRepo: commentRepo,
	}
}

func (u *ListCommentsUsecase) Execute(ctx context.Context, input dto.ListCommentsInput) (*dto.ListCommentsOutput, error) {
	postID := valueobject.NewPostID(input.PostID)

	comments, err := u.commentRepo.FindByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}

	items := make([]dto.CommentItem, len(comments))
	for i, comment := range comments {
		items[i] = dto.CommentItem{
			ID:          comment.ID().Int(),
			PostID:      comment.PostID().Int(),
			Body:        comment.Body().String(),
			CommentedAt: comment.CommentedAt(),
		}
	}

	return &dto.ListCommentsOutput{Comments: items}, nil
}
