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
			name: "pass",
			post: Post{
				Title:    strings.Repeat("t", 30),
				Content:  strings.Repeat("c", 255),
				PostedAt: time.Now(),
			},
			hasError: false,
		},
		{
			name: "failed: title length = 0",
			post: Post{
				Title:    "",
				Content:  strings.Repeat("c", 255),
				PostedAt: time.Now(),
			},
			hasError: true,
		},
		{
			name: "failed: title length > 30",
			post: Post{
				Title:    strings.Repeat("t", 31),
				Content:  strings.Repeat("c", 255),
				PostedAt: time.Now(),
			},
			hasError: true,
		},
		{
			name: "failed: content length = 0",
			post: Post{
				Title:    strings.Repeat("t", 30),
				Content:  "",
				PostedAt: time.Now(),
			},
			hasError: true,
		},
		{
			name: "failed: content length > 255",
			post: Post{
				Title:    strings.Repeat("t", 30),
				Content:  strings.Repeat("c", 256),
				PostedAt: time.Now(),
			},
			hasError: true,
		},
		{
			name: "failed: time is not set",
			post: Post{
				Title:   strings.Repeat("t", 30),
				Content: strings.Repeat("c", 255),
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
