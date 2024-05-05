package domain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name     string
		user     User
		hasError bool
	}{
		{
			name: "pass",
			user: User{
				Name:     strings.Repeat("n", 30),
				Password: strings.Repeat("p", 8),
			},
			hasError: false,
		},
		{
			name: "failed: name length = 0",
			user: User{
				Name:     strings.Repeat("n", 0),
				Password: strings.Repeat("p", 8),
			},
			hasError: true,
		},
		{
			name: "failed: name length = 31",
			user: User{
				Name:     strings.Repeat("n", 31),
				Password: strings.Repeat("p", 8),
			},
			hasError: true,
		},
		{
			name: "failed: password length = 7",
			user: User{
				Name:     strings.Repeat("n", 30),
				Password: strings.Repeat("p", 7),
			},
			hasError: true,
		},
		{
			name: "failed: password length = 31",
			user: User{
				Name:     strings.Repeat("n", 30),
				Password: strings.Repeat("p", 31),
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.user.Validate()
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
