package entity

import (
	"time"

	"github.com/ritarock/bbs-app/domain/valueobject"
)

type Post struct {
	id        valueobject.PostID
	title     valueobject.PostTitle
	content   valueobject.PostContent
	createdAt time.Time
}

func NewPost(title, content string) (*Post, error) {
	postTitle, err := valueobject.NewPostTitle(title)
	if err != nil {
		return nil, err
	}
	postContent, err := valueobject.NewPostContent(content)
	if err != nil {
		return nil, err
	}

	return &Post{
		title:     postTitle,
		content:   postContent,
		createdAt: time.Now(),
	}, nil
}

func (p *Post) ID() valueobject.PostID {
	return p.id
}

func (p *Post) Title() valueobject.PostTitle {
	return p.title
}

func (p *Post) Content() valueobject.PostContent {
	return p.content
}

func (p *Post) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Post) Update(title, content string) error {
	newTitle, err := valueobject.NewPostTitle(title)
	if err != nil {
		return err
	}
	newContent, err := valueobject.NewPostContent(content)
	if err != nil {
		return err
	}

	p.title = newTitle
	p.content = newContent
	return nil
}

func ReconstructPost(id valueobject.PostID, title, content string, createdAt time.Time) *Post {
	return &Post{
		id:        id,
		title:     valueobject.ReconstructPostTitle(title),
		content:   valueobject.ReconstructPostContent(content),
		createdAt: createdAt,
	}
}
