package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/testing/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_commentUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockCommentRepository(ctrl)
	timeout := 2 * time.Second
	usecase := NewCommentUsecase(repo, timeout)

	timeNow := time.Now()
	comment := domain.Comment{
		Id:          1,
		Content:     "content",
		CommentedAt: timeNow,
		PostId:      1,
	}
	repo.EXPECT().Create(gomock.Any(), &comment).Times(1).Return(nil)

	err := usecase.Create(context.Background(), &comment)
	assert.NoError(t, err)
}

func Test_commentUsecase_GetByPostId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockCommentRepository(ctrl)
	timeout := 2 * time.Second
	usecase := NewCommentUsecase(repo, timeout)

	timeNow := time.Now()
	comments := []domain.Comment{
		{
			Id:          1,
			Content:     "content1",
			CommentedAt: timeNow,
			PostId:      1,
		},
		{
			Id:          2,
			Content:     "content2",
			CommentedAt: timeNow,
			PostId:      1,
		},
	}
	repo.EXPECT().GetByPostId(gomock.Any(), 1).Times(1).Return(comments, nil)

	getComments, err := usecase.GetByPostId(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(getComments))
	assert.Equal(t, comments[0].Content, getComments[0].Content)
}
