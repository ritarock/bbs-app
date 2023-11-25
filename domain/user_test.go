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
			name:     "pass",
			user:     User{Name: "user", Password: "password1234"},
			hasError: false,
		},
		{
			name: "validate error: name",
			user: User{
				Name:     strings.Repeat("user", 10),
				Password: "password1234",
			},
			hasError: true,
		},
		{
			name: "validate error: password",
			user: User{
				Name:     "user",
				Password: "pass",
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
