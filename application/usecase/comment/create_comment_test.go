package comment_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/application/usecase/comment"
	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateCommentUsecase_Execute(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    dto.CreateCommentInput
		mockFunc func(m *mock.MockCommentRepository)
		want     *dto.CreateCommentOutput
		hasError bool
	}{
		{
			name:  "pass: create comment",
			input: dto.CreateCommentInput{PostID: 1, Body: "comment body"},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().
					Save(gomock.Any(), gomock.Cond(func(c any) bool {
						comment := c.(*entity.Comment)
						return comment.PostID().Int() == 1 &&
							comment.Body().String() == "comment body"
					})).
					Return(valueobject.NewCommentID(1), nil)
			},
			want: &dto.CreateCommentOutput{
				ID:     1,
				PostID: 1,
				Body:   "comment body",
			},
			hasError: false,
		},
		{
			name:     "failed: new comment error",
			input:    dto.CreateCommentInput{PostID: 1, Body: ""},
			mockFunc: func(m *mock.MockCommentRepository) {},
			want:     nil,
			hasError: true,
		},
		{
			name:  "failed: repository save error",
			input: dto.CreateCommentInput{PostID: 1, Body: "comment body"},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().Save(gomock.Any(), gomock.Any()).
					Return(valueobject.CommentID{}, errors.New("db error"))
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
			uc := comment.NewCreateCommentUsecase(mockRepo)
			got, err := uc.Execute(context.Background(), test.input)

			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, test.want.ID, got.ID)
				assert.Equal(t, test.want.PostID, got.PostID)
				assert.Equal(t, test.want.Body, got.Body)
				assert.NotZero(t, got.CommentedAt)
			}
		})
	}
}
