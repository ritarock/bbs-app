package valueobject

import "errors"

type PostContent struct {
	value string
}

func NewPostContent(content string) (PostContent, error) {
	if len(content) < 1 {
		return PostContent{}, errors.New("content must be at least 1 character")
	}
	if len(content) > 255 {
		return PostContent{}, errors.New("content must be at most 255 characters")
	}

	return PostContent{value: content}, nil
}

func ReconstructPostContent(content string) PostContent {
	return PostContent{value: content}
}

func (t PostContent) String() string {
	return t.value
}
