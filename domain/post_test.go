package domain

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPost_Validate(t *testing.T) {
	tests := []struct {
		name     string
		post     Post
		hasError bool
	}{
		{
			name:     "pass",
			post:     Post{Title: "title", Content: "content", PostedAt: time.Now()},
			hasError: false,
		},
		{
			name: "validate error: title",
			post: Post{
				Title:    strings.Repeat("title", 10),
				Content:  "content",
				PostedAt: time.Now(),
			},
			hasError: true,
		},
		{
			name: "validate error: content",
			post: Post{
				Title:    "title",
				Content:  strings.Repeat("content", 100),
				PostedAt: time.Now(),
			},
			hasError: true,
		},
		{
			name: "validate error: time",
			post: Post{
				Title:   "title",
				Content: "content",
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.post.Validate()
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
