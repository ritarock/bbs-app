package database

import (
	"context"

	"github.com/ritarock/bbs-app/domain/entity"
	"github.com/ritarock/bbs-app/domain/repository"
	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/ritarock/bbs-app/infra/database/query"
)

type PostRepository struct {
	queries *query.Queries
}

func NewPostRepository(db query.DBTX) repository.PostRepository {
	return &PostRepository{
		queries: query.New(db),
	}
}

func (p *PostRepository) Save(ctx context.Context, post *entity.Post) (valueobject.PostID, error) {
	result, err := p.queries.InsertPost(ctx, query.InsertPostParams{
		Title:     post.Title().String(),
		Content:   post.Content().String(),
		CreatedAt: post.CreatedAt(),
	})
	if err != nil {
		return valueobject.PostID{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return valueobject.PostID{}, err
	}

	return valueobject.NewPostID(int(id)), nil
}

func (p *PostRepository) FindByID(ctx context.Context, id valueobject.PostID) (*entity.Post, error) {
	result, err := p.queries.SelectPost(ctx, int64(id.Int()))
	if err != nil {
		return nil, err
	}
	return p.toEntity(result), nil
}

func (p *PostRepository) FindAll(ctx context.Context) ([]*entity.Post, error) {
	result, err := p.queries.SelectPosts(ctx)
	if err != nil {
		return nil, err
	}

	posts := make([]*entity.Post, len(result))
	for i, row := range result {
		posts[i] = p.toEntity(row)
	}

	return posts, nil
}

func (p *PostRepository) Update(ctx context.Context, post *entity.Post) error {
	return p.queries.UpdatePost(ctx, query.UpdatePostParams{
		ID:      int64(post.ID().Int()),
		Title:   post.Title().String(),
		Content: post.Content().String(),
	})
}

func (p *PostRepository) Delete(ctx context.Context, id valueobject.PostID) error {
	return p.queries.DeletePost(ctx, int64(id.Int()))
}

func (p *PostRepository) toEntity(row query.Post) *entity.Post {
	return entity.ReconstructPost(
		valueobject.NewPostID(int(row.ID)),
		row.Title,
		row.Content,
		row.CreatedAt,
	)
}
