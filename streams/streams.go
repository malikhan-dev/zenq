package streams

import (
	"context"
	"time"
)

func FilterStream[T any](ctx context.Context, BufferSize int, in <-chan T, predicate func(T) bool) <-chan T {
	out := make(chan T, BufferSize)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}

				if predicate(v) {
					select {
					case <-ctx.Done():
						return
					case out <- v:

					}
				}
			}
		}
	}()

	return out
}

func MapStream[T any, M any](ctx context.Context, in <-chan T, mapper func(T) M) <-chan M {
	out := make(chan M)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return

			case v, ok := <-in:
				if !ok {
					return
				}

				m := mapper(v)

				select {
				case <-ctx.Done():
					return
				case out <- m:
				}
			}
		}
	}()

	return out
}

func Throttle[T any](ctx context.Context, in <-chan T, duration time.Duration) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}

				select {
				case <-ctx.Done():
					return
				case out <- v:
					time.Sleep(duration * time.Millisecond)
				}

			}
		}
	}()

	return out
}
