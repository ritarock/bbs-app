package domain

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestComment_Validate(t *testing.T) {
	tests := []struct {
		name     string
		comment  Comment
		hasError bool
	}{
		{
			name: "pass",
			comment: Comment{
				PostID:      1,
				Content:     strings.Repeat("c", 255),
				CommentedAt: time.Now(),
			},
			hasError: false,
		},
		{
			name: "failed: post_id is not set",
			comment: Comment{
				Content:     strings.Repeat("c", 255),
				CommentedAt: time.Now(),
			},
			hasError: true,
		},
		{
			name: "failed: content length = 0",
			comment: Comment{
				Content:     strings.Repeat("c", 0),
				CommentedAt: time.Now(),
			},
			hasError: true,
		},
		{
			name: "failed: content length = 256",
			comment: Comment{
				Content:     strings.Repeat("c", 256),
				CommentedAt: time.Now(),
			},
			hasError: true,
		},
		{
			name: "failed: time is not set",
			comment: Comment{
				Content: strings.Repeat("c", 255),
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.comment.Validate()
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
