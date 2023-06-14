package usecase_test

import (
	"context"
	"ritarock/bbs-app/comment/usecase"
	"ritarock/bbs-app/domain"
	"ritarock/bbs-app/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepository)
	mockPost := domain.Post{
		ID: 1,
	}
	mockComment := domain.Comment{
		ID:          1,
		Content:     "content 1",
		CommentedAt: time.Now(),
		PostID:      mockPost.ID,
	}

	t.Run("success", func(t *testing.T) {
		mockCommentRepo.
			On("Create", mock.Anything, mock.AnythingOfType("*domain.Comment")).
			Once().
			Return(nil)

		u := usecase.NewCommentUsecase(mockCommentRepo, time.Second*2)

		err := u.Create(context.Background(), &mockComment)

		assert.NoError(t, err)
		assert.Equal(t, mockComment.Content, "content 1")
		mockCommentRepo.AssertExpectations(t)
	})
}

func TestGetAllByPost(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepository)
	mockPost := domain.Post{
		ID: 1,
	}
	mockComments := []domain.Comment{
		{
			ID:          1,
			Content:     "content 1",
			CommentedAt: time.Now(),
			PostID:      mockPost.ID,
		},
		{
			ID:          2,
			Content:     "content 2",
			CommentedAt: time.Now(),
			PostID:      mockPost.ID,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockCommentRepo.
			On("GetAllByPost", mock.Anything, mock.AnythingOfType("int")).
			Once().
			Return(mockComments, nil)

		u := usecase.NewCommentUsecase(mockCommentRepo, time.Second*2)

		p, err := u.GetAllByPost(context.Background(), mockPost.ID)

		assert.NoError(t, err)
		assert.NotNil(t, p)
		assert.Equal(t, len(p), 2)
		mockCommentRepo.AssertExpectations(t)
	})
}
