package main

import "fmt"

func main() {

	// array
	my_array := [5]int{4, 6, 5, 2, 4}

	for index, value := range my_array {
		fmt.Println(index, "->", value)
	}

	// slices
	my_slice := []int{23, 65, 33, 77, 33, 78, 56}
	for i, v := range my_slice {
		fmt.Println(i, "->", v)
	}

	// in slices we can just use a range of values
	fmt.Println(my_slice[1:3])

	// appending to a slice

	my_slice = append(my_slice, 100)
	fmt.Println(my_slice)

	// appending two slices
	myslice := []string{"I ", "am "}
	myother_slice := []string{"Sachith ", "Muhandiram"}

	myslice = append(myslice, myother_slice...)

	fmt.Println(myslice) // [I  am  Sachith  Muhandiram]

	// slice value selection
	//				0, 1, 2, 3, 4
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9}

	fmt.Println("Slice 1", slice1)
	fmt.Println("Value after index 2 -> ", slice1[2:]) // [3 4 5]

	fmt.Println("Slice 2", slice2)
	fmt.Println("Value upto index 2 -> ", slice2[:2])

} // end of main function
