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

func Test_postHandler_Create(t *testing.T) {
	tests := []struct {
		name     string
		post     *domain.Post
		mockFunc func(usecase *mock.MockPostUsecase)
	}{
		{
			name: "pass",
			post: &domain.Post{
				Title:   "test",
				Content: "test",
			},
			mockFunc: func(usecase *mock.MockPostUsecase) {
				usecase.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			ctx := context.Background()
			usecase := mock.NewMockPostUsecase(ctrl)
			handler := postHandler{
				postUsecase: usecase,
			}

			test.mockFunc(usecase)

			postJson, err := json.Marshal(test.post)
			assert.NoError(t, err)

			req, err := http.NewRequestWithContext(ctx,
				http.MethodPost,
				"/backend/api/v1/posts",
				bytes.NewBuffer(postJson),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, rec)

			err = handler.Create(c)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusCreated, rec.Code)
		})
	}
}

func Test_postHandler_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		mockFunc func(usecase *mock.MockPostUsecase)
	}{
		{
			name: "pass",
			mockFunc: func(usecase *mock.MockPostUsecase) {
				usecase.EXPECT().GetAll(gomock.Any()).Times(1).Return(
					[]domain.Post{
						{
							ID:       1,
							Title:    "test1",
							Content:  "test1",
							PostedAt: time.Now(),
						},
						{
							ID:       2,
							Title:    "test2",
							Content:  "test2",
							PostedAt: time.Now(),
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
			usecase := mock.NewMockPostUsecase(ctrl)
			handler := postHandler{
				postUsecase: usecase,
			}

			test.mockFunc(usecase)

			req, err := http.NewRequestWithContext(ctx,
				http.MethodGet,
				"/backend/api/v1/posts/",
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

func Test_postHandler_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		postID   string
		mockFunc func(usecase *mock.MockPostUsecase)
	}{
		{
			name:   "pass",
			postID: "1",
			mockFunc: func(usecase *mock.MockPostUsecase) {
				usecase.EXPECT().GetByID(gomock.Any(), 1).Times(1).Return(&domain.Post{
					ID:       1,
					Title:    "test",
					Content:  "test",
					PostedAt: time.Now(),
				}, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			ctx := context.Background()
			usecase := mock.NewMockPostUsecase(ctrl)
			handler := postHandler{
				postUsecase: usecase,
			}

			test.mockFunc(usecase)

			req, err := http.NewRequestWithContext(ctx,
				http.MethodGet,
				"/backend/api/v1/posts/"+test.postID,
				nil,
			)
			assert.NoError(t, err)

			rec := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/backend/api/v1/posts/:id")
			c.SetParamNames("id")
			c.SetParamValues(test.postID)

			err = handler.GetByID(c)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}
