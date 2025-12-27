package post

import (
	"context"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/repository"
)

type CreatePostUsecase struct {
	postRepo repository.PostRepository
}

func NewCreatePostUsecase(postRepo repository.PostRepository) *CreatePostUsecase {
	return &CreatePostUsecase{
		postRepo: postRepo,
	}
}

func (u *CreatePostUsecase) Execute(ctx context.Context, input dto.CreatePostInput) (*dto.CreatePostOutput, error) {
	post, err := entity.NewPost(input.Title, input.Content)
	if err != nil {
		return nil, err
	}

	postID, err := u.postRepo.Save(ctx, post)
	if err != nil {
		return nil, err
	}

	return &dto.CreatePostOutput{
		ID:       postID.Int(),
		Title:    post.Title().String(),
		Content:  post.Content().String(),
		PostedAt: post.PostedAt(),
	}, nil
}
