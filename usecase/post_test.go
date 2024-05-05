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

func Test_postUsecase_Create(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		post     *domain.Post
		mockFunc func(repo *mock.MockPostRepository)
	}{
		{
			name: "pass",
			post: &domain.Post{
				Title:    "test",
				Content:  "test",
				PostedAt: now,
			},
			mockFunc: func(repo *mock.MockPostRepository) {
				repo.EXPECT().Create(gomock.Any(),
					&domain.Post{
						Title:    "test",
						Content:  "test",
						PostedAt: now,
					}).
					Times(1).Return(nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			repo := mock.NewMockPostRepository(ctrl)
			timeout := 2 * time.Second
			usecase := NewPostUsecase(repo, timeout)
			test.mockFunc(repo)

			err := usecase.Create(context.Background(), test.post)
			assert.NoError(t, err)
		})
	}
}

func Test_postUsecase_GetAll(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		mockFunc func(repo *mock.MockPostRepository)
	}{
		{
			name: "pass",
			mockFunc: func(repo *mock.MockPostRepository) {
				repo.EXPECT().GetAll(gomock.Any()).Times(1).Return([]domain.Post{
					{
						ID:       1,
						Title:    "test1",
						Content:  "test1",
						PostedAt: now,
					},
					{
						ID:       2,
						Title:    "test2",
						Content:  "test2",
						PostedAt: now,
					},
				}, nil)
			},
		},
	}

	for _, test := range tests {
		ctrl := gomock.NewController(t)

		repo := mock.NewMockPostRepository(ctrl)
		timeout := 2 * time.Second
		usecase := NewPostUsecase(repo, timeout)
		test.mockFunc(repo)

		got, err := usecase.GetAll(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.Len(t, got, 2)
	}
}

func Test_postUsecase_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		mockFunc func(repo *mock.MockPostRepository)
	}{
		{
			name: "pass",
			mockFunc: func(repo *mock.MockPostRepository) {
				repo.EXPECT().GetByID(gomock.Any(), 1).Times(1).Return(
					&domain.Post{
						ID:       1,
						Title:    "test",
						Content:  "test",
						PostedAt: time.Now(),
					}, nil,
				)
			},
		},
	}

	for _, test := range tests {
		ctrl := gomock.NewController(t)

		repo := mock.NewMockPostRepository(ctrl)
		timeout := 2 * time.Second
		usecase := NewPostUsecase(repo, timeout)
		test.mockFunc(repo)

		got, err := usecase.GetByID(context.Background(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, got)
	}
}

func Test_postUsecase_Update(t *testing.T) {
	tests := []struct {
		name     string
		postID   int
		post     *domain.Post
		mockFunc func(repo *mock.MockPostRepository, post domain.Post)
		hasError bool
	}{
		{
			name:   "pass",
			postID: 1,
			post: &domain.Post{
				Title:    "test",
				Content:  "test",
				PostedAt: time.Now(),
			},
			mockFunc: func(repo *mock.MockPostRepository, post domain.Post) {
				repo.EXPECT().GetByID(gomock.Any(), 1).Times(1).Return(&post, nil)
				repo.EXPECT().Update(gomock.Any(), 1, &post).Times(1).Return(nil)
			},
			hasError: false,
		},
		{
			name:   "failed: not found",
			postID: 10,
			post: &domain.Post{
				Title:    "test",
				Content:  "test",
				PostedAt: time.Now(),
			},
			mockFunc: func(repo *mock.MockPostRepository, post domain.Post) {
				repo.EXPECT().GetByID(gomock.Any(), 10).Times(1).Return(nil, domain.ErrNotFound)
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			repo := mock.NewMockPostRepository(ctrl)
			timeout := 2 * time.Second
			usecase := NewPostUsecase(repo, timeout)
			test.mockFunc(repo, *test.post)

			err := usecase.Update(context.Background(), test.postID, test.post)
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_postUsecase_Delete(t *testing.T) {
	tests := []struct {
		name     string
		postID   int
		post     *domain.Post
		mockFunc func(repo *mock.MockPostRepository, post domain.Post)
		hasError bool
	}{
		{
			name:   "pass",
			postID: 1,
			post: &domain.Post{
				Title:    "test",
				Content:  "test",
				PostedAt: time.Now(),
			},
			mockFunc: func(repo *mock.MockPostRepository, post domain.Post) {
				repo.EXPECT().GetByID(gomock.Any(), 1).Times(1).Return(&post, nil)
				repo.EXPECT().Delete(gomock.Any(), 1).Times(1).Return(nil)
			},
			hasError: false,
		},
		{
			name:   "failed: not found",
			postID: 10,
			post: &domain.Post{
				Title:    "test",
				Content:  "test",
				PostedAt: time.Now(),
			},
			mockFunc: func(repo *mock.MockPostRepository, post domain.Post) {
				repo.EXPECT().GetByID(gomock.Any(), 10).Times(1).Return(nil, domain.ErrNotFound)
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			repo := mock.NewMockPostRepository(ctrl)
			timeout := 2 * time.Second
			usecase := NewPostUsecase(repo, timeout)
			test.mockFunc(repo, *test.post)

			err := usecase.Delete(context.Background(), test.postID)
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
