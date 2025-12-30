package database

import (
	"context"
	"errors"

	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/repository"
	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/ritarock/bbs-app/infra/database/query"
)

type UserRepository struct {
	queries *query.Queries
}

func NewUserRepository(db query.DBTX) repository.UserRepository {
	return &UserRepository{
		queries: query.New(db),
	}
}

func (r *UserRepository) Save(ctx context.Context, user *entity.User) (valueobject.UserID, error) {
	result, err := r.queries.InsertUser(ctx, query.InsertUserParams{
		Email:        user.Email().String(),
		PasswordHash: user.PasswordHash(),
		CreatedAt:    user.CreatedAt(),
	})
	if err != nil {
		return valueobject.UserID{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return valueobject.UserID{}, err
	}

	return valueobject.NewUserID(int(id)), nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	result, err := r.queries.SelectUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return r.toEntity(result), nil
}

func (r *UserRepository) FindByID(ctx context.Context, id valueobject.UserID) (*entity.User, error) {
	result, err := r.queries.SelectUserByID(ctx, int64(id.Int()))
	if err != nil {
		return nil, errors.New("user not found")
	}
	return r.toEntity(result), nil
}

func (r *UserRepository) toEntity(row query.User) *entity.User {
	return entity.ReconstructUser(
		valueobject.NewUserID(int(row.ID)),
		row.Email,
		row.PasswordHash,
		row.CreatedAt,
	)
}
