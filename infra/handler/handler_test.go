package handler_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/ritarock/bbs-app/infra/handler"
	"github.com/stretchr/testify/assert"
)

func TestHandler_NewError(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		err            error
		wantStatusCode int
		wantMessage    string
	}{
		{
			name:           "post not found",
			err:            errors.New("post not found"),
			wantStatusCode: http.StatusNotFound,
			wantMessage:    "Post not found",
		},
		{
			name:           "comment not found",
			err:            errors.New("comment not found"),
			wantStatusCode: http.StatusNotFound,
			wantMessage:    "Comment not found",
		},
		{
			name:           "user not found",
			err:            errors.New("user not found"),
			wantStatusCode: http.StatusNotFound,
			wantMessage:    "User not found",
		},
		{
			name:           "unauthorized",
			err:            errors.New("unauthorized"),
			wantStatusCode: http.StatusUnauthorized,
			wantMessage:    "Unauthorized",
		},
		{
			name:           "invalid email or password",
			err:            errors.New("invalid email or password"),
			wantStatusCode: http.StatusUnauthorized,
			wantMessage:    "Invalid email or password",
		},
		{
			name:           "user already exists",
			err:            errors.New("user already exists"),
			wantStatusCode: http.StatusConflict,
			wantMessage:    "User already exists",
		},
		{
			name:           "title validation error",
			err:            errors.New("title must be between 1 and 30 characters"),
			wantStatusCode: http.StatusBadRequest,
			wantMessage:    "title must be between 1 and 30 characters",
		},
		{
			name:           "content validation error",
			err:            errors.New("content must be between 1 and 255 characters"),
			wantStatusCode: http.StatusBadRequest,
			wantMessage:    "content must be between 1 and 255 characters",
		},
		{
			name:           "body validation error",
			err:            errors.New("body must be at least 1 character"),
			wantStatusCode: http.StatusBadRequest,
			wantMessage:    "body must be at least 1 character",
		},
		{
			name:           "email validation error",
			err:            errors.New("email is required"),
			wantStatusCode: http.StatusBadRequest,
			wantMessage:    "email is required",
		},
		{
			name:           "password validation error",
			err:            errors.New("password must be at least 8 characters"),
			wantStatusCode: http.StatusBadRequest,
			wantMessage:    "password must be at least 8 characters",
		},
		{
			name:           "unknown error",
			err:            errors.New("database error"),
			wantStatusCode: http.StatusInternalServerError,
			wantMessage:    "Internal Server Error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			h := handler.NewHandler(nil, nil, nil)
			got := h.NewError(context.Background(), test.err)

			assert.Equal(t, test.wantStatusCode, got.StatusCode)
			assert.Equal(t, int32(test.wantStatusCode), got.Response.Code)
			assert.Equal(t, test.wantMessage, got.Response.Message)
		})
	}
}
