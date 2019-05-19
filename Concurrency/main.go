package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mtx sync.Mutex

func foo() {
	for i := 0; i < 50; i++ {
		fmt.Printf("Foo : %d\n", i)
	}
	wg.Done() // passing the message when this routine done
}

func bar() {
	for j := 0; j < 50; j++ {
		fmt.Printf("Bar : %d\n", j)
	}
	wg.Done()
}
func main() {

	wg.Add(2) // add 2 routine
	go foo()
	go bar()
	wg.Wait() // wait for those routine to finish
}
