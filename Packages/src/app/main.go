package main

import (
	"fmt"

	"../greeting"
)

/*
	Importing other path should be taken from where your source file located.
	Just need to link the package (Folder) not the file where our codes are located.
*/
func main() {
	fmt.Println(greeting.Morning)
}
