package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ritarock/bbs-app/domain"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) domain.PostRepository {
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
		t := domain.Post{}
		if err := row.Scan(
			&t.Id,
			&t.Title,
			&t.Content,
			&t.PostedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, t)
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

	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	post.Id = int(lastId)
	return nil
}

func (p *postRepository) GetAll(ctx context.Context) ([]domain.Post, error) {
	query := "SELECT id, title, content, posted_at FROM posts"
	return p.fetch(ctx, query)
}

func (p *postRepository) GetById(ctx context.Context, id int) (*domain.Post, error) {
	query := "SELECT id, title, content, posted_at FROM posts WHERE id = ?"
	list, err := p.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return &list[0], nil
	} else {
		return nil, domain.ErrNotFound
	}
}

func (p *postRepository) Update(ctx context.Context, post *domain.Post, id int) error {
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

	query := "UPDATE post SET title = ?, content = ?, WHERE id = ?"

	res, err := tx.ExecContext(ctx, query, post.Title, post.Content, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return fmt.Errorf("weird Behavior Total Affected: %d", affected)
	}

	return nil
}

func (p *postRepository) Delete(ctx context.Context, id int) error {
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

	query := "DELETE FROM post WHERE id = ?"

	res, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return fmt.Errorf("weird Behavior Total Affected: %d", affected)
	}

	return nil
}
