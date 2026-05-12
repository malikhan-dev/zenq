package streams

import (
	"time"
)

type OperatorType int

const (
	BuildQueryable    = 1
	FilterQueryable   = 2
	ThrottleQueryable = 3
)

func (currentOps *CompiledQueryable[T]) Filter(filter func(item T) bool) *CompiledQueryable[T] {

	currentOps.Operators = append(currentOps.Operators, LingoOperator[T]{
		MetaData: OpData[T]{
			MetaData: "FFilter",
			Function: filter,
		},
		OperatorType: 2,
	})

	return currentOps
}

func (currentOps *CompiledQueryable[T]) Throttle(duration time.Duration) *CompiledQueryable[T] {

	currentOps.Operators = append(currentOps.Operators, LingoOperator[T]{
		MetaData: OpData[T]{
			MetaData: "0",
			Function: func(item T) bool {
				time.Sleep(duration)
				return true
			},
		},
		OperatorType: 3,
	})

	return currentOps
}
