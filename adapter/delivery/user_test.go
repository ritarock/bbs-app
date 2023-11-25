package delivery

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_userHandler_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	usecase := mock.NewMockUserUsecase(ctrl)
	user := domain.User{
		Name:     "user",
		Password: "password1234",
	}

	j, err := json.Marshal(user)
	assert.NoError(t, err)

	usecase.EXPECT().SignUp(ctx, &user).Times(1).Return(nil)

	e := echo.New()
	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		"/backend/signup",
		strings.NewReader(string(j)),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/backend/signup")

	handler := userHandler{
		userUsecase: usecase,
	}
	err = handler.SignUp(c)
	assert.NoError(t, err)
}

func Test_userHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	usecase := mock.NewMockUserUsecase(ctrl)
	user := domain.User{
		Name:     "user",
		Password: "password1234",
	}

	j, err := json.Marshal(user)
	assert.NoError(t, err)

	usecase.EXPECT().Login(ctx, &user).Times(1).Return(true, &user)
	usecase.EXPECT().SetToken(ctx, gomock.Any(), gomock.Any()).Times(1).Return(nil)

	e := echo.New()
	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		"/backend/login",
		strings.NewReader(string(j)),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/backend/login")

	handler := userHandler{
		userUsecase: usecase,
	}
	err = handler.Login(c)
	assert.NoError(t, err)
}
