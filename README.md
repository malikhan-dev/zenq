What Is Lingo?

As Of Now Lingo Is A Set Of Functions or Extensions That Can Help You Querying Data from any Structures. Something Like Linq In C# Or Streams In Java. Detailed Documentation Available In Wiki Section. This Project Is An Open Source Project Under No Guarantees, Though We Use Test Driven Approach To Ensure Stable Releases.



Rich Syntax

Ease Of Use

Open Source

Fast Development

Integrated Set Of Tools


Our Benchmark Shows The Following Lines Of Code Took About 4.8 seconds To Run In A Slice Of 50,000,000 records. And The Data Validates. The Performance Looks Solid.
	
	Benchmark Pc Specs:
	Laptop 
	Intel Core I7 12700
	16Gb Of Ram
	SSD
	Linux Ubuntu
    8 Seconds < 1.3.0
	4.8 Seconds On Latest Version
	
The Benchmark Included In The Test File


	res, err := From(items).Where("Flag", true).Filter(func(item ComplexObjectToSearch) bool {
		return item.Id > 200000
	}).AllOrDefault()

	if err != nil {
		b.Error(err)

	}

	if Any(*res, func(item ComplexObjectToSearch) bool {
		return !item.Flag
	}) {
		b.Error("Wrong Data Fetched")
	}





(V1.3.2) Or Higher (Breaking Change)

---Introducing Collectors---

After A Chained Operation Like Where(etc).AllOrDefault(), If We Want To Have The Result And Unwrap Queryable To A Tuple Of []T, and []err... We Use Collect And CollectRange functions. Its Important To Know That After Calling Collect The Result Is Not A Pointer Anymore. And Methods Like All And AllOrDefault Returns Queryable From Now On.



	res, err := From(items).Filter(func(item ComplexObjectToSearch) bool {
		return item.Id > 200000
	}).AllOrDefault().CollectRange(200)

	res2, err2 := From(items).Filter(func(item ComplexObjectToSearch) bool {
		return item.Id > 200000
	}).AllOrDefault().Collect()



