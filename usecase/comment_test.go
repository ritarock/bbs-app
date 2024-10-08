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

func Test_commentUsecase_Create(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		comment  *domain.Comment
		mockFunc func(repo *mock.MockCommentRepository)
	}{
		{
			name: "pass",
			comment: &domain.Comment{
				PostID:      1,
				Content:     "test content",
				CommentedAt: now,
			},
			mockFunc: func(repo *mock.MockCommentRepository) {
				repo.EXPECT().Create(gomock.Any(), &domain.Comment{
					PostID:      1,
					Content:     "test content",
					CommentedAt: now,
				}).Return(nil)
			},
		},
	}

	for _, test := range tests {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockCommentRepository(ctrl)
		timeout := 2 * time.Second
		usecase := NewCommentUsecase(repo, timeout)
		test.mockFunc(repo)

		err := usecase.Create(context.Background(), test.comment)
		assert.NoError(t, err)
	}
}

func Test_commentUsecase_GetAll(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		mockFunc func(repo *mock.MockCommentRepository)
	}{
		{
			name: "pass",
			mockFunc: func(repo *mock.MockCommentRepository) {
				repo.EXPECT().GetAll(gomock.Any(), 1).Return([]*domain.Comment{
					{
						ID:          1,
						PostID:      1,
						Content:     "test content",
						CommentedAt: now,
					},
					{
						ID:          2,
						PostID:      1,
						Content:     "test content",
						CommentedAt: now,
					},
				}, nil)
			},
		},
	}

	for _, test := range tests {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockCommentRepository(ctrl)
		timeout := 2 * time.Second
		usecase := NewCommentUsecase(repo, timeout)
		test.mockFunc(repo)

		got, err := usecase.GetAll(context.Background(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.Len(t, got, 2)
	}
}
