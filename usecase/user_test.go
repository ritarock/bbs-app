package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_userUsecase_Signup(t *testing.T) {
	tests := []struct {
		name     string
		user     *domain.User
		mockFunc func(repo *mock.MockUserRepository)
		hasError bool
	}{
		{
			name: "pass",
			user: &domain.User{
				Name:     "test",
				Password: "test",
			},
			mockFunc: func(repo *mock.MockUserRepository) {
				repo.EXPECT().
					GetByNameAndPassword(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).Return(nil, domain.ErrNotFound)
				repo.EXPECT().Create(gomock.Any(), &domain.User{
					Name:     "test",
					Password: "test",
				}).Times(1).Return(nil)
			},
			hasError: false,
		},
		{
			name: "failed: already user exists",
			user: &domain.User{
				Name:     "test",
				Password: "test",
			},
			mockFunc: func(repo *mock.MockUserRepository) {
				repo.EXPECT().
					GetByNameAndPassword(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).Return(&domain.User{}, nil)
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			repo := mock.NewMockUserRepository(ctrl)
			timeout := 2 * time.Second
			usecase := NewUserUsecase(repo, timeout)
			test.mockFunc(repo)

			err := usecase.Signup(context.Background(), test.user)
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_userUsecase_Login(t *testing.T) {
	tests := []struct {
		name     string
		mockFunc func(repo *mock.MockUserRepository)
	}{
		{
			name: "pass",
			mockFunc: func(repo *mock.MockUserRepository) {
				repo.EXPECT().
					GetByNameAndPassword(gomock.Any(), "test", "test").
					Times(1).Return(&domain.User{}, nil)
			},
		},
	}

	for _, test := range tests {
		ctrl := gomock.NewController(t)

		repo := mock.NewMockUserRepository(ctrl)
		timeout := 2 * time.Second
		usecase := NewUserUsecase(repo, timeout)
		test.mockFunc(repo)

		got, err := usecase.Login(context.Background(), "test", "test")
		assert.NoError(t, err)
		assert.NotNil(t, got)
	}
}

func Test_userUsecase_UpdateToken(t *testing.T) {
	tests := []struct {
		name     string
		mockFunc func(repo *mock.MockUserRepository)
	}{
		{
			name: "pass",
			mockFunc: func(repo *mock.MockUserRepository) {
				repo.EXPECT().UpdateToken(gomock.Any(), 1, "test").
					Times(1).Return(nil)
			},
		},
	}

	for _, test := range tests {
		ctrl := gomock.NewController(t)

		repo := mock.NewMockUserRepository(ctrl)
		timeout := 2 * time.Second
		usecase := NewUserUsecase(repo, timeout)
		test.mockFunc(repo)

		err := usecase.UpdateToken(context.Background(), 1, "test")
		assert.NoError(t, err)
	}
}

func Test_userUsecase_IsTokenAvailable(t *testing.T) {
	tests := []struct {
		name     string
		token    string
		mockFunc func(repo *mock.MockUserRepository)
	}{
		{
			name:  "pass",
			token: "test",
			mockFunc: func(repo *mock.MockUserRepository) {
				repo.EXPECT().ExistToken(gomock.Any(), "test").Times(1).Return(true, nil)
			},
		},
	}

	for _, test := range tests {
		ctrl := gomock.NewController(t)

		repo := mock.NewMockUserRepository(ctrl)
		timeout := 2 * time.Second
		usecase := NewUserUsecase(repo, timeout)
		test.mockFunc(repo)

		result := usecase.IsTokenAvailable(context.Background(), test.token)
		assert.Equal(t, true, result)
	}
}
