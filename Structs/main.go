package main

import "fmt"

type person struct {
	First_Name string
	Last_Name  string
	Age        int
}

type programmar struct {
	Details   person // person struct
	Languages []string
}

func main() {
	p1 := person{
		First_Name: "Sachith",
		Last_Name:  "Muhandiram",
		Age:        29,
	}

	// prints full struct
	fmt.Println(p1) // {Sachith Muhandiram 29}

	// to access a value
	fmt.Println(p1.First_Name) // Sachith

	// programmar struct
	programmar1 := programmar{
		Details: person{
			First_Name: "Linus",
			Last_Name:  "Linux",
			Age:        45,
		},
		Languages: []string{"C", "C++"}, // you must use []string to insert data
	}

	fmt.Println(programmar1)
}
