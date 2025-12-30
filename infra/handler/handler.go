package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/ritarock/bbs-app/infra/handler/api"
)

type Handler struct {
	*PostHandler
	*CommentHandler
}

var _ api.Handler = (*Handler)(nil)

func NewHandler(postHandler *PostHandler, commentHandler *CommentHandler) *Handler {
	return &Handler{
		PostHandler:    postHandler,
		CommentHandler: commentHandler,
	}
}

func (h *Handler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	statusCode := http.StatusInternalServerError
	message := "Internal Server Error"

	if err != nil {
		errMsg := err.Error()
		switch {
		case errMsg == "post not found":
			statusCode = http.StatusNotFound
			message = "Post not found"
		case errMsg == "comment not found":
			statusCode = http.StatusNotFound
			message = "Comment not found"
		case strings.Contains(errMsg, "title") || strings.Contains(errMsg, "content") || strings.Contains(errMsg, "body"):
			statusCode = http.StatusBadRequest
			message = errMsg
		}
	}

	return &api.ErrorStatusCode{
		StatusCode: statusCode,
		Response: api.Error{
			Code:    int32(statusCode),
			Message: message,
		},
	}
}
