package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) // un buffered integer channel

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Added value to channel")
			ch <- i // putting value to channel
		}
		close(ch) // closing the channel after getting all data out
	}()

	for v := range ch {
		fmt.Println("Getting value from channel ")
		fmt.Println(v) // getting value from channel
	}

} // end of main
