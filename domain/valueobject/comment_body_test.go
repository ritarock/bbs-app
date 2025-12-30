package valueobject_test

import (
	"strings"
	"testing"

	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestNewCommentBody(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		body     string
		hasError bool
	}{
		{
			name:     "pass: 1 char",
			body:     strings.Repeat("b", 1),
			hasError: false,
		},
		{
			name:     "pass: 500 chars",
			body:     strings.Repeat("b", 500),
			hasError: false,
		},
		{
			name:     "failed: empty",
			body:     "",
			hasError: true,
		},
		{
			name:     "failed: 501 chars",
			body:     strings.Repeat("b", 501),
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := valueobject.NewCommentBody(test.body)
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.body, got.String())
			}
		})
	}
}
