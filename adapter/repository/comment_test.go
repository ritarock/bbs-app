package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ritarock/bbs-app/domain"
	"github.com/stretchr/testify/assert"
)

func Test_commentRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	timeNow := time.Now()

	comment := &domain.Comment{
		Content:     "content",
		CommentedAt: timeNow,
		PostId:      1,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO comments").
		WithArgs(comment.Content, comment.CommentedAt, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := NewCommentRepository(db)
	err := repo.Create(context.Background(), comment)
	assert.NoError(t, err)
	assert.Equal(t, 1, comment.Id)
	assert.Equal(t, "content", comment.Content)
	assert.Equal(t, timeNow, comment.CommentedAt)
	assert.Equal(t, 1, comment.PostId)

}

func Test_commentRepository_GetByPostId(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	timeNow := time.Now()

	mockComments := []domain.Comment{
		{
			Id:          1,
			Content:     "content1",
			CommentedAt: timeNow,
			PostId:      1,
		},
		{
			Id:          2,
			Content:     "content1",
			CommentedAt: timeNow,
			PostId:      1,
		},
	}

	rows := sqlmock.NewRows([]string{
		"id",
		"content",
		"commented_at",
		"post_id",
	}).AddRow(
		mockComments[0].Id,
		mockComments[0].Content,
		mockComments[0].CommentedAt,
		mockComments[0].PostId,
	).AddRow(
		mockComments[1].Id,
		mockComments[1].Content,
		mockComments[1].CommentedAt,
		mockComments[1].PostId,
	)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	repo := NewCommentRepository(db)
	comments, err := repo.GetByPostId(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, comments)
}
