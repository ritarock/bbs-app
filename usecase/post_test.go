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
				Title:    "test title",
				Content:  "test content",
				PostedAt: now,
			},
			mockFunc: func(repo *mock.MockPostRepository) {
				repo.EXPECT().Create(gomock.Any(), &domain.Post{
					Title:    "test title",
					Content:  "test content",
					PostedAt: now,
				}).Return(nil)
			},
		},
	}

	for _, test := range tests {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockPostRepository(ctrl)
		timeout := 2 * time.Second
		usecase := NewPostUsecase(repo, timeout)
		test.mockFunc(repo)

		err := usecase.Create(context.Background(), test.post)
		assert.NoError(t, err)
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
				repo.EXPECT().GetAll(gomock.Any()).Return([]*domain.Post{
					{
						ID:       1,
						Title:    "test title",
						Content:  "test content",
						PostedAt: now,
					},
					{
						ID:       2,
						Title:    "test title2",
						Content:  "test content2",
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
		id       int
		mockFunc func(repo *mock.MockPostRepository)
	}{
		{
			name: "pass",
			id:   1,
			mockFunc: func(repo *mock.MockPostRepository) {
				repo.EXPECT().GetByID(gomock.Any(), 1).Return(&domain.Post{
					ID:       1,
					Title:    "test title",
					Content:  "test content",
					PostedAt: time.Now(),
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

		got, err := usecase.GetByID(context.Background(), test.id)
		assert.NoError(t, err)
		assert.NotNil(t, got)
	}
}

func Test_postUsecase_Update(t *testing.T) {
	tests := []struct {
		name     string
		post     *domain.Post
		id       int
		mockFunc func(repo *mock.MockPostRepository, post *domain.Post, id int)
		hasError bool
	}{
		{
			name: "pass",
			post: &domain.Post{
				Title:    "test title",
				Content:  "test content",
				PostedAt: time.Now(),
			},
			id: 1,
			mockFunc: func(repo *mock.MockPostRepository, post *domain.Post, id int) {
				repo.EXPECT().GetByID(gomock.Any(), id).Return(post, nil)
				repo.EXPECT().Update(gomock.Any(), post, id).Return(nil)
			},
			hasError: false,
		},
		{
			name: "failed: not found",
			post: &domain.Post{
				Title:    "test title",
				Content:  "test content",
				PostedAt: time.Now(),
			},
			id: 10,
			mockFunc: func(repo *mock.MockPostRepository, post *domain.Post, id int) {
				repo.EXPECT().GetByID(gomock.Any(), id).Return(nil, domain.ErrNotFound)
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
			test.mockFunc(repo, test.post, test.id)

			err := usecase.Update(context.Background(), test.post, test.id)
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
		post     *domain.Post
		id       int
		mockFunc func(repo *mock.MockPostRepository, post *domain.Post, id int)
		hasError bool
	}{
		{
			name: "pass",
			post: &domain.Post{
				Title:    "test title",
				Content:  "test content",
				PostedAt: time.Now(),
			},
			id: 1,
			mockFunc: func(repo *mock.MockPostRepository, post *domain.Post, id int) {
				repo.EXPECT().GetByID(gomock.Any(), id).Return(post, nil)
				repo.EXPECT().Delete(gomock.Any(), id).Return(nil)
			},
			hasError: false,
		},
		{
			name: "failed: not found",
			post: &domain.Post{
				Title:    "test title",
				Content:  "test content",
				PostedAt: time.Now(),
			},
			id: 10,
			mockFunc: func(repo *mock.MockPostRepository, post *domain.Post, id int) {
				repo.EXPECT().GetByID(gomock.Any(), id).Return(nil, domain.ErrNotFound)
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
			test.mockFunc(repo, test.post, test.id)

			err := usecase.Delete(context.Background(), test.id)
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
