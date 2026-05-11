package streams

import (
	lingo "Lingo/src"
	"context"
)

func FromData[T any](ctx context.Context, BufferSize int, items []T) <-chan T {
	out := make(chan T, BufferSize)

	go func() {
		defer close(out)

		for _, v := range items {
			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}()

	return out
}

func FromChannel[T any](ctx context.Context, BufferSize int, items <-chan T) <-chan T {
	out := make(chan T, BufferSize)

	go func() {
		defer close(out)

		for val := range items {
			select {
			case <-ctx.Done():
				return
			case out <- val:
			}
		}
	}()

	return out
}

func FromQueryable[T any](ctx context.Context, BufferSize int, items lingo.Queryable[T]) <-chan T {
	out := make(chan T, BufferSize)

	go func() {
		defer close(out)

		for _, v := range items.Items {

			select {
			case <-ctx.Done():
				return
			case out <- v:
			}

		}

	}()
	return out
}
