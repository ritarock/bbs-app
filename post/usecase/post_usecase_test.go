package usecase_test

import (
	"context"
	"errors"
	"ritarock/bbs-app/domain"
	"ritarock/bbs-app/domain/mocks"
	"ritarock/bbs-app/post/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockPostRepo := new(mocks.PostRepository)
	mockPost := domain.Post{
		ID:       1,
		Title:    "title 1",
		Content:  "content 1",
		PostedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockPostRepo.
			On("Create", mock.Anything, mock.AnythingOfType("*domain.Post")).
			Once().
			Return(nil)

		u := usecase.NewPostUsecase(mockPostRepo, time.Second*2)

		err := u.Create(context.Background(), &mockPost)

		assert.NoError(t, err)
		assert.Equal(t, mockPost.Title, "title 1")
		mockPostRepo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	mockPostRepo := new(mocks.PostRepository)
	mockPost := domain.Post{
		ID:       1,
		Title:    "title 1",
		Content:  "content 1",
		PostedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockPostRepo.
			On("GetById", mock.Anything, mock.AnythingOfType("int")).
			Once().
			Return(mockPost, nil)

		u := usecase.NewPostUsecase(mockPostRepo, time.Second*2)

		p, err := u.GetById(context.Background(), mockPost.ID)

		assert.NoError(t, err)
		assert.NotNil(t, p)
		mockPostRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockPostRepo.
			On("GetById", mock.Anything, mock.AnythingOfType("int")).
			Once().
			Return(domain.Post{}, errors.New("Unexpected"))

		u := usecase.NewPostUsecase(mockPostRepo, time.Second*2)

		p, err := u.GetById(context.Background(), mockPost.ID)

		assert.Error(t, err)
		assert.Equal(t, domain.Post{}, p)
		mockPostRepo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	mockPostRepo := new(mocks.PostRepository)
	mockPosts := []domain.Post{
		{
			ID:       1,
			Title:    "title 1",
			Content:  "content 1",
			PostedAt: time.Now(),
		},
		{
			ID:       2,
			Title:    "title 2",
			Content:  "content 2",
			PostedAt: time.Now(),
		},
	}

	t.Run("success", func(t *testing.T) {
		mockPostRepo.
			On("GetAll", mock.Anything).
			Once().
			Return(mockPosts, nil)

		u := usecase.NewPostUsecase(mockPostRepo, time.Second*2)

		p, err := u.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, p)
		assert.Equal(t, len(p), 2)
		mockPostRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockPostRepo := new(mocks.PostRepository)
	mockPost := domain.Post{
		ID:       1,
		Title:    "title 1",
		Content:  "content 1",
		PostedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockPostRepo.
			On("Update", mock.Anything, &mockPost).
			Once().
			Return(nil)

		u := usecase.NewPostUsecase(mockPostRepo, time.Second*2)

		err := u.Update(context.Background(), &mockPost)
		assert.NoError(t, err)
		mockPostRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockPostRepo := new(mocks.PostRepository)
	mockPost := domain.Post{
		ID:       1,
		Title:    "title 1",
		Content:  "content 1",
		PostedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockPostRepo.
			On("GetById", mock.Anything, mock.AnythingOfType("int")).
			Once().
			Return(mockPost, nil)

		mockPostRepo.
			On("Delete", mock.Anything, mock.AnythingOfType("int")).
			Once().
			Return(nil)

		u := usecase.NewPostUsecase(mockPostRepo, time.Second*2)

		err := u.Delete(context.Background(), mockPost.ID)
		assert.NoError(t, err)
		mockPostRepo.AssertExpectations(t)
	})

	t.Run("failed-post-is-not-exist", func(t *testing.T) {
		mockPostRepo.
			On("GetById", mock.Anything, mock.AnythingOfType("int")).
			Once().
			Return(domain.Post{}, nil)

		u := usecase.NewPostUsecase(mockPostRepo, time.Second*2)

		err := u.Delete(context.Background(), mockPost.ID)
		assert.Error(t, err)
		mockPostRepo.AssertExpectations(t)
	})

	t.Run("failed-happens-in-db", func(t *testing.T) {
		mockPostRepo.
			On("GetById", mock.Anything, mock.AnythingOfType("int")).
			Once().
			Return(domain.Post{}, errors.New("Unexpected Error"))

		u := usecase.NewPostUsecase(mockPostRepo, time.Second*2)

		err := u.Delete(context.Background(), mockPost.ID)
		assert.Error(t, err)
		mockPostRepo.AssertExpectations(t)
	})
}
