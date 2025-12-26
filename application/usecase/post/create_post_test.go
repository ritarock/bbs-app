package post_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/application/usecase/post"
	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreatePostUsecase_Execute(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    dto.CreatePostInput
		mockFunc func(m *mock.MockPostRepository)
		want     *dto.CreatePostOutput
		hasError bool
	}{
		{
			name:  "pass: create post",
			input: dto.CreatePostInput{Title: "title", Content: "content"},
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().
					Save(gomock.Any(), gomock.Cond(func(p any) bool {
						post := p.(*entity.Post)
						return post.Title().String() == "title" &&
							post.Content().String() == "content"
					})).
					Return(valueobject.NewPostID(1), nil)
			},
			want: &dto.CreatePostOutput{
				ID:      1,
				Title:   "title",
				Content: "content",
			},
			hasError: false,
		},
		{
			name:     "failed: new post error",
			input:    dto.CreatePostInput{Title: "", Content: ""},
			mockFunc: func(m *mock.MockPostRepository) {},
			want:     nil,
			hasError: true,
		},
		{
			name:  "failed: repository save error",
			input: dto.CreatePostInput{Title: "title", Content: "content"},
			mockFunc: func(m *mock.MockPostRepository) {
				m.EXPECT().Save(gomock.Any(), gomock.Any()).
					Return(valueobject.PostID{}, errors.New("db error"))
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
			uc := post.NewCreatePostUsecase(mockRepo)
			got, err := uc.Execute(context.Background(), test.input)

			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, test.want.ID, got.ID)
				assert.Equal(t, test.want.Title, got.Title)
				assert.Equal(t, test.want.Content, got.Content)
				assert.NotZero(t, got.CreatedAt)
			}
		})
	}
}
