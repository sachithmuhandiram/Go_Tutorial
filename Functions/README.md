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