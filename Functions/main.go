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
}
