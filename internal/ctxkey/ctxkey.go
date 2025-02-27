// This code is inspired by https://hypirion.com/musings/spectral-contexts-in-go

package ctxkey

import (
	"context"
)

type key[T any] struct{}

func WithKey[T any](ctx context.Context, v T) context.Context {
	return context.WithValue(ctx, key[T]{}, v)
}

func From[T any](ctx context.Context) T {
	v, _ := ctx.Value(key[T]{}).(T)
	return v
}
