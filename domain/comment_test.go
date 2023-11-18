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
			name:     "pass",
			comment:  Comment{Content: "content", CommentedAt: time.Now(), PostId: 1},
			hasError: false,
		},
		{
			name: "validate error: content",
			comment: Comment{
				Content:     strings.Repeat("content", 100),
				CommentedAt: time.Now(),
				PostId:      1,
			},
			hasError: true,
		},
		{
			name: "validate error: time",
			comment: Comment{
				Content: "content",
				PostId:  1,
			},
			hasError: true,
		},
		{
			name: "validate error: time",
			comment: Comment{
				Content:     "content",
				CommentedAt: time.Now(),
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
