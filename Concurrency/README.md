## Concurrency

This is very useful when we want to run many threads ( go routine).
When one finishes execution or wait for io etc, other routine starts execution.


For doing this mainly sync package is used.

`var wg sync.WaitGroup`

IN main function or where we call those concurrently running function, we have to use `Add()` and `Wait()` functions in `WaitGroup`

```
func foo(){
    code goes here
    wg.Done() // passing I am done
}

func main(){
    wg.Add(1) // number of go routines

    /*  It waits till these routines to finish and then finishes the
        main function.
        When one function returns I am Done(), it decrement the counter.
    */
    go foo()

    wg.Wait()
}

```

### Mutex
Mutex are used as guard for variables access by one go routine.
Mutex also in `sync` package.

```
var mtx sync.Mutex

// when one function is used that shared variable

mtx.Lock()
variable used here

mtx.UnLock()

```

#### Important
If your code has many concurrent routines, you must check for race conditions using built in functionality.