package delivery

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ritarock/bbs-app/domain"
	timeutil "github.com/ritarock/bbs-app/internal/time_util"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_postHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	usecase := mock.NewMockPostUsecase(ctrl)
	ctx = timeutil.SetMockNow(ctx, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local))
	post := &domain.Post{
		Id:       1,
		Title:    "title",
		Content:  "content",
		PostedAt: timeutil.Now(ctx),
	}

	j, err := json.Marshal(post)
	assert.NoError(t, err)

	usecase.EXPECT().Create(ctx, post).Times(1).Return(nil)

	e := echo.New()
	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		"/backend/api/v1/posts",
		strings.NewReader(string(j)),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/backend/api/v1/posts")

	handler := postHandler{
		postUsecase: usecase,
	}
	err = handler.Create(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func Test_postHandler_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	usecase := mock.NewMockPostUsecase(ctrl)

	posts := []domain.Post{
		{
			Id:      1,
			Title:   "title1",
			Content: "content1",
		},
		{
			Id:      2,
			Title:   "title2",
			Content: "content2",
		},
	}

	usecase.EXPECT().GetAll(ctx).Times(1).Return(posts, nil)

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
	handler := postHandler{
		postUsecase: usecase,
	}

	handler.GetAll(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func Test_postHandler_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	usecase := mock.NewMockPostUsecase(ctrl)
	post := domain.Post{
		Id:      1,
		Title:   "title1",
		Content: "content1",
	}

	usecase.EXPECT().GetById(ctx, post.Id).Times(1).Return(&post, nil)

	e := echo.New()
	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		"/backend/api/v1/posts",
		strings.NewReader(""),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/backend/api/v1/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(post.Id))
	handler := postHandler{
		postUsecase: usecase,
	}

	handler.GetById(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
