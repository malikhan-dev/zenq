package Thor


import "github.com/malikhan-dev/lingo/contracts"

type CollectionCompiledQueryable[T any] struct {
	Queryable contracts.CompiledQueryable[T]
}

type AssertCompiledQueryable[T any] struct {
	Queryable contracts.CompiledQueryable[T]
}
type GroupCompiledQueryable[K comparable, T any] struct {
	Queryable         contracts.CompiledQueryable[T]
	GroupPropertyName string
}
