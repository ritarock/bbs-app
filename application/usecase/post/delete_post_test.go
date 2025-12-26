package post_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/application/usecase/post"
	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeletePostUsecase_Execute(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    dto.DeletePostInput
		mockFunc func(m *mock.MockPostRepository)
		hasError bool
	}{
		{
			name:  "pass: delete post",
			input: dto.DeletePostInput{ID: 1},
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().
					Delete(gomock.Any(), valueobject.NewPostID(1)).
					Return(nil)
			},
			hasError: false,
		},
		{
			name:  "failed: repository error",
			input: dto.DeletePostInput{ID: 1},
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().
					Delete(gomock.Any(), valueobject.NewPostID(1)).
					Return(errors.New("db error"))
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockRepo := mock.NewMockPostRepository(ctrl)
			test.mockFunc(mockRepo)
			uc := post.NewDeletePostUsecase(mockRepo)
			err := uc.Execute(context.Background(), test.input)

			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
