package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // un buffered integer channel

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Added value to channel")
			ch <- i // putting value to channel
		}
	}()

	go func() {
		for {
			fmt.Println("Getting value from channel ")
			fmt.Println(<-ch) // getting value from channel
		}
	}()

	time.Sleep(time.Second)
} // end of main
