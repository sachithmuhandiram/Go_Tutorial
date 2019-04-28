package main

import "fmt"

func main() {
	age := 29

	//if statement
	if age == 29 {
		fmt.Println("You are ", age)
	}

	// if else
	if age == 25 {
		fmt.Println("You are ", age)
	} else {
		fmt.Println("You are not 25 years old")
	}

	// if else if

	if age == 25 {
		fmt.Println("You are : ", age)
	} else if age == 27 {
		fmt.Println("You are 27 years old")
	} else {
		fmt.Println("You are 29 years old from else if")
	}

	// FOR loop

	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	// loop through a slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // integer slice

	for index, value := range numbers {
		fmt.Println(index, "->", value)
	}

}
