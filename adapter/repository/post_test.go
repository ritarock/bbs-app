package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ritarock/bbs-app/domain"
	"github.com/stretchr/testify/assert"
)

func Test_postRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	timeNow := time.Now()

	post := &domain.Post{
		Title:    "title",
		Content:  "content",
		PostedAt: timeNow,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO posts").
		WithArgs(post.Title, post.Content, timeNow).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := NewPostRepository(db)
	err := repo.Create(context.Background(), post)
	assert.NoError(t, err)
	assert.Equal(t, 1, post.Id)
	assert.Equal(t, "title", post.Title)
	assert.Equal(t, "content", post.Content)
	assert.Equal(t, timeNow, post.PostedAt)
}

func Test_postRepository_GetAll(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	timeNow := time.Now()

	mockPosts := []domain.Post{
		{
			Id:       1,
			Title:    "title1",
			Content:  "content1",
			PostedAt: timeNow,
		},
		{
			Id:       2,
			Title:    "title2",
			Content:  "content2",
			PostedAt: timeNow,
		},
	}

	rows := sqlmock.NewRows([]string{
		"id",
		"title",
		"content",
		"posted_at",
	}).AddRow(
		mockPosts[0].Id,
		mockPosts[0].Title,
		mockPosts[0].Content,
		mockPosts[0].PostedAt,
	).AddRow(
		mockPosts[1].Id,
		mockPosts[1].Title,
		mockPosts[1].Content,
		mockPosts[1].PostedAt,
	)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	repo := NewPostRepository(db)
	posts, err := repo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, posts)
}

func Test_postRepository_GetById(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	timeNow := time.Now()

	row := sqlmock.NewRows([]string{
		"id",
		"title",
		"content",
		"posted_at",
	}).AddRow(
		1,
		"title",
		"content",
		timeNow,
	)
	mock.ExpectQuery("SELECT").WithArgs(1).WillReturnRows(row)
	repo := NewPostRepository(db)
	posts, err := repo.GetById(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, posts)
}

func Test_postRepository_Update(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	timeNow := time.Now()

	post := &domain.Post{
		Id:       10,
		Title:    "title",
		Content:  "content",
		PostedAt: timeNow,
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").
		WithArgs(post.Title, post.Content, post.Id).
		WillReturnResult(sqlmock.NewResult(10, 1))
	mock.ExpectCommit()

	repo := NewPostRepository(db)
	err := repo.Update(context.Background(), post, 10)
	assert.NoError(t, err)
}

func Test_postRepository_Delete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := NewPostRepository(db)
	err := repo.Delete(context.Background(), 1)
	assert.NoError(t, err)
}
