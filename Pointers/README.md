## Pointers

This is same as pointers in C langauge.
A pointer only holds a memory location.

```
var age int
// the memory address of this varibale can be access by using & sing infront of it

fmt.Println(&age)
// This will print the memory address of age varibale

```

To access the data in a pointer, we use *varibale


### Usage

Simple usage of pointers is that is used for getting user inputs. (if we use fmt.Scan())

```
var age int

fmt.Println("Please enter your age :" )
fmt.Scan(&age)

fmt.Println("Your age is : ",age)

```

### Important 
array,slices,maps and channels are considered as refereces. So if we pass one of these to a function and modify it, it means we are modifying the memory location.