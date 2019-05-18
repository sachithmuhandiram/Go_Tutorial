package main

import (
	"fmt"
	"net/http"
)

// globle definations here
type new_type int

// constants
const (
	Name = "Sachith Muhandiram"
	Age  = 29
)

func main() {
	// long way
	// var variable data type
	var f_name string = "sachith"

	fmt.Println("Hello ,your first name is ", f_name)
	//short way
	// variable := value
	l_name := "muhandiram"
	fmt.Println("Hello your last name is ", l_name)

	type my_type int

	var age my_type
	age = 29

	fmt.Println("Age is : ", age)
	fmt.Printf("%T\n", age) // to print type of the variable
	// Output : main.my_type -> package.data type

	var born new_type
	born = 1990

	fmt.Println("Born : ", born)
	fmt.Printf("%T\n", born)

	// type conversion
	var born1 int

	born1 = int(born)

	fmt.Println("Born 2 : ", born1)
	fmt.Printf("%T\n", born1)

	// blink identifier
	web_data, _ := http.Get("www.google.com")

	fmt.Println(web_data)

	// accessing constants
	fmt.Println(Name)
	fmt.Println(Age)

}
