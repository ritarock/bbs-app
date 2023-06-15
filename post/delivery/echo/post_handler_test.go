package echo_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ritarock/bbs-app/domain"
	"ritarock/bbs-app/domain/mocks"
	postEcho "ritarock/bbs-app/post/delivery/echo"
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

	e := echo.New()
	req, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		"/backend/api/v1/posts",
		strings.NewReader(string(j)),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/backend/api/v1/posts")

	handler := postEcho.PostHandler{
		PUsecase: mockUcase,
	}

	handler.Create(c)
	assert.Equal(t, http.StatusCreated, rec.Code)
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
	c.SetPath("/backend/api/v1/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(mockPost.ID))
	handler := postEcho.PostHandler{
		PUsecase: mockUcase,
	}

	handler.GetById(c)
	assert.Equal(t, http.StatusOK, rec.Code)
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

	e := echo.New()
	req, err := http.NewRequestWithContext(
		context.Background(),
		"GET",
		"/backend/api/v1/posts",
		strings.NewReader(""),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/backend/api/v1/posts")
	handler := postEcho.PostHandler{
		PUsecase: mockUcase,
	}

	handler.GetAll(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUcase.AssertExpectations(t)
}
