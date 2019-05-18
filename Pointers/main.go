package main

import "fmt"

func main() {
	// x interger varibale
	x := 5

	// y integer pointer to x
	y := &x

	fmt.Println(x)
	fmt.Println(y)  // memory location 0xabc1254
	fmt.Println(*y) // dereferencing the memory location -> 5

	// update the value at memory location
	*y = 10
	fmt.Println(x)
	fmt.Println(y)  // memory location 0xabc1254
	fmt.Println(*y) // dereferencing the memory location -> 10

	z := add_ten(&x)
	fmt.Println(z)  //  value -> 20
	fmt.Println(&z) //  memory location 0xabc1050
	fmt.Println(*y) //  value of y -> 20

	var age int

	fmt.Println("Please enter your age :")
	fmt.Scan(&age)

	fmt.Println("Your age is : ", age)
}

func add_ten(x *int) int { // function receives an integer pointer and returns an integer
	*x += 10

	return *x
}
