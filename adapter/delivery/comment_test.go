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

func Test_commentHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	ctx = timeutil.SetMockNow(ctx, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local))
	usecase := mock.NewMockCommentUsecase(ctrl)
	comment := &domain.Comment{
		Id:          1,
		Content:     "content",
		CommentedAt: timeutil.Now(ctx),
		PostId:      1,
	}

	j, err := json.Marshal(comment)
	assert.NoError(t, err)

	usecase.EXPECT().Create(ctx, comment).Times(1).Return(nil)

	e := echo.New()
	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		"/backend/api/v1/post/1/comments",
		strings.NewReader(string(j)),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/backend/api/v1/post/:id/comments")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(comment.PostId))

	handler := commentHandler{
		commentUsecase: usecase,
	}
	err = handler.Create(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func Test_commentHandler_GetAllByPostId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	usecase := mock.NewMockCommentUsecase(ctrl)
	comments := []domain.Comment{
		{
			Id:      1,
			Content: "content1",
			PostId:  1,
		},
		{
			Id:      2,
			Content: "content2",
			PostId:  1,
		},
	}

	usecase.EXPECT().GetByPostId(ctx, 1).Times(1).Return(comments, nil)

	e := echo.New()
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		"/backend/api/v1/post/1/comments",
		strings.NewReader(""),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/backend/api/v1/post/:id/comments")
	c.SetParamNames("id")
	c.SetParamValues("1")
	handler := commentHandler{
		commentUsecase: usecase,
	}

	handler.GetAllByPostId(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
