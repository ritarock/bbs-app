package comment_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/application/usecase/comment"
	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestListCommentsUsecase_Execute(t *testing.T) {
	t.Parallel()
	commentedAt := time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name     string
		input    dto.ListCommentsInput
		mockFunc func(m *mock.MockCommentRepository)
		want     *dto.ListCommentsOutput
		hasError bool
	}{
		{
			name:  "pass: list comments",
			input: dto.ListCommentsInput{PostID: 1},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().FindByPostID(gomock.Any(), valueobject.NewPostID(1)).
					Return([]*entity.Comment{
						entity.ReconstructComment(
							valueobject.NewCommentID(1),
							valueobject.NewPostID(1),
							"comment1",
							commentedAt,
						),
						entity.ReconstructComment(
							valueobject.NewCommentID(2),
							valueobject.NewPostID(1),
							"comment2",
							commentedAt,
						),
					}, nil)

			},
			want: &dto.ListCommentsOutput{
				Comments: []dto.CommentItem{
					{ID: 1, PostID: 1, Body: "comment1", CommentedAt: commentedAt},
					{ID: 2, PostID: 1, Body: "comment2", CommentedAt: commentedAt},
				},
			},
			hasError: false,
		},
		{
			name:  "pass: empty list",
			input: dto.ListCommentsInput{PostID: 1},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().FindByPostID(gomock.Any(), valueobject.NewPostID(1)).
					Return([]*entity.Comment{}, nil)

			},
			want:     &dto.ListCommentsOutput{Comments: []dto.CommentItem{}},
			hasError: false,
		},
		{
			name:  "failed: repository list error",
			input: dto.ListCommentsInput{PostID: 1},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().FindByPostID(gomock.Any(), valueobject.NewPostID(1)).
					Return([]*entity.Comment{}, errors.New("db error"))

			},
			want:     nil,
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockRepo := mock.NewMockCommentRepository(ctrl)
			test.mockFunc(mockRepo)
			uc := comment.NewListCommentsUsecase(mockRepo)
			got, err := uc.Execute(context.Background(), test.input)

			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
