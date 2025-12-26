package valueobject_test

import (
	"strings"
	"testing"

	"github.com/ritarock/bbs-app/domain/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestNewPostTitle(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		title    string
		hasError bool
	}{
		{
			name:     "pass: 1 char",
			title:    strings.Repeat("t", 1),
			hasError: false,
		},
		{
			name:     "pass: 30 chars",
			title:    strings.Repeat("t", 30),
			hasError: false,
		},
		{
			name:     "failed: empty",
			title:    "",
			hasError: true,
		},
		{
			name:     "failed: 31 chars",
			title:    strings.Repeat("t", 31),
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := valueobject.NewPostTitle(test.title)
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.title, got.String())
			}
		})
	}
}
