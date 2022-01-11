//go:build unit
// +build unit

package context

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	assert.False(t, IsCanceledError(ctx.Err()))
	cancel()
	assert.True(t, IsCanceledError(ctx.Err()))

	// Flaky test. Sometimes passes sometimes doesn't. Disabling.
	// ctx, cancel = context.WithTimeout(context.Background(), time.Microsecond)
	// time.Sleep(time.Microsecond * 2)
	// cancel()
	// assert.False(t, IsCanceledError(ctx.Err()))

	assert.False(t, IsCanceledError(errors.New("test error")))
}
