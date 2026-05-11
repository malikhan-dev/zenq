package lingo

import (
	"testing"
)

func TestCopyTo(t *testing.T) {
	var dest *[]ComplexObjectToSearch

	dest = DeepCopy(items)

	dest = RemoveFirstByPredicate(*dest, func(search ComplexObjectToSearch) bool {
		return search.Name == "Jane"
	})

}

func TestToSlice(t *testing.T) {

	var m map[int]ComplexObjectToSearch

	m = make(map[int]ComplexObjectToSearch, 2)

	m[1] = ComplexObjectToSearch{
		Name: "joe",
		Age:  20,
		Id:   1,
		Flag: false,
	}
	m[2] = ComplexObjectToSearch{
		Name: "john",
		Age:  27,
		Id:   2,
		Flag: true,
	}

	_ = ToSlice(m)

}
