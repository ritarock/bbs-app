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

func TestUpdateCommentUsecase_Execute(t *testing.T) {
	t.Parallel()
	commentedAt := time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name     string
		input    dto.UpdateCommentInput
		mockFunc func(m *mock.MockCommentRepository)
		want     *dto.UpdateCommentOutput
		hasError bool
	}{
		{
			name:  "pass: update comment",
			input: dto.UpdateCommentInput{ID: 1, Body: "new body"},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewCommentID(1)).
					Return(entity.ReconstructComment(
						valueobject.NewCommentID(1),
						valueobject.NewPostID(1),
						"old body",
						commentedAt,
					), nil)
				m.EXPECT().
					Update(gomock.Any(), gomock.Cond(func(c any) bool {
						comment := c.(*entity.Comment)
						return comment.Body().String() == "new body"
					})).
					Return(nil)
			},
			want: &dto.UpdateCommentOutput{
				ID:          1,
				PostID:      1,
				Body:        "new body",
				CommentedAt: commentedAt,
			},
			hasError: false,
		},
		{
			name:  "failed: comment not found",
			input: dto.UpdateCommentInput{ID: 10, Body: "new body"},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewCommentID(10)).
					Return(nil, nil)
			},
			want:     nil,
			hasError: true,
		},
		{
			name:  "failed: repository FindByID error",
			input: dto.UpdateCommentInput{ID: 10, Body: "new body"},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewCommentID(10)).
					Return(nil, errors.New("db error"))
			},
			want:     nil,
			hasError: true,
		},
		{
			name:  "failed: entity update error",
			input: dto.UpdateCommentInput{ID: 1, Body: ""},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewCommentID(1)).
					Return(entity.ReconstructComment(
						valueobject.NewCommentID(1),
						valueobject.NewPostID(1),
						"old body",
						commentedAt,
					), nil)

			},
			want:     nil,
			hasError: true,
		},
		{
			name:  "failed: repository Update error",
			input: dto.UpdateCommentInput{ID: 1, Body: "new body"},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewCommentID(1)).
					Return(entity.ReconstructComment(
						valueobject.NewCommentID(1),
						valueobject.NewPostID(1),
						"old body",
						commentedAt,
					), nil)
				m.EXPECT().
					Update(gomock.Any(), gomock.Cond(func(c any) bool {
						comment := c.(*entity.Comment)
						return comment.Body().String() == "new body"
					})).
					Return(errors.New("db error"))
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
			uc := comment.NewUpdateCommentUsecase(mockRepo)
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
