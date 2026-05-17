package Thor

import (
	"github.com/malikhan-dev/zenq/collections"
	"github.com/malikhan-dev/zenq/contracts"
)

// Hi My Name Is Thor. Im The Collections Query Engine Of Lingo

// Author: Mohammadreza Malikhan

type OperatorType int

const (
	FromItems       = 1
	WhereCollection = 2
	AnyCollection   = 4
	GroupCollection = 5
)

func From[T any](items []T) *CollectionCompiledQueryable[T] {

	initiateOperator := make([]contracts.LingoOperator[T], 0)
	initiateOperator = append(initiateOperator, contracts.LingoOperator[T]{
		OperatorType: FromItems,
		MetaData: contracts.OpData[T]{
			MetaData: "from",
			Function: func(t T) bool {
				return true
			},
		},
	})
	queryData := contracts.CompiledQueryable[T]{
		Items:     &items,
		Operators: initiateOperator,
	}

	return &CollectionCompiledQueryable[T]{
		Queryable: queryData,
	}
}
func (op *CollectionCompiledQueryable[T]) Where(function func(T) bool) *CollectionCompiledQueryable[T] {

	op.Queryable.Operators = append(op.Queryable.Operators, contracts.LingoOperator[T]{
		OperatorType: WhereCollection,
		MetaData: contracts.OpData[T]{
			MetaData: "where",
			Function: function,
		},
	})
	return op
}

func (op *CollectionCompiledQueryable[T]) Any(function func(T) bool) *AssertCompiledQueryable[T] {

	op.Queryable.Operators = append(op.Queryable.Operators, contracts.LingoOperator[T]{
		OperatorType: AnyCollection,
		MetaData: contracts.OpData[T]{
			MetaData: "any",
			Function: function,
		},
	})
	return &AssertCompiledQueryable[T]{
		Queryable: op.Queryable,
	}
}

func Group[K comparable, T any](op *CollectionCompiledQueryable[T], locator func(T) K) *GroupCompiledQueryable[K, T] {

	op.Queryable.Operators = append(op.Queryable.Operators, contracts.LingoOperator[T]{
		OperatorType: GroupCollection,
		MetaData: contracts.OpData[T]{
			MetaData: "group",
			Function: func(t T) bool {
				return true
			},
		},
	})
	return &GroupCompiledQueryable[K, T]{
		Queryable:   op.Queryable,
		PropLocator: locator,
	}
}

func (op *AssertCompiledQueryable[T]) Assert() bool {

	Any := false

	for _, item := range *op.Queryable.Items {

		for _, op := range op.Queryable.Operators {

			switch op.OperatorType {

			case AnyCollection:
				if op.MetaData.Function(item) {
					return true

				}

			}

		}
	}
	return Any
}

func (op *CollectionCompiledQueryable[T]) Collect() []T {

	var result []T

	result = make([]T, 0)

	for _, item := range *op.Queryable.Items {

		keep := true

		for _, op := range op.Queryable.Operators {

			switch op.OperatorType {

			case FromItems:
				if !op.MetaData.Function(item) {
					keep = false
					break
				}

			case WhereCollection:
				if !op.MetaData.Function(item) {
					keep = false
					break
				}

			}

		}

		if keep {
			result = append(result, item)
		}

	}
	return result
}

func Collect[K comparable, T any](op *GroupCompiledQueryable[K, T]) *collections.GroupedQueryable[K, T] {

	var result collections.GroupedQueryable[K, T]

	result.Items = make(map[K][]T)

	var LocatedKey K

	for _, item := range *op.Queryable.Items {

		LocatedKey = op.PropLocator(item)

		keep := true

		for _, operator := range op.Queryable.Operators {

			switch operator.OperatorType {

			case FromItems:
				if !operator.MetaData.Function(item) {
					keep = false
					break
				}

			case WhereCollection:
				if !operator.MetaData.Function(item) {
					keep = false
					break
				}
			}
		}

		if !keep {
			continue
		}
		/*
			v := reflect.ValueOf(item)

			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}

			field := v.FieldByIndex(targetField.Index)

			if !field.IsValid() {
				continue
			}

			key, ok := field.Interface().(K)

			if !ok {
				continue
			}*/

		result.Items[LocatedKey] = append(result.Items[LocatedKey], item)
	}

	return &result
}
