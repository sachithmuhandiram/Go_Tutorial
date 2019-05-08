package main

func main()  {
	
	// array
	my_array := [5]int{4,6,5,2,4}

	for index,value := range my_array{
		fmt.Println(index, "->", value)
	}

	// slices
	my_slice := []int{23,65,33,77,33,78,56}
	for i,v := range my_slice{
		fmt.Println(i, "->", v)
	}

	// in slices we can just use a range of values
	fmt.Println(my_slice[1:3])

	// appending to a slice

	my_slice = append(my_slice,100)
	fmt.Println(my_slice)

}