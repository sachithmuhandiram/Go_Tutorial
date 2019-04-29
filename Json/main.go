package main

import (
	"encoding/json"
	"fmt"
)
/*

here I have updated struct with json tags, which are used to do
Unmarshal. Tags are necessory for mapping byte of data into data

*/
type person struct {
	Fname string `json:"Fname"`
	Lname string `json:"Lname"`
	Age   int `json:"Age"`
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
	// [{Sachith Muhandiram 29} {Linus Linux 45}]

	ppl, err := json.Marshal(people)

	if err != nil {
		fmt.Println("There is an error : ", err)
	}

	fmt.Println(string(ppl)) // remember to convert byte stream to string
	fmt.Println(ppl)
	/*
		The result byte stream
		[91 123 34 70 110 97 109 101 34 58 34 83 97 99 104 105 116 104 34 44 34 76 110 97 109 101 34 58 34 77 117 104 97 110 100 105 114 97 109 34 44 34 65 103 101 34 58 50 57 125 44 123 34 70 110 97 109 101 34 58 34 76 105 110 117 115 34 44 34 76 110 97 109 101 34 58 34 76 105 110 117 120 34 44 34 65 103 101 34 58 52 53 125 93]
	*/

	err = json.Unmarshal(ppl, &people) // byte stream and mem location of struct values
	
	if err != nil{
		fmt.Println("Error : ",err)
	}

	fmt.Println(people)
}
