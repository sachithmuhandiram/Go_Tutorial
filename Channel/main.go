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
	/*
		Here range works as a guard. It waits for ch channel to close.
		When channel is closed, it gets the data from that channel.

	*/
	c := make(chan int) // integer channel

	go sender(c)
	receiver(c) // no need to have a separate routine as its inside main

} // end of main

// send only
func sender(c chan<- int) { // sending to channel
	c <- 34
}

// receive only
func receiver(c <-chan int) { // receving from channel
	fmt.Println(<-c)
}
