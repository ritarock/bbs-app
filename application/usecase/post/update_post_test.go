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

func TestUpdatePostUsecase_Execute(t *testing.T) {
	t.Parallel()
	postedAt := time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name     string
		input    dto.UpdatePostInput
		mockFunc func(m *mock.MockPostRepository)
		want     *dto.UpdatePostOutput
		hasError bool
	}{
		{
			name:  "pass: update post",
			input: dto.UpdatePostInput{ID: 1, Title: "new title", Content: "new content"},
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewPostID(1)).
					Return(entity.ReconstructPost(
						valueobject.NewPostID(1),
						"title",
						"content",
						postedAt,
					), nil)
				m.EXPECT().
					Update(gomock.Any(), gomock.Cond(func(p any) bool {
						post := p.(*entity.Post)
						return post.Title().String() == "new title" &&
							post.Content().String() == "new content"
					})).
					Return(nil)
			},
			want: &dto.UpdatePostOutput{
				ID:       1,
				Title:    "new title",
				Content:  "new content",
				PostedAt: postedAt,
			},
			hasError: false,
		},
		{
			name:  "failed: post not found",
			input: dto.UpdatePostInput{ID: 10, Title: "new title", Content: "new content"},
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewPostID(10)).
					Return(nil, nil)
			},
			want:     nil,
			hasError: true,
		},
		{
			name:  "failed: repository FindByID error",
			input: dto.UpdatePostInput{ID: 10, Title: "new title", Content: "new content"},
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewPostID(10)).
					Return(nil, errors.New("db error"))
			},
			want:     nil,
			hasError: true,
		},
		{
			name:  "failed: entity update error",
			input: dto.UpdatePostInput{ID: 1, Title: "", Content: "new content"},
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
			want:     nil,
			hasError: true,
		},
		{
			name:  "failed: repository Update error",
			input: dto.UpdatePostInput{ID: 1, Title: "new title", Content: "new content"},
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().
					FindByID(gomock.Any(), valueobject.NewPostID(1)).
					Return(entity.ReconstructPost(
						valueobject.NewPostID(1),
						"title",
						"content",
						postedAt,
					), nil)
				m.EXPECT().
					Update(gomock.Any(), gomock.Cond(func(p any) bool {
						post := p.(*entity.Post)
						return post.Title().String() == "new title" &&
							post.Content().String() == "new content"
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
			mockRepo := mock.NewMockPostRepository(ctrl)
			test.mockFunc(mockRepo)
			uc := post.NewUpdatePostUsecase(mockRepo)
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
