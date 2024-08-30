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

func Test_userUsecase_Create(t *testing.T) {
	tests := []struct {
		name     string
		user     *domain.User
		mockFunc func(repo *mock.MockUserRepository)
	}{
		{
			name: "pass",
			user: &domain.User{
				Name:     "test name",
				Password: "test password",
			},
			mockFunc: func(repo *mock.MockUserRepository) {
				repo.EXPECT().Create(gomock.Any(), &domain.User{
					Name:     "test name",
					Password: "test password",
				}).Return(nil)
			},
		},
	}

	for _, test := range tests {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockUserRepository(ctrl)
		timeout := 2 * time.Second
		usecase := NewUserUsecase(repo, timeout)
		test.mockFunc(repo)

		err := usecase.Create(context.Background(), test.user)
		assert.NoError(t, err)
	}
}

func Test_userUsecase_Find(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		password string
		mockFunc func(repo *mock.MockUserRepository)
	}{
		{
			name:     "pass",
			userName: "test name",
			password: "test password",
			mockFunc: func(repo *mock.MockUserRepository) {
				repo.EXPECT().
					GetByNameAndPasswd(gomock.Any(), "test name", "test password").
					Return(&domain.User{
						ID:       1,
						Name:     "test name",
						Password: "test password",
					}, nil)
			},
		},
	}

	for _, test := range tests {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockUserRepository(ctrl)
		timeout := 2 * time.Second
		usecase := NewUserUsecase(repo, timeout)
		test.mockFunc(repo)

		got, err := usecase.Find(context.Background(), test.userName, test.password)
		assert.NoError(t, err)
		assert.NotNil(t, got)
	}
}
