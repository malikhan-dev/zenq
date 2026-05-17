package streams

import (
	"github.com/malikhan-dev/zenq/contracts"
	"context"
)

func CompileFromQueryable[T any](items []T) *contracts.CompiledQueryable[T] {

	var result contracts.CompiledQueryable[T]

	result.Operators = make([]contracts.LingoOperator[T], 0)

	result.Items = &items

	var operator contracts.LingoOperator[T]

	operator.OperatorType = 1

	operator.MetaData = contracts.OpData[T]{
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
