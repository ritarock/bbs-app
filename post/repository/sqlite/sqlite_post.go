package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"ritarock/bbs-app/domain"
	"time"
)

type sqlitePostRepository struct {
	Conn *sql.DB
}

func NewSqlitePostRepository(conn *sql.DB) domain.PostRepository {
	return &sqlitePostRepository{
		Conn: conn,
	}
}

func (s *sqlitePostRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.Post, error) {
	rows, err := s.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]domain.Post, 0)
	for rows.Next() {
		t := domain.Post{}
		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&t.PostedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (s sqlitePostRepository) Create(ctx context.Context, post *domain.Post) error {
	query := "INSERT post (title, content, posted_at) VALUES (?, ?, ?)"
	stmt, err := s.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, post.Title, post.Content, time.Now())
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

func (s sqlitePostRepository) GetById(ctx context.Context, id int) (domain.Post, error) {
	query := `SELECT id, title, content, posted_at FROM post WHERE id = ?`

	list, err := s.fetch(ctx, query, id)
	if err != nil {
		return domain.Post{}, err
	}
	var res domain.Post
	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return res, nil
}

func (s sqlitePostRepository) GetAll(ctx context.Context) ([]domain.Post, error) {
	query := `SELECT id, title, content, posted_at FROM post`

	list, err := s.fetch(ctx, query)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s sqlitePostRepository) Update(ctx context.Context, post *domain.Post) error {
	query := `UPDATE post SET title = ?, content = ? WHERE id = ?`

	stmt, err := s.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, post.Title, post.Content, post.ID)
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

func (s sqlitePostRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM post WHERE id = ?`

	stmt, err := s.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
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
