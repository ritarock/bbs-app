package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_userUsecase_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUserRepository(ctrl)
	timeout := 2 * time.Second
	usecase := NewUserUsecase(repo, timeout)

	tests := []struct {
		name     string
		user     domain.User
		mockFunc func(t *testing.T, repo *mock.MockUserRepository, user domain.User)
		hasError bool
	}{
		{
			name: "pass",
			user: domain.User{
				Id:       1,
				Name:     "user",
				Password: "password1234",
			},
			mockFunc: func(t *testing.T, repo *mock.MockUserRepository, user domain.User) {
				repo.EXPECT().FindUser(gomock.Any(), user.Name, user.Password).Times(1).Return(nil, domain.ErrNotFound)
				repo.EXPECT().Create(gomock.Any(), &user).Times(1).Return(nil)
			},
			hasError: false,
		},
		{
			name: "failed",
			user: domain.User{
				Id:       1,
				Name:     "user",
				Password: "password1234",
			},
			mockFunc: func(t *testing.T, repo *mock.MockUserRepository, user domain.User) {
				repo.EXPECT().FindUser(gomock.Any(), user.Name, user.Password).Times(1).Return(&user, nil)
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		test.mockFunc(t, repo, test.user)
		err := usecase.SignUp(context.Background(), &test.user)
		if test.hasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}

}

func Test_userUsecase_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUserRepository(ctrl)
	timeout := 2 * time.Second
	usecase := NewUserUsecase(repo, timeout)

	user := domain.User{
		Id:       1,
		Name:     "user",
		Password: "password1234",
	}

	tests := []struct {
		name      string
		user      domain.User
		mockFunc  func(t *testing.T, repo *mock.MockUserRepository, user domain.User)
		existUser bool
	}{
		{
			name: "pass: return true",
			user: user,
			mockFunc: func(t *testing.T, repo *mock.MockUserRepository, user domain.User) {
				repo.EXPECT().FindUser(gomock.Any(), user.Name, user.Password).Times(1).Return(&user, nil)
			},
			existUser: true,
		},
		{
			name: "pass: return false",
			user: user,
			mockFunc: func(t *testing.T, repo *mock.MockUserRepository, user domain.User) {
				repo.EXPECT().FindUser(gomock.Any(), user.Name, user.Password).Times(1).Return(nil, domain.ErrNotFound)
			},
			existUser: false,
		},
	}

	for _, test := range tests {
		test.mockFunc(t, repo, test.user)
		reslut, _ := usecase.Login(context.Background(), &test.user)
		assert.Equal(t, test.existUser, reslut)
	}
}

func Test_userUsecase_SetToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUserRepository(ctrl)
	timeout := 2 * time.Second
	usecase := NewUserUsecase(repo, timeout)

	tests := []struct {
		name     string
		user     domain.User
		userId   int
		token    string
		mockFunc func(t *testing.T, repo *mock.MockUserRepository, user domain.User)
		hasError bool
	}{
		{
			name: "pass",
			user: domain.User{
				Id:    1,
				Token: "token",
			},
			userId: 1,
			token:  "token",
			mockFunc: func(t *testing.T, repo *mock.MockUserRepository, user domain.User) {
				repo.EXPECT().SetToken(gomock.Any(), 1, "token").Times(1).Return(nil)
			},
			hasError: false,
		},
		{
			name: "pass",
			user: domain.User{
				Id:    1,
				Token: "token",
			},
			userId: 10,
			token:  "token",
			mockFunc: func(t *testing.T, repo *mock.MockUserRepository, user domain.User) {
				repo.EXPECT().SetToken(gomock.Any(), 10, "token").Times(1).Return(errors.New(""))
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		test.mockFunc(t, repo, test.user)
		err := usecase.SetToken(context.Background(), test.userId, test.token)
		if test.hasError {

		} else {
			assert.NoError(t, err)
		}
	}
}

func Test_userUsecase_ValidateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUserRepository(ctrl)
	timeout := 2 * time.Second
	usecase := NewUserUsecase(repo, timeout)

	user := domain.User{
		Token: "token",
	}

	tests := []struct {
		name       string
		user       domain.User
		token      string
		mockFunc   func(t *testing.T, repo *mock.MockUserRepository, user domain.User)
		existToken bool
	}{
		{
			name:  "pass: return true",
			user:  user,
			token: "token",
			mockFunc: func(t *testing.T, repo *mock.MockUserRepository, user domain.User) {
				repo.EXPECT().ExistToken(gomock.Any(), "token").Times(1).Return(true)
			},
			existToken: true,
		},
		{
			name: "pass: return false",
			user: user,
			mockFunc: func(t *testing.T, repo *mock.MockUserRepository, user domain.User) {
				repo.EXPECT().ExistToken(gomock.Any(), "").Times(1).Return(false)
			},
			existToken: false,
		},
	}

	for _, test := range tests {
		test.mockFunc(t, repo, test.user)
		reslut := usecase.ValidateToken(context.Background(), test.token)
		assert.Equal(t, test.existToken, reslut)
	}

}
