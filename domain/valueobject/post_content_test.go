package valueobject_test

import (
	"strings"
	"testing"

	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestNewPostContent(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		content  string
		hasError bool
	}{
		{
			name:     "pass: 1 char",
			content:  strings.Repeat("c", 1),
			hasError: false,
		},
		{
			name:     "pass: 255 chars",
			content:  strings.Repeat("c", 255),
			hasError: false,
		},
		{
			name:     "failed: empty",
			content:  "",
			hasError: true,
		},
		{
			name:     "failed: 256 chars",
			content:  strings.Repeat("c", 256),
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := valueobject.NewPostContent(test.content)
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.content, got.String())
			}
		})
	}
}
