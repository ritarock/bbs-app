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

func TestListPostUsecase_Execute(t *testing.T) {
	t.Parallel()
	postedAt := time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name     string
		mockFunc func(m *mock.MockPostRepository)
		want     *dto.ListPostOutput
		hasError bool
	}{
		{
			name: "pass: list posts",
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().FindAll(gomock.Any()).
					Return([]*entity.Post{
						entity.ReconstructPost(valueobject.NewPostID(1),
							"title1", "content1", postedAt,
						),
						entity.ReconstructPost(valueobject.NewPostID(2),
							"title2", "content2", postedAt,
						),
					}, nil)

			},
			want: &dto.ListPostOutput{
				Posts: []dto.PostItem{
					{ID: 1, Title: "title1", Content: "content1", PostedAt: postedAt},
					{ID: 2, Title: "title2", Content: "content2", PostedAt: postedAt},
				},
			},
			hasError: false,
		},
		{
			name: "pass: empty list",
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().FindAll(gomock.Any()).
					Return([]*entity.Post{}, nil)

			},
			want:     &dto.ListPostOutput{Posts: []dto.PostItem{}},
			hasError: false,
		},
		{
			name: "failed: repository list error",
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().FindAll(gomock.Any()).
					Return([]*entity.Post{}, errors.New("db error"))

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
			uc := post.NewListPostUsecase(mockRepo)
			got, err := uc.Execute(context.Background())

			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
