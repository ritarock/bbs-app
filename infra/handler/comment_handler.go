package handler

import (
	"context"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/application/usecase/comment"
	"github.com/ritarock/bbs-app/infra/handler/api"
)

type CommentHandler struct {
	createCommentUsecase *comment.CreateCommentUsecase
	getCommentUsecase    *comment.GetCommentUsecase
	listCommentsUsecase  *comment.ListCommentsUsecase
	updateCommentUsecase *comment.UpdateCommentUsecase
	deleteCommentUsecase *comment.DeleteCommentUsecase
}

func NewCommentHandler(
	createCommentUsecase *comment.CreateCommentUsecase,
	getCommentUsecase *comment.GetCommentUsecase,
	listCommentsUsecase *comment.ListCommentsUsecase,
	updateCommentUsecase *comment.UpdateCommentUsecase,
	deleteCommentUsecase *comment.DeleteCommentUsecase,
) *CommentHandler {
	return &CommentHandler{
		createCommentUsecase: createCommentUsecase,
		getCommentUsecase:    getCommentUsecase,
		listCommentsUsecase:  listCommentsUsecase,
		updateCommentUsecase: updateCommentUsecase,
		deleteCommentUsecase: deleteCommentUsecase,
	}
}

func (h *CommentHandler) CommentsCreate(ctx context.Context, req *api.CreateCommentRequest, params api.CommentsCreateParams) (*api.Comment, error) {
	input := dto.CreateCommentInput{
		PostID: int(params.PostId),
		Body:   req.GetBody(),
	}

	output, err := h.createCommentUsecase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.Comment{
		ID:          int64(output.ID),
		PostId:      int64(output.PostID),
		Body:        output.Body,
		CommentedAt: output.CommentedAt,
	}, nil
}

func (h *CommentHandler) CommentsRead(ctx context.Context, params api.CommentsReadParams) (*api.Comment, error) {
	input := dto.GetCommentInput{
		ID: int(params.ID),
	}

	output, err := h.getCommentUsecase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.Comment{
		ID:          int64(output.ID),
		PostId:      int64(output.PostID),
		Body:        output.Body,
		CommentedAt: output.CommentedAt,
	}, nil
}

func (h *CommentHandler) CommentsList(ctx context.Context, params api.CommentsListParams) (*api.CommentList, error) {
	input := dto.ListCommentsInput{
		PostID: int(params.PostId),
	}

	output, err := h.listCommentsUsecase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	items := make([]api.Comment, len(output.Comments))
	for i, c := range output.Comments {
		items[i] = api.Comment{
			ID:          int64(c.ID),
			PostId:      int64(c.PostID),
			Body:        c.Body,
			CommentedAt: c.CommentedAt,
		}
	}

	return &api.CommentList{Items: items}, nil
}

func (h *CommentHandler) CommentsUpdate(ctx context.Context, req *api.UpdateCommentRequest, params api.CommentsUpdateParams) (*api.Comment, error) {
	input := dto.UpdateCommentInput{
		ID:   int(params.ID),
		Body: req.GetBody(),
	}

	output, err := h.updateCommentUsecase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.Comment{
		ID:          int64(output.ID),
		PostId:      int64(output.PostID),
		Body:        output.Body,
		CommentedAt: output.CommentedAt,
	}, nil
}

func (h *CommentHandler) CommentsDelete(ctx context.Context, params api.CommentsDeleteParams) error {
	input := dto.DeleteCommentInput{
		ID: int(params.ID),
	}

	return h.deleteCommentUsecase.Execute(ctx, input)
}
