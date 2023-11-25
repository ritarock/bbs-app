package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ritarock/bbs-app/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
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
		t := domain.User{}
		if err := row.Scan(
			&t.Id,
			&t.Name,
			&t.Password,
			&t.Token,
		); err != nil {
			fmt.Println(err)
			return nil, err
		}
		result = append(result, t)
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

	query := "INSERT INTO users (name, password, token) VALUES (?, ?, '')"
	res, err := tx.ExecContext(ctx, query, user.Name, user.Password)
	if err != nil {
		return err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = int(lastId)
	return nil
}

func (u *userRepository) FindUser(ctx context.Context, name string, password string) (*domain.User, error) {
	query := `
		SELECT
			id, name, password, token
		FROM
			users
		WHERE
			name = ? AND password = ?
	`
	list, err := u.fetch(ctx, query, name, password)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, domain.ErrNotFound
	}
	if len(list) > 1 {
		return nil, errors.New("system error")
	}

	return &list[0], nil
}

func (u *userRepository) SetToken(ctx context.Context, userId int, token string) error {
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

	query := `
		UPDATE
			users
		SET
			token = ?
		WHERE
			id = ?
		`
	_, err = tx.ExecContext(ctx, query, token, userId)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) ExistToken(ctx context.Context, token string) bool {
	query := `
		SELECT
			id, name, password, token
		FROM
			users
		WHERE
			token = ?
	`

	list, err := u.fetch(ctx, query, token)
	if err != nil {
		return false
	}

	if len(list) == 0 {
		return false
	}

	return true
}
