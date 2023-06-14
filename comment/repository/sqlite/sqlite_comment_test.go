package sqlite_test

import (
	"context"
	"ritarock/bbs-app/domain"
	"testing"
	"time"

	commentRepo "ritarock/bbs-app/comment/repository/sqlite"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	now := time.Now()
	post := &domain.Post{
		ID: 1,
	}
	comment := &domain.Comment{
		Content:     "content",
		CommentedAt: now,
		PostID:      1,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT INTO comment \\(content, commented_at, post_id\\) VALUES \\(\\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(comment.Content, sqlmock.AnyArg(), post.ID).
		WillReturnResult(sqlmock.NewResult(10, 1))

	cRepo := commentRepo.NewsqliteCommentRepository(db)

	err = cRepo.Create(context.Background(), comment)
	assert.NoError(t, err)
	assert.Equal(t, 10, comment.ID)
}

func TestGetAllByPost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockPost := domain.Post{
		ID: 1,
	}
	mockComments := []domain.Comment{
		{
			ID:          1,
			Content:     "content 1",
			CommentedAt: time.Now(),
			PostID:      mockPost.ID,
		},
		{
			ID:          2,
			Content:     "content 2",
			CommentedAt: time.Now(),
			PostID:      mockPost.ID,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "content", "commented_at", "post_id"}).
		AddRow(
			mockComments[0].ID,
			mockComments[0].Content,
			mockComments[0].CommentedAt.Format("2006-01-02 15:04:05.999999999-07:00"),
			mockComments[0].PostID,
		).
		AddRow(
			mockComments[1].ID,
			mockComments[1].Content,
			mockComments[1].CommentedAt.Format("2006-01-02 15:04:05.999999999-07:00"),
			mockComments[1].PostID,
		)

	query := "SELECT id, content, commented_at, post_id FROM comment WHERE post_id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)

	cRepo := commentRepo.NewsqliteCommentRepository(db)

	comments, err := cRepo.GetAllByPost(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, comments)
	assert.Equal(t, 2, len(comments))
}
