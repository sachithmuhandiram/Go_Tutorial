## Channels

Channels are used when you need to insure data passing between two processes.
There are two types of channels, buffered and unbuffered.

`ch := make(chan int) ` // un buffered integer channel

`ch := make(chan int,2)` // buffered channel with two buffered values

When we put a value/s to a channel, it waits until that value is taken from the buffer.

```
//putting a value to buffer

ch <- 6  // adding 6 to the channel

   <- ch // removing value from channel 

