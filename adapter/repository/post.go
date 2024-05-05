package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/port"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) port.PostRepository {
	return &postRepository{
		db: db,
	}
}

func (p *postRepository) fetch(ctx context.Context, query string, args ...any) ([]domain.Post, error) {
	row, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	result := []domain.Post{}
	for row.Next() {
		post := domain.Post{}
		err := row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.PostedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, post)
	}

	return result, nil
}

func (p *postRepository) Create(ctx context.Context, post *domain.Post) error {
	tx, err := p.db.Begin()
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

	query := "INSERT INTO posts (title, content, posted_at) VALUES (?, ?, ?)"
	res, err := tx.ExecContext(ctx, query, post.Title, post.Content, post.PostedAt)
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	post.ID = int(lastID)
	return nil
}

func (p *postRepository) GetAll(ctx context.Context) ([]domain.Post, error) {
	query := "SELECT id, title, content, posted_at FROM posts"
	return p.fetch(ctx, query)
}

func (p *postRepository) GetByID(ctx context.Context, id int) (*domain.Post, error) {
	query := "SELECT id, title, content, posted_at FROM posts WHERE id = ?"
	result, err := p.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, domain.ErrNotFound
	}

	return &result[0], nil
}

func (p *postRepository) Update(ctx context.Context, postID int, post *domain.Post) error {
	tx, err := p.db.Begin()
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

	query := "UPDATE posts SET title = ?, content = ?, posted_at = ? WHERE id = ?"
	res, err := tx.ExecContext(ctx, query, post.Title, post.Content, post.PostedAt, postID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return fmt.Errorf("weird behavior total affected: %d", affected)
	}

	return nil
}

func (p *postRepository) Delete(ctx context.Context, postID int) error {
	tx, err := p.db.Begin()
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

	query := "DELETE FROM posts WHERE id = ?"
	res, err := tx.ExecContext(ctx, query, postID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return fmt.Errorf("weird behavior total affected: %d", affected)
	}

	return nil
}
