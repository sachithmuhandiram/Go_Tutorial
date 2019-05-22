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
```

Adding and getting data from the channel should happen simultaniously.
To achieve this, those two should run with different go routine or should use buffered channels.

Channels can be only sending and receiving. We can not do the opposite with that channel.

#### Working with multiple channels
If we have multiple channels, we can use `select{}` for getting data from different channels.

```
select {
    case data:= <- ch1 :
        fmt.Println("Channel 1")
    case data:= <- ch2 :
        fmt.Println("Channel 2")
    return
   }
```

##### Altanative way for select
I found [this is] intersting as an altanative way for `select` (https://medium.com/@pedram.esmaeeli/a-pattern-for-overcoming-non-determinism-of-golang-select-statement-139dbe93db98)



