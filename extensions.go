package lingo

import (
	"errors"
	"fmt"
	"reflect"
)

func FindByPredicate[T any](items []T, predicate func(T) bool) *[]T {

	var result []T

	for _, v := range items {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return &result

}

func FindFirstByPredicate[T any](items []T, predicate func(T) bool) *T {

	var result T

	for _, v := range items {
		if predicate(v) {
			result = v
			break
		}
	}

	return &result
}

func RemoveFirstByPredicate[T any](items []T, predicate func(T) bool) *[]T {

	var result []T

	conditionMet := false

	for _, v := range items {

		if predicate(v) && !conditionMet {
			conditionMet = true
			continue

		} else {

			result = append(result, v)

		}

	}

	return &result
}

func RemoveByPredicate[T any](items []T, predicate func(T) bool) *[]T {

	var result []T

	for _, v := range items {
		if predicate(v) {
			continue
		} else {
			result = append(result, v)
		}

	}

	return &result
}

func (query *Queryable[T]) Filter(predicate func(T) bool) *Queryable[T] {

	var result []T
	result = make([]T, 0)

	for _, v := range query.Items {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return &Queryable[T]{
		Items: result,
		err:   nil,
	}
}

func From[T any](items []T) *Queryable[T] {

	return &Queryable[T]{
		Items: items,
		err:   nil,
	}
}

func Any[T any](Items []T, Condition func(T) bool) bool {

	for _, v := range Items {

		if Condition(v) {
			return true
		}
	}
	return false
}

func (query *Queryable[T]) Where(fieldName string, fieldValue any) *Queryable[T] {

	fnc := func(t *Queryable[T]) *Queryable[T] {

		var Out Queryable[T]

		strType := reflect.TypeFor[T]()

		if strType.Kind() != reflect.Struct {
			Out.err = append(Out.err, errors.New("Expected a struct"))
		}

		if strType.Kind() == reflect.Ptr {
			strType = strType.Elem()
		}

		field, ok := strType.FieldByName(fieldName)

		newItems := make([]T, 0)

		if ok {

			for _, val := range query.Items {

				v := reflect.ValueOf(val)

				f := v.FieldByIndex(field.Index)

				if f.Interface() == fieldValue {
					newItems = append(newItems, val)
				}
			}

		} else {
			Out.err = append(Out.err, errors.New(fmt.Sprintf("Invalid Property Name: %s", fieldName)))
		}
		for _, val := range t.err {

			Out.err = append(Out.err, val)
		}

		Out.Items = newItems
		return &Out
	}

	return fnc(query)
}

func FirstOrDefault[T any](items *Queryable[T]) (*T, []error) {

	var result T

	if len(items.Items) > 0 {
		result = items.Items[0]
	} else {
		items.err = append(items.err, errors.New("No items found"))
	}
	return &result, items.err
}

func All[T any](items *Queryable[T]) (*[]T, []error) {
	if len(items.Items) > 0 {
		return &items.Items, items.err
	} else {
		panic(errors.New("Cant Perform All() On Empty Slice"))
	}
}

func First[T any](items *Queryable[T]) (*T, []error) {
	if len(items.Items) > 0 {
		return &items.Items[0], items.err
	} else {
		panic(errors.New("Cant Perform First() On Empty Slice"))
	}
}

func AllOrDefault[T any](items *Queryable[T]) (*[]T, []error) {
	if len(items.Items) > 0 {
		return &items.Items, items.err
	} else {
		var toReturn *[]T
		items.err = append(items.err, errors.New("No Result."))
		return toReturn, items.err
	}
}
