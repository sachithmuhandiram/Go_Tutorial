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