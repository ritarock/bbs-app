package http_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ritarock/bbs-app/domain"
	"ritarock/bbs-app/domain/mocks"
	postHttp "ritarock/bbs-app/post/delivery/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockPost := domain.Post{
		ID:       1,
		Title:    "title 1",
		Content:  "content 1",
		PostedAt: time.Now(),
	}
	mockUcase := new(mocks.PostUsecase)

	j, err := json.Marshal(mockPost)
	assert.NoError(t, err)

	mockUcase.
		On("Create", mock.Anything, mock.AnythingOfType("*domain.Post")).
		Return(nil)

	req, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		"/backend/api/v1/posts",
		strings.NewReader(string(j)),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	got := httptest.NewRecorder()

	handler := postHttp.PostHandler{
		PUsecase: mockUcase,
	}

	handler.Create(got, req)
	assert.Equal(t, http.StatusCreated, got.Code)
	mockUcase.AssertExpectations(t)
}

func TestGetById(t *testing.T) {
	mockPost := domain.Post{
		ID:       1,
		Title:    "title 1",
		Content:  "content 1",
		PostedAt: time.Now(),
	}
	mockUcase := new(mocks.PostUsecase)

	mockUcase.
		On("GetById", mock.Anything, mockPost.ID).
		Return(mockPost, nil)

	req, err := http.NewRequestWithContext(
		context.Background(),
		"GET",
		"/backend/api/v1/posts/1",
		strings.NewReader(""),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	got := httptest.NewRecorder()

	handler := postHttp.PostHandler{
		PUsecase: mockUcase,
	}

	handler.GetByID(got, req)
	assert.Equal(t, http.StatusOK, got.Code)
	mockUcase.AssertExpectations(t)
}

func TestGetAll(t *testing.T) {
	mockPosts := []domain.Post{
		{
			ID:       1,
			Title:    "title 1",
			Content:  "content 1",
			PostedAt: time.Now(),
		},
		{
			ID:       2,
			Title:    "title 2",
			Content:  "content 2",
			PostedAt: time.Now(),
		},
	}
	mockUcase := new(mocks.PostUsecase)

	mockUcase.
		On("GetAll", mock.Anything).
		Return(mockPosts, nil)

	req, err := http.NewRequestWithContext(
		context.Background(),
		"GET",
		"/backend/api/v1/posts",
		strings.NewReader(""),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	got := httptest.NewRecorder()

	handler := postHttp.PostHandler{
		PUsecase: mockUcase,
	}

	handler.GetAll(got, req)
	assert.Equal(t, http.StatusOK, got.Code)
	mockUcase.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockPost := domain.Post{
		ID:       1,
		Title:    "title 1",
		Content:  "content 1",
		PostedAt: time.Now(),
	}
	mockUcase := new(mocks.PostUsecase)

	j, err := json.Marshal(mockPost)
	assert.NoError(t, err)

	mockUcase.
		On("GetById", mock.Anything, mockPost.ID).
		Return(mockPost, nil)
	mockUcase.
		On("Update", mock.Anything, &mockPost).
		Return(nil)

	req, err := http.NewRequestWithContext(
		context.Background(),
		"PUT",
		"/backend/api/v1/posts/1",
		strings.NewReader(string(j)),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	got := httptest.NewRecorder()

	handler := postHttp.PostHandler{
		PUsecase: mockUcase,
	}

	handler.Update(got, req)
	assert.Equal(t, http.StatusOK, got.Code)
	mockUcase.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockPost := domain.Post{
		ID:       1,
		Title:    "title 1",
		Content:  "content 1",
		PostedAt: time.Now(),
	}
	mockUcase := new(mocks.PostUsecase)

	mockUcase.On("Delete", mock.Anything, mockPost.ID).Return(nil)

	req, err := http.NewRequestWithContext(
		context.Background(),
		"DELETE",
		"/backend/api/v1/posts/1",
		strings.NewReader(""),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	got := httptest.NewRecorder()

	handler := postHttp.PostHandler{
		PUsecase: mockUcase,
	}

	handler.Delete(got, req)
	assert.Equal(t, http.StatusNoContent, got.Code)
	mockUcase.AssertExpectations(t)
}
