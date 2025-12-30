package entity

import (
	"time"

	"github.com/ritarock/bbs-app/domain/valueobject"
)

type Comment struct {
	id          valueobject.CommentID
	postID      valueobject.PostID
	body        valueobject.CommentBody
	commentedAt time.Time
}

func NewComment(postID int, body string) (*Comment, error) {
	commentBody, err := valueobject.NewCommentBody(body)
	if err != nil {
		return nil, err
	}

	return &Comment{
		postID:      valueobject.NewPostID(postID),
		body:        commentBody,
		commentedAt: time.Now(),
	}, nil
}

func (c *Comment) ID() valueobject.CommentID {
	return c.id
}

func (c *Comment) PostID() valueobject.PostID {
	return c.postID
}

func (c *Comment) Body() valueobject.CommentBody {
	return c.body
}

func (c *Comment) CommentedAt() time.Time {
	return c.commentedAt
}

func (c *Comment) Update(body string) error {
	newBody, err := valueobject.NewCommentBody(body)
	if err != nil {
		return err
	}

	c.body = newBody
	return nil
}

func ReconstructComment(id valueobject.CommentID, postID valueobject.PostID, body string, commentedAt time.Time) *Comment {
	return &Comment{
		id:          id,
		postID:      postID,
		body:        valueobject.ReconstructCommentBody(body),
		commentedAt: commentedAt,
	}
}
