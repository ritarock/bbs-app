package http_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	commentHttp "ritarock/bbs-app/comment/delivery/http"
	"ritarock/bbs-app/domain"
	"ritarock/bbs-app/domain/mocks"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockPost := domain.Post{
		ID: 1,
	}
	mockComment := domain.Comment{
		ID:          1,
		Content:     "content 1",
		CommentedAt: time.Now(),
		PostID:      mockPost.ID,
	}
	mockUcase := new(mocks.CommentUsecase)

	j, err := json.Marshal(mockComment)
	assert.NoError(t, err)

	mockUcase.
		On("Create", mock.Anything, mock.AnythingOfType("*domain.Comment")).
		Return(nil)

	req, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		"/backend/api/v1/comments",
		strings.NewReader(string(j)),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	got := httptest.NewRecorder()

	handler := commentHttp.CommentHandler{
		CUsecase: mockUcase,
	}

	handler.Create(got, req)
	assert.Equal(t, http.StatusCreated, got.Code)
	mockUcase.AssertExpectations(t)
}

func TestGetAllByPost(t *testing.T) {
	mockPost := domain.Post{
		ID: 1,
	}
	mockComments := []domain.Comment{
		{
			ID:          1,
			Content:     "content 1",
			CommentedAt: time.Now(),
			PostID:      mockPost.ID,
		},
		{
			ID:          2,
			Content:     "content 2",
			CommentedAt: time.Now(),
			PostID:      mockPost.ID,
		},
	}
	mockUcase := new(mocks.CommentUsecase)

	mockUcase.
		On("GetAllByPost", mock.Anything, mockPost.ID).
		Return(mockComments, nil)

	req, err := http.NewRequestWithContext(
		context.Background(),
		"GET",
		"/backend/api/v1/comments/1",
		strings.NewReader(""),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	got := httptest.NewRecorder()

	handler := commentHttp.CommentHandler{
		CUsecase: mockUcase,
	}

	handler.GetAllByPost(got, req)
	assert.Equal(t, http.StatusOK, got.Code)
	mockUcase.AssertExpectations(t)
}
