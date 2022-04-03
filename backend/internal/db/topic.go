package db

import (
	"bbs-app/backend/ent/topic"
	"context"
	"time"
)

type Topic struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetTopicAll() ([]Topic, error) {
	client, err := connection()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	searched, err := client.Topic.
		Query().
		All(context.Background())
	if err != nil {
		return nil, err
	}

	topics := make([]Topic, len(searched))
	for i, topic := range searched {
		topics[i] = Topic{
			ID:        topic.ID,
			Name:      topic.Name,
			Detail:    topic.Detail,
			CreatedAt: topic.CreatedAt,
			UpdatedAt: topic.UpdatedAt,
		}
	}

	return topics, nil
}

func (t *Topic) Create() (*Topic, error) {
	client, err := connection()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	created, err := client.Topic.
		Create().
		SetName(t.Name).
		SetDetail(t.Detail).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &Topic{
		ID:        created.ID,
		Name:      created.Name,
		Detail:    created.Detail,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.UpdatedAt,
	}, nil
}

func (t *Topic) Get() (*Topic, error) {
	client, err := connection()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	searched, err := client.Topic.
		Query().
		Where(topic.IDIn(t.ID)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return &Topic{
		ID:        searched.ID,
		Name:      searched.Name,
		Detail:    searched.Detail,
		CreatedAt: searched.CreatedAt,
		UpdatedAt: searched.UpdatedAt,
	}, nil
}

func (t *Topic) Update(name, detail string) (*Topic, error) {
	client, err := connection()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	updated, err := client.Topic.
		UpdateOneID(t.ID).
		SetName(name).
		SetDetail(detail).
		SetUpdatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &Topic{
		ID:        updated.ID,
		Name:      updated.Name,
		Detail:    updated.Detail,
		CreatedAt: updated.CreatedAt,
		UpdatedAt: updated.UpdatedAt,
	}, nil
}

func (t *Topic) Delete() error {
	client, err := connection()
	if err != nil {
		return err
	}
	defer client.Close()

	return client.Topic.
		DeleteOneID(t.ID).
		Exec(context.Background())
}
