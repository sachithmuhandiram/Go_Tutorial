package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname string
	Lname string
	Age   int
}

func main() {
	p1 := person{
		Fname: "Sachith",
		Lname: "Muhandiram",
		Age:   29,
	}

	p2 := person{
		Fname: "Linus",
		Lname: "Linux",
		Age:   45,
	}

	// now we have multiple instantces of person, if we need to export these
	// we have to use json.Marshal

	people := []person{p1, p2} // slice of persons

	fmt.Println(people)

	ppl, err := json.Marshal(people)

	if err != nil {
		fmt.Println("There is an error : ", err)
	}

	fmt.Println(string(ppl)) // remember to convert byte stream to string
}
