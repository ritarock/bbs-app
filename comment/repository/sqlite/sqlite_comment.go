package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"ritarock/bbs-app/domain"
	"time"
)

type sqliteCommentRepository struct {
	Conn *sql.DB
}

func NewsqliteCommentRepository(conn *sql.DB) domain.CommentRepository {
	return &sqliteCommentRepository{
		Conn: conn,
	}
}

func (s *sqliteCommentRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.Comment, error) {
	rows, err := s.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]domain.Comment, 0)
	for rows.Next() {
		t := domain.Comment{}
		var commentAtStr string
		err := rows.Scan(
			&t.ID,
			&t.Content,
			&commentAtStr,
			&t.PostID,
		)
		if err != nil {
			return nil, err
		}
		t.CommentedAt, err = time.Parse("2006-01-02 15:04:05.999999999-07:00", commentAtStr)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (s *sqliteCommentRepository) Create(ctx context.Context, comment *domain.Comment) error {
	query := "INSERT INTO comment (content, commented_at, post_id) VALUES (?, ?, ?)"
	stmt, err := s.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, comment.Content, time.Now(), comment.PostID)
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return err
	}
	comment.ID = int(lastID)
	return nil
}

func (s *sqliteCommentRepository) GetAllByPost(ctx context.Context, postId int) ([]domain.Comment, error) {
	query := `SELECT id, content, commented_at, post_id FROM comment WHERE post_id = ?`

	list, err := s.fetch(ctx, query, postId)
	if err != nil {
		return nil, err
	}
	return list, nil
}
