## Functions
Functions are one of the key thing in go or any other programming langauge.

Here we can use functions for doing many things. I will explain them one by one.

##### Basic function syntax
```
func functionName(variables...) return_type {

}
```

here return type and varibales can be empty.

#### Anonymus function
Its a function where you can define in your code without a function name. Simply as a variable.

```
my_anonymus := func adding(x int) int{
    return x+10
}

my_val := my_anonymus(1)

```
Here anonymus function needs to be called. If we want to self execute anonynus function. we can add () at the end.

```
func anonymus_self() int{
    fmt.Println("Hello, I am anonymus self")
}()
```

#### Variadic parameters

These are special cases where a function accepts zero or more parameters.
One useful case is dealing with slices, when an unknown slice is passed to a function we can use this.

```
func variatic(v ...int) int{
    total :=0
    for _,v := range v{
        total += v
    }

    return total
}
```

#### Variadic arguments
You can pass variable number of input to a function using this.

`num := my_function(x...)`

This is useful when we try to send differnt data type other than function expecting.

#### Function callbacks
A function can be passed to an another function as an parameter.

```
func my_function(x int, func my_func func(int)) {
    my_func(x)
}

```
#### init function
This function is used when we need to configure special parameters for the code.
This is the first thing to get executed (before main).

#### defer
`defer` keyword is used when we call a function, it makes sure that called function is the last function to call in that block.
This does opposite of `init` function.



