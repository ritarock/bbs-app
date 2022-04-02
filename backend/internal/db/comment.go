package db

import (
	"bbs-app/backend/ent/comment"
	"context"
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	TopicId   int       `json:"topic_id"`
}

func GetCommentAll() ([]Comment, error) {
	client, err := connection()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	searched, err := client.Comment.
		Query().
		All(context.Background())
	if err != nil {
		return nil, err
	}

	comments := make([]Comment, len(searched))
	for _, comment := range searched {
		comments = append(comments, Comment{
			ID:        comment.ID,
			Body:      comment.Body,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			TopicId:   comment.TopicID,
		})
	}

	return comments, nil
}

func GetCommentAllByTopic(topicId int) ([]Comment, error) {
	client, err := connection()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	searched, err := client.Comment.
		Query().
		Where(comment.TopicID(topicId)).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	comments := make([]Comment, len(searched))
	for _, comment := range searched {
		comments = append(comments, Comment{
			ID:        comment.ID,
			Body:      comment.Body,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			TopicId:   comment.TopicID,
		})
	}

	return comments, nil
}

func (c *Comment) Create() (*Comment, error) {
	client, err := connection()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	created, err := client.Comment.
		Create().
		SetBody(c.Body).
		SetTopicID(c.TopicId).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &Comment{
		ID:        created.ID,
		Body:      created.Body,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.UpdatedAt,
		TopicId:   created.TopicID,
	}, nil
}

func (c *Comment) Get() (*Comment, error) {
	client, err := connection()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	searched, err := client.Comment.
		Query().
		Where(comment.IDIn(c.ID)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return &Comment{
		ID:        searched.ID,
		Body:      searched.Body,
		CreatedAt: searched.CreatedAt,
		UpdatedAt: searched.UpdatedAt,
		TopicId:   searched.TopicID,
	}, nil
}

func (c *Comment) Update(body string) (*Comment, error) {
	client, err := connection()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	updated, err := client.Comment.
		UpdateOneID(c.ID).
		SetBody(body).
		SetUpdatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &Comment{
		ID:        updated.ID,
		Body:      updated.Body,
		CreatedAt: updated.CreatedAt,
		UpdatedAt: updated.UpdatedAt,
		TopicId:   updated.TopicID,
	}, err
}

func (c *Comment) Delete() error {
	client, err := connection()
	if err != nil {
		return err
	}
	defer client.Close()

	return client.Comment.
		DeleteOneID(c.ID).
		Exec(context.Background())
}
