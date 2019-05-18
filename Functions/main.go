package main

import "fmt"

func foo() {
	fmt.Println("Hello from foo function")
}

func bar(x int) int {
	y := x + 10
	return y
}

func hello() {
	fmt.Println("Hello ")
}

func world() {
	fmt.Println("World")
}

// variadic function
func variadic(v ...int) int {
	total := 0
	for _, v := range v {
		total += v
	}

	return total
}

// variadic arguments
func variadicArg(v ...int) int {
	fmt.Println("Value : ", v)
	total := 0
	for _, v := range v {
		total += v
	}

	return total
}

func main() {

	foo()
	bar_val := bar(8) // we have to assign to a varible as bar() returns an integer
	fmt.Println("Bar value : ", bar_val)

	// anonymus function
	my_anonymus := func(x int) int {
		return x + 10
	}

	my_val := my_anonymus(1)
	fmt.Println("Anonymus value : ", my_val)

	// variadic with values
	fmt.Println(variadic()) // zero paramters -> 0

	fmt.Println(variadic(10, 20)) // -> 30

	fmt.Println(variadic(10, 89, 54, 3, 1, 22, 56, 15, 656, 23, 88, 02, 85))

	// variadic arguments
	var_data := []int{12, 85, 8, 96, 55, 54} // int slice
	fmt.Println(variadicArg(var_data...))    // variadicArg expecting an int

	//calling hello() and world() functions
	world()
	hello()

	// with defer
	defer world() // in this block, this is executed just before main() ends
	hello()

} // end of main function
