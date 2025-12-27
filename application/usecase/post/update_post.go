package post

import (
	"context"
	"errors"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/domain/repository"
	"github.com/ritarock/bbs-app/domain/valueobject"
)

type UpdatePostUsecase struct {
	postRepo repository.PostRepository
}

func NewUpdatePostUsecase(postRepo repository.PostRepository) *UpdatePostUsecase {
	return &UpdatePostUsecase{
		postRepo: postRepo,
	}
}

func (u *UpdatePostUsecase) Execute(ctx context.Context, input dto.UpdatePostInput) (*dto.UpdatePostOutput, error) {
	postID := valueobject.NewPostID(input.ID)

	post, err := u.postRepo.FindByID(ctx, postID)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, errors.New("post not found")
	}

	if err := post.Update(input.Title, input.Content); err != nil {
		return nil, err
	}

	if err := u.postRepo.Update(ctx, post); err != nil {
		return nil, err
	}

	return &dto.UpdatePostOutput{
		ID:       postID.Int(),
		Title:    post.Title().String(),
		Content:  post.Content().String(),
		PostedAt: post.PostedAt(),
	}, nil
}
