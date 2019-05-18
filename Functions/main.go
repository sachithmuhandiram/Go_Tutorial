package main

import "fmt"

func foo() {
	fmt.Println("Hello from foo function")
}

func bar(x int) int {
	y := x + 10
	return y
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
}
