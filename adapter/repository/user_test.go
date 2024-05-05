package repository

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ritarock/bbs-app/domain"
	"github.com/stretchr/testify/assert"
)

func Test_userRepository_Create(t *testing.T) {
	tests := []struct {
		name    string
		user    *domain.User
		mockSql func(mock sqlmock.Sqlmock)
	}{
		{
			name: "pass",
			user: &domain.User{
				Name:     "test",
				Password: "test",
			},
			mockSql: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(
					"INSERT INTO users (name, password) VALUES (?, ?)")).
					WithArgs("test", "test").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			test.mockSql(mock)
			repo := NewUserRepository(db)
			err := repo.Create(context.Background(), test.user)
			assert.NoError(t, err)
		})
	}
}

func Test_userRepository_GetByNameAndPassword(t *testing.T) {
	tests := []struct {
		name        string
		argName     string
		argPassword string
		mockSql     func(mock sqlmock.Sqlmock)
	}{
		{
			name:        "pass",
			argName:     "test-name",
			argPassword: "test-password",
			mockSql: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "password", "token"}).
					AddRow(1, "test-name", "test-paasword", "test")
				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT id, name, password, token FROM users WHERE name = ? AND password = ?")).
					WithArgs("test-name", "test-password").
					WillReturnRows(rows)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			test.mockSql(mock)
			repo := NewUserRepository(db)
			got, err := repo.GetByNameAndPassword(context.Background(), test.argName, test.argPassword)
			assert.NoError(t, err)
			assert.NotNil(t, got)
		})
	}
}

func Test_userRepository_UpdateToken(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		mockSql func(mock sqlmock.Sqlmock)
	}{
		{
			name:  "pass",
			token: "test",
			mockSql: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(
					"UPDATE users SET token = ? WHERE id = ?")).
					WithArgs("test", 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			test.mockSql(mock)
			repo := NewUserRepository(db)
			err := repo.UpdateToken(context.Background(), 1, test.token)
			assert.NoError(t, err)
		})
	}
}

func Test_userRepository_ExistToken(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		mockSql func(mock sqlmock.Sqlmock)
	}{
		{
			name:  "pass",
			token: "test",
			mockSql: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "password", "token"}).
					AddRow(1, "test", "test", "test")
				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT id, name, password, token FROM users WHERE token = ?")).
					WithArgs("test").
					WillReturnRows(rows)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			test.mockSql(mock)
			repo := NewUserRepository(db)
			got, err := repo.ExistToken(context.Background(), test.token)
			assert.NoError(t, err)
			assert.NotNil(t, got)
		})
	}
}
