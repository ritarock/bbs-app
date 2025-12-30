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
	*AuthHandler
}

var _ api.Handler = (*Handler)(nil)

func NewHandler(postHandler *PostHandler, commentHandler *CommentHandler, authHandler *AuthHandler) *Handler {
	return &Handler{
		PostHandler:    postHandler,
		CommentHandler: commentHandler,
		AuthHandler:    authHandler,
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
		case errMsg == "user not found":
			statusCode = http.StatusNotFound
			message = "User not found"
		case errMsg == "unauthorized":
			statusCode = http.StatusUnauthorized
			message = "Unauthorized"
		case errMsg == "invalid email or password":
			statusCode = http.StatusUnauthorized
			message = "Invalid email or password"
		case errMsg == "user already exists":
			statusCode = http.StatusConflict
			message = "User already exists"
		case strings.Contains(errMsg, "title") || strings.Contains(errMsg, "content") || strings.Contains(errMsg, "body"):
			statusCode = http.StatusBadRequest
			message = errMsg
		case strings.Contains(errMsg, "email") || strings.Contains(errMsg, "password"):
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
