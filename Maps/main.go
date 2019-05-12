package main

import "fmt"

func main() {

	// I will use only shorthand method
	my_names := map[string]string{"first": "sachith", "last": "muhandiram"}

	fmt.Println(my_names) // map[first:sachith last:muhandiram]

	// interate through a map
	for key, value := range my_names {
		fmt.Println(key, " : ", value)
	}

	//delete a value
	delete(my_names, "last")
	fmt.Println(my_names) // map[first:sachith]

}
