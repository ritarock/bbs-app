package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ritarock/bbs-app/domain"
	"github.com/stretchr/testify/assert"
)

func Test_userRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	user := &domain.User{
		Name:     "user",
		Password: "password1234",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.Name, user.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := NewUserRepository(db)
	err := repo.Create(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)
	assert.Equal(t, "user", user.Name)
	assert.Equal(t, "password1234", user.Password)
}

func Test_userRepository_FindUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	row := sqlmock.NewRows([]string{
		"id",
		"name",
		"password",
		"token",
	}).AddRow(
		1,
		"user",
		"password1234",
		"",
	)
	mock.ExpectQuery("SELECT").WithArgs("user", "password1234").WillReturnRows(row)
	repo := NewUserRepository(db)
	user, err := repo.FindUser(context.Background(), "user", "password1234")
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func Test_userRepository_SetToken(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	user := &domain.User{
		Id:       1,
		Name:     "user",
		Password: "password",
		Token:    "token",
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").
		WithArgs(user.Token, user.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := NewUserRepository(db)
	err := repo.SetToken(context.Background(), 1, "token")
	assert.NoError(t, err)
}

func Test_userRepository_ExistToken(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	row := sqlmock.NewRows([]string{
		"id",
		"name",
		"password",
		"token",
	}).AddRow(
		1,
		"user",
		"password1234",
		"token",
	)

	tests := []struct {
		name     string
		token    string
		hasError bool
	}{
		{
			name:     "pass",
			token:    "token",
			hasError: false,
		},
		{
			name:     "pass",
			token:    "tokennnn",
			hasError: true,
		},
	}
	mock.ExpectQuery("SELECT").WithArgs("token").WillReturnRows(row)
	repo := NewUserRepository(db)
	for _, test := range tests {
		result := repo.ExistToken(context.Background(), test.token)
		if test.hasError {
			assert.Equal(t, false, result)
		} else {
			assert.Equal(t, true, result)
		}
	}
}
