package streams

type OpData[T any] struct {
	MetaData string
	Function func(T) bool
}

type CompiledQueryable[T any] struct {
	Operators []LingoOperator[T]
	Items     *[]T
}
type LingoOperator[T any] struct {
	MetaData     OpData[T]
	OperatorType int
}
