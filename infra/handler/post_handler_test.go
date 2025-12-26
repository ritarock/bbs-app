package handler_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/ritarock/bbs-app/infra/handler"
	"github.com/stretchr/testify/assert"
)

func TestPostHandler_NewError(t *testing.T) {
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
			name:           "unknown error",
			err:            errors.New("database error"),
			wantStatusCode: http.StatusInternalServerError,
			wantMessage:    "Internal Server Error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			h := handler.NewPostHandler(nil, nil, nil, nil, nil)
			got := h.NewError(context.Background(), test.err)

			assert.Equal(t, test.wantStatusCode, got.StatusCode)
			assert.Equal(t, int32(test.wantStatusCode), got.Response.Code)
			assert.Equal(t, test.wantMessage, got.Response.Message)
		})
	}
}
