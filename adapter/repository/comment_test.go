package repository

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ritarock/bbs-app/domain"
	"github.com/stretchr/testify/assert"
)

func Test_commentRepository_Create(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name    string
		comment *domain.Comment
		mockSql func(mock sqlmock.Sqlmock)
	}{
		{
			name: "pass",
			comment: &domain.Comment{
				PostID:      1,
				Content:     "test",
				CommentedAt: now,
			},
			mockSql: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(
					"INSERT INTO comments (post_id, content, commented_at) VALUES (?, ?, ?)")).
					WithArgs(1, "test", now).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			test.mockSql(mock)
			repo := NewCommentRepository(db)
			err := repo.Create(context.Background(), test.comment)
			assert.NoError(t, err)
		})
	}
}

func Test_commentRepository_GetAll(t *testing.T) {
	tests := []struct {
		name    string
		mockSql func(mock sqlmock.Sqlmock)
	}{
		{
			name: "pass",
			mockSql: func(mock sqlmock.Sqlmock) {
				now := time.Now()
				rows := sqlmock.NewRows([]string{"id", "post_id", "content", "commented_at"}).
					AddRow(1, 1, "test1", now).
					AddRow(2, 1, "test2", now)
				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT id, posted_at, content, commented_at FROM comments")).
					WillReturnRows(rows)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			test.mockSql(mock)
			repo := NewCommentRepository(db)
			got, err := repo.GetAll(context.Background())
			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.Len(t, got, 2)
		})
	}
}

func Test_commentRepository_GetByPostID(t *testing.T) {
	tests := []struct {
		name    string
		mockSql func(mock sqlmock.Sqlmock)
	}{
		{
			name: "pass",
			mockSql: func(mock sqlmock.Sqlmock) {
				now := time.Now()
				rows := sqlmock.NewRows([]string{"id", "post_id", "content", "commented_at"}).
					AddRow(1, 1, "test1", now).
					AddRow(2, 1, "test2", now)
				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT id, post_id, content, commented_at FROM comments WHERE post_id = ?")).
					WillReturnRows(rows)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			test.mockSql(mock)
			repo := NewCommentRepository(db)
			got, err := repo.GetByPostID(context.Background(), 1)
			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.Len(t, got, 2)
		})
	}
}
