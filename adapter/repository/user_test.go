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
				Name:     "test name",
				Password: "test password",
			},
			mockSql: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(
					"INSERT INTO users (name, password) VALUES (?, ?)")).
					WithArgs("test name", "test password").
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

func Test_userRepository_GetByNameAndPasswd(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		password string
		mockSql  func(mock sqlmock.Sqlmock)
	}{
		{
			name:     "pass",
			userName: "test name",
			password: "test password",
			mockSql: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "password"}).
					AddRow(1, "test name", "test password")
				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT id, name, password FROM users WHERE name = ? AND password = ?")).
					WithArgs("test name", "test password").
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
			got, err := repo.GetByNameAndPasswd(context.Background(), test.userName, test.password)
			assert.NoError(t, err)
			assert.NotNil(t, got)
		})
	}
}
