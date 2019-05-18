## Descrption

This explains how to define varibales in go lang

There is one important thing, you **MUST** use a defined varible, if not go compiler wont compile the code and give you an error.

>> variable declared and not used
### Important

variables declared with lowercase letters are private to the package. 
We will go to using diffent packages later.


### How to define a new data type
You can define your own data type and use that data type to define variable.

`type my_type int`

So `my_type` will be the new data type and its integer.

### How to convert data type

We have data type called `my_type`, to do that following syntax is used.

```
var age my_type = 29

var now_age int 
now_age = int(age) // here age which is my_type converted to int

```

#### Blink identifier
If you do not need to use a variable returning from a function (commonly used there), then you use blink identifier

```
web_data,_ := http.Get("www.google.com")

fmt.Println(web_data)

If we dont use _, we will have to use that varibale in our code.
http.Get(url) -> returns a response and an error (if there)
```

#### Constants
Constant are defined and values can not be changed.

`const my_name string = "Sachith Muhandiram"`

If we want to declare multiple constant for our program, then we use

```
const (
    Name = "Sachith Muhandiram"
    Age  = 29
)
```