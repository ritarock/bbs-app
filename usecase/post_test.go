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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockPostRepository(ctrl)
	timeout := 2 * time.Second
	usecase := NewPostUsecase(repo, timeout)

	timeNow := time.Now()
	post := domain.Post{
		Id:       1,
		Title:    "title",
		Content:  "content",
		PostedAt: timeNow,
	}
	repo.EXPECT().Create(gomock.Any(), &post).Times(1).Return(nil)

	err := usecase.Create(context.Background(), &post)
	assert.NoError(t, err)
}

func Test_postUsecase_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockPostRepository(ctrl)
	timeout := 2 * time.Second
	usecase := NewPostUsecase(repo, timeout)

	timeNow := time.Now()

	posts := []domain.Post{
		{
			Id:       1,
			Title:    "title1",
			Content:  "content1",
			PostedAt: timeNow,
		},
		{
			Id:       2,
			Title:    "title2",
			Content:  "content2",
			PostedAt: timeNow,
		},
	}
	repo.EXPECT().GetAll(gomock.Any()).Times(1).Return(posts, nil)

	getPosts, err := usecase.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 2, len(getPosts))
	assert.Equal(t, posts[0].Title, getPosts[0].Title)
}

func Test_postUsecase_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockPostRepository(ctrl)
	timeout := 2 * time.Second
	usecase := NewPostUsecase(repo, timeout)

	timeNow := time.Now()

	post := &domain.Post{
		Id:       1,
		Title:    "title",
		Content:  "content",
		PostedAt: timeNow,
	}
	repo.EXPECT().GetById(gomock.Any(), 1).Times(1).Return(post, nil)

	getPost, err := usecase.GetById(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, post.Title, getPost.Title)
}

func Test_postUsecase_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockPostRepository(ctrl)
	timeout := 2 * time.Second
	usecase := NewPostUsecase(repo, timeout)

	timeNow := time.Now()

	tests := []struct {
		name     string
		post     domain.Post
		id       int
		mockFunc func(t *testing.T, repo *mock.MockPostRepository, post domain.Post)
		hasError bool
	}{
		{
			name: "pass",
			post: domain.Post{
				Id:       1,
				Title:    "title",
				Content:  "content",
				PostedAt: timeNow,
			},
			id: 1,
			mockFunc: func(t *testing.T, repo *mock.MockPostRepository, post domain.Post) {
				repo.EXPECT().GetById(gomock.Any(), 1).Times(1).Return(&post, nil)
				repo.EXPECT().Update(gomock.Any(), &post, 1).Times(1).Return(nil)
			},
			hasError: false,
		},
		{
			name: "failed",
			post: domain.Post{
				Id:       1,
				Title:    "title",
				Content:  "content",
				PostedAt: timeNow,
			},
			id: 10,
			mockFunc: func(t *testing.T, repo *mock.MockPostRepository, post domain.Post) {
				repo.EXPECT().GetById(gomock.Any(), 10).Times(1).
					Return(nil, domain.ErrNotFound)
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		test.mockFunc(t, repo, test.post)
		err := usecase.Update(context.Background(), &test.post, test.id)
		if test.hasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func Test_postUsecase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockPostRepository(ctrl)
	timeout := 2 * time.Second
	usecase := NewPostUsecase(repo, timeout)

	timeNow := time.Now()

	tests := []struct {
		name     string
		post     domain.Post
		id       int
		mockFunc func(t *testing.T, repo *mock.MockPostRepository, post domain.Post)
		hasError bool
	}{
		{
			name: "pass",
			post: domain.Post{
				Id:       1,
				Title:    "title",
				Content:  "content",
				PostedAt: timeNow,
			},
			id: 1,
			mockFunc: func(t *testing.T, repo *mock.MockPostRepository, post domain.Post) {
				repo.EXPECT().GetById(gomock.Any(), 1).Times(1).Return(&post, nil)
				repo.EXPECT().Delete(gomock.Any(), 1).Times(1).Return(nil)
			},
			hasError: false,
		},
		{
			name: "failed",
			post: domain.Post{
				Id:       1,
				Title:    "title",
				Content:  "content",
				PostedAt: timeNow,
			},
			id: 10,
			mockFunc: func(t *testing.T, repo *mock.MockPostRepository, post domain.Post) {
				repo.EXPECT().GetById(gomock.Any(), 10).Times(1).
					Return(nil, domain.ErrNotFound)
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		test.mockFunc(t, repo, test.post)
		err := usecase.Delete(context.Background(), test.id)
		if test.hasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
