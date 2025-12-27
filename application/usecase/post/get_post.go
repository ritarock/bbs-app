package post

import (
	"context"
	"errors"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/domain/repository"
	"github.com/ritarock/bbs-app/domain/valueobject"
)

type GetPostUsecase struct {
	postRepo repository.PostRepository
}

func NewGetPostUsecase(postRepo repository.PostRepository) *GetPostUsecase {
	return &GetPostUsecase{
		postRepo: postRepo,
	}
}

func (u *GetPostUsecase) Execute(ctx context.Context, input dto.GetPostInput) (*dto.GetPostOutput, error) {
	postID := valueobject.NewPostID(input.ID)

	post, err := u.postRepo.FindByID(ctx, postID)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, errors.New("post not found")
	}

	return &dto.GetPostOutput{
		ID:       post.ID().Int(),
		Title:    post.Title().String(),
		Content:  post.Content().String(),
		PostedAt: post.PostedAt(),
	}, nil
}
