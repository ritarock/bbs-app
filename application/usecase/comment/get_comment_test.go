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

func TestGetCommentUsecase_Execute(t *testing.T) {
	t.Parallel()
	commentedAt := time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name     string
		input    dto.GetCommentInput
		mockFunc func(m *mock.MockCommentRepository)
		want     *dto.GetCommentOutput
		hasError bool
	}{
		{
			name:  "pass: get comment",
			input: dto.GetCommentInput{ID: 1},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewCommentID(1)).
					Return(entity.ReconstructComment(
						valueobject.NewCommentID(1),
						valueobject.NewPostID(1),
						"comment body",
						commentedAt,
					), nil)
			},
			want: &dto.GetCommentOutput{
				ID:          1,
				PostID:      1,
				Body:        "comment body",
				CommentedAt: commentedAt,
			},
			hasError: false,
		},
		{
			name:  "failed: comment not found",
			input: dto.GetCommentInput{ID: 10},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewCommentID(10)).
					Return(nil, nil)
			},
			want:     nil,
			hasError: true,
		},
		{
			name:  "failed: repository get error",
			input: dto.GetCommentInput{ID: 1},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewCommentID(1)).
					Return(nil, errors.New("db error"))
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
			uc := comment.NewGetCommentUsecase(mockRepo)
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
