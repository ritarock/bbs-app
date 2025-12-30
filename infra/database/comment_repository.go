package database

import (
	"context"

	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/repository"
	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/ritarock/bbs-app/infra/database/query"
)

type CommentRepository struct {
	queries *query.Queries
}

func NewCommentRepository(db query.DBTX) repository.CommentRepository {
	return &CommentRepository{
		queries: query.New(db),
	}
}

func (r *CommentRepository) Save(ctx context.Context, comment *entity.Comment) (valueobject.CommentID, error) {
	result, err := r.queries.InsertComment(ctx, query.InsertCommentParams{
		PostID:      int64(comment.PostID().Int()),
		Body:        comment.Body().String(),
		CommentedAt: comment.CommentedAt(),
	})
	if err != nil {
		return valueobject.CommentID{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return valueobject.CommentID{}, err
	}

	return valueobject.NewCommentID(int(id)), nil
}

func (r *CommentRepository) FindByID(ctx context.Context, id valueobject.CommentID) (*entity.Comment, error) {
	result, err := r.queries.SelectComment(ctx, int64(id.Int()))
	if err != nil {
		return nil, err
	}
	return r.toEntity(result), nil
}

func (r *CommentRepository) FindAll(ctx context.Context) ([]*entity.Comment, error) {
	result, err := r.queries.SelectCommentsByPostId(ctx, 0)
	if err != nil {
		return nil, err
	}

	comments := make([]*entity.Comment, len(result))
	for i, row := range result {
		comments[i] = r.toEntity(row)
	}

	return comments, nil
}

func (r *CommentRepository) FindByPostID(ctx context.Context, postID valueobject.PostID) ([]*entity.Comment, error) {
	result, err := r.queries.SelectCommentsByPostId(ctx, int64(postID.Int()))
	if err != nil {
		return nil, err
	}

	comments := make([]*entity.Comment, len(result))
	for i, row := range result {
		comments[i] = r.toEntity(row)
	}

	return comments, nil
}

func (r *CommentRepository) Update(ctx context.Context, comment *entity.Comment) error {
	return r.queries.UpdateComment(ctx, query.UpdateCommentParams{
		ID:   int64(comment.ID().Int()),
		Body: comment.Body().String(),
	})
}

func (r *CommentRepository) Delete(ctx context.Context, id valueobject.CommentID) error {
	return r.queries.DeleteComment(ctx, int64(id.Int()))
}

func (r *CommentRepository) toEntity(row query.Comment) *entity.Comment {
	return entity.ReconstructComment(
		valueobject.NewCommentID(int(row.ID)),
		valueobject.NewPostID(int(row.PostID)),
		row.Body,
		row.CommentedAt,
	)
}
