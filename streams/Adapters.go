package streams

import (
	"context"

	lingo "github.com/malikhan-dev/lingo"
)

func CompileFromQueryable[T any](items []T) *CompiledQueryable[T] {

	var result CompiledQueryable[T]

	result.Operators = make([]LingoOperator[T], 0)

	result.Items = &items

	var operator LingoOperator[T]

	operator.OperatorType = 1

	operator.MetaData = OpData[T]{
		MetaData: "FromQueryable",
		Function: func(item T) bool {
			return true
		},
	}
	return &result
}

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
