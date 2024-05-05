package delivery

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_commentHandler_Create(t *testing.T) {
	tests := []struct {
		name     string
		comment  *domain.Comment
		postID   string
		mockFunc func(usecase *mock.MockCommentUsecase)
	}{
		{
			name: "pass",
			comment: &domain.Comment{
				Content: "test",
			},
			postID: "1",
			mockFunc: func(usecase *mock.MockCommentUsecase) {
				usecase.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			ctx := context.Background()
			usecase := mock.NewMockCommentUsecase(ctrl)
			handler := commentHandler{
				commentUsecase: usecase,
			}

			test.mockFunc(usecase)

			commentJson, err := json.Marshal(test.comment)
			assert.NoError(t, err)

			req, err := http.NewRequestWithContext(ctx,
				http.MethodPost,
				"/backend/api/v1/post/"+test.postID+"/comments",
				bytes.NewBuffer(commentJson),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/backend/api/v1/post/:post_id/comments")
			c.SetParamNames("post_id")
			c.SetParamValues(test.postID)

			err = handler.Create(c)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusCreated, rec.Code)
		})
	}
}

func Test_commentHandler_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		mockFunc func(usecase *mock.MockCommentUsecase)
	}{
		{
			name: "pass",
			mockFunc: func(usecase *mock.MockCommentUsecase) {
				usecase.EXPECT().GetAll(gomock.Any()).Times(1).Return(
					[]domain.Comment{
						{
							ID:          1,
							PostID:      1,
							Content:     "test1",
							CommentedAt: time.Now(),
						},
						{
							ID:          2,
							PostID:      1,
							Content:     "test2",
							CommentedAt: time.Now(),
						},
					}, nil,
				)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			ctx := context.Background()
			usecase := mock.NewMockCommentUsecase(ctrl)
			handler := commentHandler{
				commentUsecase: usecase,
			}

			test.mockFunc(usecase)

			req, err := http.NewRequestWithContext(ctx,
				http.MethodGet,
				"/backend/api/v1/comments",
				nil,
			)
			assert.NoError(t, err)

			rec := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, rec)

			err = handler.GetAll(c)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func Test_commentHandler_GetByPostID(t *testing.T) {
	tests := []struct {
		name     string
		postID   string
		mockFunc func(usecase *mock.MockCommentUsecase)
	}{
		{
			name:   "pass",
			postID: "1",
			mockFunc: func(usecase *mock.MockCommentUsecase) {
				usecase.EXPECT().GetByPostID(gomock.Any(), 1).Times(1).Return(
					[]domain.Comment{
						{
							ID:          1,
							PostID:      1,
							Content:     "test1",
							CommentedAt: time.Now(),
						},
						{
							ID:          2,
							PostID:      1,
							Content:     "test2",
							CommentedAt: time.Now(),
						},
					}, nil,
				)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			ctx := context.Background()
			usecase := mock.NewMockCommentUsecase(ctrl)
			handler := commentHandler{
				commentUsecase: usecase,
			}

			test.mockFunc(usecase)

			req, err := http.NewRequestWithContext(ctx,
				http.MethodGet,
				"/backend/api/v1/post/"+test.postID+"/comments",
				nil,
			)
			assert.NoError(t, err)

			rec := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/backend/api/v1/post/:post_id/comments")
			c.SetParamNames("post_id")
			c.SetParamValues(test.postID)

			err = handler.GetByPostID(c)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}

}
