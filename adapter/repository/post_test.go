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

func Test_postRepository_Create(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name    string
		post    *domain.Post
		mockSql func(mock sqlmock.Sqlmock)
	}{
		{
			name: "pass",
			post: &domain.Post{
				Title:    "test",
				Content:  "test",
				PostedAt: now,
			},
			mockSql: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(
					"INSERT INTO posts (title, content, posted_at) VALUES (?, ?, ?)")).
					WithArgs("test", "test", now).
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
			repo := NewPostRepository(db)
			err := repo.Create(context.Background(), test.post)
			assert.NoError(t, err)
		})
	}
}

func Test_postRepository_GetAll(t *testing.T) {
	tests := []struct {
		name    string
		mockSql func(mock sqlmock.Sqlmock)
	}{
		{
			name: "pass",
			mockSql: func(mock sqlmock.Sqlmock) {
				now := time.Now()
				rows := sqlmock.NewRows([]string{"id", "title", "content", "posted_at"}).
					AddRow(1, "test1", "test1", now).
					AddRow(2, "test2", "test2", now)
				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT id, title, content, posted_at FROM posts")).
					WillReturnRows(rows)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			test.mockSql(mock)
			repo := NewPostRepository(db)
			got, err := repo.GetAll(context.Background())
			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.Len(t, got, 2)
		})
	}
}

func Test_postRepository_GetByID(t *testing.T) {
	tests := []struct {
		name    string
		mockSql func(mock sqlmock.Sqlmock)
	}{
		{
			name: "pass",
			mockSql: func(mock sqlmock.Sqlmock) {
				now := time.Now()
				rows := sqlmock.NewRows([]string{"id", "title", "content", "posted_at"}).
					AddRow(1, "test1", "test1", now)
				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT id, title, content, posted_at FROM posts WHERE id = ?")).
					WithArgs(1).
					WillReturnRows(rows)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			test.mockSql(mock)
			repo := NewPostRepository(db)
			got, err := repo.GetByID(context.Background(), 1)
			assert.NoError(t, err)
			assert.NotNil(t, got)
		})
	}
}

func Test_postRepository_Update(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name    string
		post    *domain.Post
		mockSql func(mock sqlmock.Sqlmock)
	}{
		{
			name: "pass",
			post: &domain.Post{
				Title:    "test",
				Content:  "test",
				PostedAt: now,
			},
			mockSql: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(
					"UPDATE posts SET title = ?, content = ?, posted_at = ? WHERE id = ?")).
					WithArgs("test", "test", now, 1).
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
			repo := NewPostRepository(db)
			err := repo.Update(context.Background(), 1, test.post)
			assert.NoError(t, err)
		})
	}
}

func Test_postRepository_Delete(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name    string
		post    *domain.Post
		mockSql func(mock sqlmock.Sqlmock)
	}{
		{
			name: "pass",
			post: &domain.Post{
				Title:    "test",
				Content:  "test",
				PostedAt: now,
			},
			mockSql: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(
					"DELETE FROM posts WHERE id = ?")).
					WithArgs(1).
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
			repo := NewPostRepository(db)
			err := repo.Delete(context.Background(), 1)
			assert.NoError(t, err)
		})
	}
}
