package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/port"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) port.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) fetch(ctx context.Context, query string, args ...any) ([]domain.User, error) {
	row, err := u.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	result := []domain.User{}
	for row.Next() {
		user := domain.User{}
		err := row.Scan(
			&user.ID,
			&user.Name,
			&user.Password,
			&user.Token,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}

	return result, nil
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) error {
	tx, err := u.db.Begin()
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

	query := "INSERT INTO users (name, password) VALUES (?, ?)"
	res, err := tx.ExecContext(ctx, query, user.Name, user.Password)
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(lastID)
	return nil
}

func (u *userRepository) GetByNameAndPassword(ctx context.Context, name string, password string) (*domain.User, error) {
	query := `
		SELECT id, name, password, token
		FROM users
		WHERE name = ? AND password = ?
	`
	result, err := u.fetch(ctx, query, name, password)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, domain.ErrNotFound
	}

	return &result[0], nil
}

func (u *userRepository) UpdateToken(ctx context.Context, userID int, token string) error {
	tx, err := u.db.Begin()
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

	query := "UPDATE users SET token = ? WHERE id = ?"
	res, err := tx.ExecContext(ctx, query, token, userID)
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

func (u *userRepository) ExistToken(ctx context.Context, token string) (bool, error) {
	query := `
		SELECT id, name, password, token
		FROM users
		WHERE token = ?
	`
	result, err := u.fetch(ctx, query, token)
	if err != nil {
		return false, err
	}

	if len(result) == 0 {
		return false, err
	}

	return true, nil
}
