package timeutil

import (
	"context"
	"time"
)

type ctxKey int

const (
	CtxFreezeTimeKey ctxKey = iota
)

func Now(ctx context.Context) time.Time {
	_, ok := ctx.Value(CtxFreezeTimeKey).(time.Time)
	if !ok {
		ctx = setNow(ctx)
	}
	return ctx.Value(CtxFreezeTimeKey).(time.Time)
}

func setNow(ctx context.Context) context.Context {
	return context.WithValue(ctx, CtxFreezeTimeKey, time.Now())
}

func SetMockNow(ctx context.Context, mockTime time.Time) context.Context {
	return context.WithValue(ctx, CtxFreezeTimeKey, mockTime)
}
