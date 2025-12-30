package handler

import (
	"context"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/application/usecase/post"
	"github.com/ritarock/bbs-app/infra/handler/api"
)

type PostHandler struct {
	createPostUsecase *post.CreatePostUsecase
	getPostUsecase    *post.GetPostUsecase
	listPostUsecase   *post.ListPostUsecase
	updatePostUsecase *post.UpdatePostUsecase
	deletePostUsecase *post.DeletePostUsecase
}

func NewPostHandler(
	createPostUsecase *post.CreatePostUsecase,
	getPostUsecase *post.GetPostUsecase,
	listPostUsecase *post.ListPostUsecase,
	updatePostUsecase *post.UpdatePostUsecase,
	deletePostUsecase *post.DeletePostUsecase,
) *PostHandler {
	return &PostHandler{
		createPostUsecase: createPostUsecase,
		getPostUsecase:    getPostUsecase,
		listPostUsecase:   listPostUsecase,
		updatePostUsecase: updatePostUsecase,
		deletePostUsecase: deletePostUsecase,
	}
}

func (p *PostHandler) PostsCreate(ctx context.Context, req *api.CreatePostRequest) (*api.Post, error) {
	input := dto.CreatePostInput{
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}

	output, err := p.createPostUsecase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.Post{
		ID:       int64(output.ID),
		Title:    output.Title,
		Content:  output.Content,
		PostedAt: output.PostedAt,
	}, nil
}

func (p *PostHandler) PostsRead(ctx context.Context, params api.PostsReadParams) (*api.Post, error) {
	input := dto.GetPostInput{
		ID: int(params.ID),
	}

	output, err := p.getPostUsecase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.Post{
		ID:       int64(output.ID),
		Title:    output.Title,
		Content:  output.Content,
		PostedAt: output.PostedAt,
	}, nil
}

func (p *PostHandler) PostsList(ctx context.Context) (*api.PostList, error) {
	output, err := p.listPostUsecase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]api.Post, len(output.Posts))
	for i, post := range output.Posts {
		items[i] = api.Post{
			ID:       int64(post.ID),
			Title:    post.Title,
			Content:  post.Content,
			PostedAt: post.PostedAt,
		}
	}

	return &api.PostList{Items: items}, nil
}

func (p *PostHandler) PostsUpdate(ctx context.Context, req *api.UpdatePostRequest, params api.PostsUpdateParams) (*api.Post, error) {
	input := dto.UpdatePostInput{
		ID:      int(params.ID),
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}

	output, err := p.updatePostUsecase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.Post{
		ID:       int64(output.ID),
		Title:    output.Title,
		Content:  output.Content,
		PostedAt: output.PostedAt,
	}, nil
}

func (p *PostHandler) PostsDelete(ctx context.Context, params api.PostsDeleteParams) error {
	input := dto.DeletePostInput{
		ID: int(params.ID),
	}

	return p.deletePostUsecase.Execute(ctx, input)
}
