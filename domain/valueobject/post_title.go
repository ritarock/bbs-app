package valueobject

import "errors"

type PostTitle struct {
	value string
}

func NewPostTitle(title string) (PostTitle, error) {
	if len(title) < 1 {
		return PostTitle{}, errors.New("title must be at least 1 character")
	}
	if len(title) > 30 {
		return PostTitle{}, errors.New("title must be at most 30 characters")
	}

	return PostTitle{value: title}, nil
}

func ReconstructPostTitle(title string) PostTitle {
	return PostTitle{value: title}
}

func (t PostTitle) String() string {
	return t.value
}
