package delivery

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_userHandler_Signup(t *testing.T) {
	tests := []struct {
		name     string
		user     *domain.User
		mockFunc func(usecase *mock.MockUserUsecase)
	}{
		{
			name: "pass",
			user: &domain.User{
				Name:     "test",
				Password: "testtest",
			},
			mockFunc: func(usecase *mock.MockUserUsecase) {
				usecase.EXPECT().Signup(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			ctx := context.Background()
			usecase := mock.NewMockUserUsecase(ctrl)
			handler := userHandler{
				userUsecase: usecase,
			}

			test.mockFunc(usecase)

			userJson, err := json.Marshal(test.user)
			assert.NoError(t, err)

			req, err := http.NewRequestWithContext(ctx,
				http.MethodPost,
				"/backend/signup",
				bytes.NewBuffer(userJson),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, rec)

			err = handler.Signup(c)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusCreated, rec.Code)
		})
	}
}

func Test_userHandler_Login(t *testing.T) {
	tests := []struct {
		name     string
		user     *domain.User
		mockFunc func(usecase *mock.MockUserUsecase)
	}{
		{
			name: "pass",
			user: &domain.User{
				Name:     "test",
				Password: "testtest",
			},
			mockFunc: func(usecase *mock.MockUserUsecase) {
				usecase.EXPECT().Login(gomock.Any(), "test", "testtest").Times(1).Return(&domain.User{ID: 1}, nil)
				usecase.EXPECT().UpdateToken(gomock.Any(), 1, gomock.Any())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			ctx := context.Background()
			usecase := mock.NewMockUserUsecase(ctrl)
			handler := userHandler{
				userUsecase: usecase,
			}

			test.mockFunc(usecase)

			userJson, err := json.Marshal(test.user)
			assert.NoError(t, err)

			req, err := http.NewRequestWithContext(ctx,
				http.MethodPost,
				"/backend/login",
				bytes.NewBuffer(userJson),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, rec)

			err = handler.Login(c)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}
