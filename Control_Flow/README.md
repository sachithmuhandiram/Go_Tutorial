## Control flows

Here I will explain basic control flow loops used in go..

### IF loop
 If loop works same as other programming languages. Similarity between C language.

##### single if
 ```
 if condition {
     run this....
 }
```

##### if else
```
 if condition_1 {
     condition 1 run
 }else{
     not condition 1
 }
```

##### if else if
```
 if condition_1 {
     condition 1 run
 }else if condition_2{
     condition 2 run
 }else{
     no condition matched
 }
```

### FOR loop
 for loop similar to other langauge defination, except for syntax declaration

 ```
 for x:=0 ; x<10; x++{
     fmt.Println(x)
 }

 ```

we can use for loop for interating through an array/slice etc
for that we have to use range key word.

```
 numbers := []int{1,2,3,4,5,6,7,8,9} // integer slice

  for index,value := range numbers {
     fmt.Println(index "->" value)
 }
```
### switch statment

The only other control flow statment is `switch` in go. It doesnt have while loop etc.

```
switch condtion {
    case statement:
        do_something
    case statment2 :
        do_something_else
    default:
        do_nothing
}

This checks with the different parameters given in `case` statments and execute the matched one.
If none matched, then default condition is executed. 



