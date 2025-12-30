package comment_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/application/usecase/comment"
	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteCommentUsecase_Execute(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    dto.DeleteCommentInput
		mockFunc func(m *mock.MockCommentRepository)
		hasError bool
	}{
		{
			name:  "pass: delete comment",
			input: dto.DeleteCommentInput{ID: 1},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().
					Delete(gomock.Any(), valueobject.NewCommentID(1)).
					Return(nil)
			},
			hasError: false,
		},
		{
			name:  "failed: repository error",
			input: dto.DeleteCommentInput{ID: 1},
			mockFunc: func(m *mock.MockCommentRepository) {
				m.EXPECT().
					Delete(gomock.Any(), valueobject.NewCommentID(1)).
					Return(errors.New("db error"))
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockRepo := mock.NewMockCommentRepository(ctrl)
			test.mockFunc(mockRepo)
			uc := comment.NewDeleteCommentUsecase(mockRepo)
			err := uc.Execute(context.Background(), test.input)

			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
