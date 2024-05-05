package domain

import (
	"errors"
	"time"
)

type Post struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	PostedAt time.Time `json:"posted_at"`
}

func (p *Post) Validate() error {
	if len(p.Title) == 0 || len(p.Title) > 30 {
		return errors.New("length of title must range from 1 to 30 characters")
	}
	if len(p.Content) == 0 || len(p.Content) > 255 {
		return errors.New("length of content must range from 1 to 255 characters")
	}
	if p.PostedAt.IsZero() {
		return errors.New("time is not set")
	}

	return nil
}
