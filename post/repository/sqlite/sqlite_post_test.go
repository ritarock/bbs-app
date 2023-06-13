package sqlite_test

import (
	"context"
	"ritarock/bbs-app/domain"
	postRepo "ritarock/bbs-app/post/repository/sqlite"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	now := time.Now()
	post := &domain.Post{
		Title:    "title 1",
		Content:  "content 1",
		PostedAt: now,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT post \\(title, content, posted_at\\) VALUES \\(\\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(post.Title, post.Content, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(10, 1))

	pRepo := postRepo.NewSqlitePostRepository(db)

	err = pRepo.Create(context.Background(), post)
	assert.NoError(t, err)
	assert.Equal(t, 10, post.ID)
}

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "content", "posted_at"}).
		AddRow(1, "title 1", "content 1", time.Now())

	query := "SELECT id, title, content, posted_at FROM post WHERE id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)

	pRepo := postRepo.NewSqlitePostRepository(db)

	post, err := pRepo.GetById(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, post)
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockPosts := []domain.Post{
		{
			ID:       1,
			Title:    "title 1",
			Content:  "content 1",
			PostedAt: time.Now(),
		},
		{
			ID:       2,
			Title:    "title 2",
			Content:  "content 2",
			PostedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "content", "posted_at"}).
		AddRow(
			mockPosts[0].ID,
			mockPosts[0].Title,
			mockPosts[0].Content,
			mockPosts[0].PostedAt,
		).
		AddRow(
			mockPosts[1].ID,
			mockPosts[1].Title,
			mockPosts[1].Content,
			mockPosts[1].PostedAt,
		)

	query := "SELECT id, title, content, posted_at FROM post"

	mock.ExpectQuery(query).WillReturnRows(rows)

	pRepo := postRepo.NewSqlitePostRepository(db)

	post, err := pRepo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, post)
}

func TestUpdate(t *testing.T) {
	now := time.Now()
	post := &domain.Post{
		ID:       10,
		Title:    "title 1",
		Content:  "content 1",
		PostedAt: now,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "UPDATE post SET title = \\?, content = \\? WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(post.Title, post.Content, post.ID).
		WillReturnResult(sqlmock.NewResult(10, 1))

	pRepo := postRepo.NewSqlitePostRepository(db)

	err = pRepo.Update(context.Background(), post)
	assert.NoError(t, err)
	assert.NotNil(t, post)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM post WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	pRepo := postRepo.NewSqlitePostRepository(db)

	err = pRepo.Delete(context.Background(), 1)
	assert.NoError(t, err)
}
