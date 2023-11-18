package timeutil

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	ctx := context.Background()
	now := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	c := SetMockNow(ctx, now)

	ti := Now(c)
	assert.Equal(t, now, ti)
}
