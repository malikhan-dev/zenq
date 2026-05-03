What Is Lingo?

As Of Now Lingo Is A Set Of Functions or Extensions That Can Help You Querying Data from any Structures. Something Like Linq In C# Or Streams In Java. 








1 - Predicate Function...

	result := Predicate(items, "Name", "John")


2 - any function...

	result := Any(items, func(item ComplexObjectToSearch) bool {
		return item.Flag
	})


3 - RemoveFirstByPredicate

	dest = RemoveFirstByPredicate(*dest, func(search ComplexObjectToSearch) bool {
		return search.Name == "Jane"
	})


4- CopyToSlice

  var dest *[]ComplexObjectToSearch
	dest = CopyToSlice(items)

5- FilterFirst

	item := FilterFirst(items, func(search ComplexObjectToSearch) bool {
		return search.Flag == true
	})


6- Filter Function

  var newItems *[]ComplexObjectToSearch

	newItems = Filter(items, func(search ComplexObjectToSearch) bool {
		return search.Flag == true
	})



and many more handy tools to come.

  
