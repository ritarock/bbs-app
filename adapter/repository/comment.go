package repository

import (
	"context"
	"database/sql"

	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/port"
)

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) port.CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (c *commentRepository) fetch(ctx context.Context, query string, args ...any) ([]*domain.Comment, error) {
	row, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	result := []*domain.Comment{}
	for row.Next() {
		comment := &domain.Comment{}
		err := row.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.Content,
			&comment.CommentedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, comment)
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

	query := "INSERT INTO comments (post_id, content, commented_at) VALUES (?, ?, ?)"
	res, err := tx.ExecContext(ctx, query, comment.PostID, comment.Content, comment.CommentedAt)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	comment.ID = int(lastID)

	return nil

}

func (c *commentRepository) GetAll(ctx context.Context, postID int) ([]*domain.Comment, error) {
	query := "SELECT id, post_id, content, commented_at FROM comments WHERE post_id = ?"
	return c.fetch(ctx, query, postID)
}
