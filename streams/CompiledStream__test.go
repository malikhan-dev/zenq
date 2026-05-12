package streams

import (
	"context"
	"fmt"
	"testing"
)

var data []ComplexObjectToSearch

const StressTest = true

func init() {

	items = []ComplexObjectToSearch{
		ComplexObjectToSearch{
			Name: "John",
			Age:  20,
			Id:   1,
			Flag: true,
		},
		ComplexObjectToSearch{
			Name: "Jane",
			Age:  20,
			Id:   2,
			Flag: false,
		},
		ComplexObjectToSearch{
			Name: "Jane",
			Age:  20,
			Id:   3,
			Flag: true,
		},
		ComplexObjectToSearch{
			Name: "jack",
			Age:  20,
			Id:   4,
			Flag: true,
		},
	}

	if StressTest {
		LoadLargeData()
	}

}

func TestCompiledQuery(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	type student struct {
		Id   int
		Name string
		Age  int
	}

	counter := 1

	for i := range CompileFromQueryable(items).
		Filter(func(student ComplexObjectToSearch) bool {
			return !student.Flag
		}).
		Throttle(0).Stream(ctx) {

		fmt.Println(i)
		if counter >= 10 {
			cancel()
			break
		} else {
			counter++

		}

	}

	cancel()

}

func TestCompiledQueryWithMapping(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	type student struct {
		Id   int
		Name string
		Age  int
	}
	counter := 1

	for i := range StreamWithMapping(CompileFromQueryable(items).
		Filter(func(student ComplexObjectToSearch) bool {
			return !student.Flag
		}).
		Throttle(0),
		func(items ComplexObjectToSearch) student {
			return student{
				Id:   items.Id,
				Name: items.Name + " student",
				Age:  items.Age,
			}
		}, ctx) {
		fmt.Println(i)
		if counter >= 10 {
			cancel()
			break
		} else {
			counter++

		}
	}

	cancel()
}
