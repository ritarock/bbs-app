package post_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/application/usecase/post"
	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetPostUsecase_Execute(t *testing.T) {
	t.Parallel()
	postedAt := time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name     string
		input    dto.GetPostInput
		mockFunc func(m *mock.MockPostRepository)
		want     *dto.GetPostOutput
		hasError bool
	}{
		{
			name:  "pass: get post",
			input: dto.GetPostInput{ID: 1},
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewPostID(1)).
					Return(entity.ReconstructPost(
						valueobject.NewPostID(1),
						"title",
						"content",
						postedAt,
					), nil)
			},
			want: &dto.GetPostOutput{
				ID:       1,
				Title:    "title",
				Content:  "content",
				PostedAt: postedAt,
			},
			hasError: false,
		},
		{
			name:  "failed: post not found",
			input: dto.GetPostInput{ID: 10},
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewPostID(10)).
					Return(nil, nil)
			},
			want:     nil,
			hasError: true,
		},
		{
			name:  "failed: repository get error",
			input: dto.GetPostInput{ID: 1},
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewPostID(1)).
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
			mockRepo := mock.NewMockPostRepository(ctrl)
			test.mockFunc(mockRepo)
			uc := post.NewGetPostUsecase(mockRepo)
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
