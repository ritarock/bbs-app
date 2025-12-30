package valueobject

import "errors"

type CommentBody struct {
	value string
}

func NewCommentBody(body string) (CommentBody, error) {
	if len(body) < 1 {
		return CommentBody{}, errors.New("body must be at least 1 character")
	}
	if len(body) > 500 {
		return CommentBody{}, errors.New("body must be at most 500 characters")
	}

	return CommentBody{value: body}, nil
}

func ReconstructCommentBody(body string) CommentBody {
	return CommentBody{value: body}
}

func (b CommentBody) String() string {
	return b.value
}
