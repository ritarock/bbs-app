package repository

import (
	"context"
	"database/sql"

	"github.com/ritarock/bbs-app/domain"
)

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) domain.CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (c *commentRepository) fetch(ctx context.Context, query string, args ...any) ([]domain.Comment, error) {
	row, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	result := []domain.Comment{}
	for row.Next() {
		t := domain.Comment{}
		if err := row.Scan(
			&t.Id,
			&t.Content,
			&t.CommentedAt,
			&t.PostId,
		); err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (c *commentRepository) Create(ctx context.Context, comment *domain.Comment) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	query := "INSERT INTO comments (content, commented_at, post_id) VALUES (?, ?, ?)"
	res, err := tx.ExecContext(ctx, query, comment.Content, comment.CommentedAt, comment.PostId)
	if err != nil {
		return err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	comment.Id = int(lastId)
	return nil
}

func (c *commentRepository) GetByPostId(ctx context.Context, postId int) ([]domain.Comment, error) {
	query := "SELECT id, content, commented_at, post_id FROM comments WHERE post_id = ?"
	return c.fetch(ctx, query, postId)
}
