## Arrays and Slices

Slices are commonly used in go than arrays. Both have similarities and differences.
Here I will give a basic dscription about these two.

### Arrays

```
// var my_array[array_size] data_type
var my_array[5] int
my_array[0] = 1

// shorthand method
my_array := [size]data_type{values....}

```

### Slices

Here we do not need to define the size of the array. Thats the main visible difference between arrays and slices. (much more behind the scene)

But slice is built on top of an array. 

```
my_slice := []data_type{values...}
```

Slices, we can use a range of values. To do it, we can use following syntax.

`my_slice[start_index:end_index]`

Here new slice will have values with start_index and upto end_index (without it)

#### Appending a value to slice

`append` function can be used to add a new value to a slice.

`slice_name = append(slice_name,value)`

#### Appending two slices
To append two slices, go uses a special syntax

`myslice = append(myslice,myslice_other...)`

```
myslice := []string{"I ", "am "}
myother_slice := []string{"Sachith ","Muhandiram"}

myslice = append(myslice,myother_slice...)

fmt.Println(myslice)

```

#### Handling slices
There is a special syntax for dealing with slices, specially selecting values.
`myslice([start_index:end_index])`

Here slice selects from `start_index`, with that `start_index` index to `end_index` (not including it it)