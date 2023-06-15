package echo_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	commentEcho "ritarock/bbs-app/comment/delivery/echo"
	"ritarock/bbs-app/domain"
	"ritarock/bbs-app/domain/mocks"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
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

	e := echo.New()
	req, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		"/backend/api/v1/comment",
		strings.NewReader(string(j)),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/backend/api/v1/comment")

	handler := commentEcho.CommentHandler{
		CUsecase: mockUcase,
	}

	handler.Create(c)
	assert.Equal(t, http.StatusCreated, rec.Code)
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

	e := echo.New()
	req, err := http.NewRequestWithContext(
		context.Background(),
		"GET",
		"/backend/api/v1/posts/1",
		strings.NewReader(""),
	)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/backend/api/v1/comments/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(mockPost.ID))
	handler := commentEcho.CommentHandler{
		CUsecase: mockUcase,
	}

	handler.GetAllByPost(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUcase.AssertExpectations(t)
}
