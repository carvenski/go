## learn golang
```
go和python真的很像... 显然go的设计就是参考了C+Python...
对比学习:go可以促使对比思考python的很多东西,很多相通的编程概念/思路等,有很多收获...

[go教程英文原版](https://golangbot.com/learn-golang-series/)
[go教程中文版](http://blog.csdn.net/u011304970/article/details/76168257)
[不错的go教程](https://zengweigang.gitbooks.io/core-go/content/eBook/03.8.html)
[go评论](https://s.geekbang.org/search/c=0/k=/t=Go)
[go socket实现聊天](https://victoriest.gitbooks.io/golang-tcp-server/content/chapter2.html)
[国外牛人的go技术博客](https://www.goinggo.net/2014/01/concurrency-goroutines-and-gomaxprocs.html)
[深入go](https://tiancaiamao.gitbooks.io/go-internals/content/zh/01.1.html)
```



#### installion
	https://golang.org/dl/	
    note that must set $GOROOT/$GOPATH

#### golang features:
     Concurrency is an inherent part of the language. As a result writing multithreaded programs is a piece of cake. This is achieved by Goroutines and channels which we will discuss later in the upcoming tutorials.
     Golang is a compiled language. The source code is compiled to native binary.
     The language spec is pretty simple.
     The go compiler supports static linking. All the go code can be statically linked into one big fat binary and it can be deployed in cloud servers easily without worrying about dependencies.

#### Hello World 
```
//your project code structure is like this:
    GoProject/
             bin/  
	            hello
             src/
                hello/
                     helloworld.go

// helloworld.go
package main

import "fmt"

func main() {  
    fmt.Println("Hello World")
	}

// then use `go run/build helloworld.go` or `go install hello` to run it.
```

#### Variable
`var name type` is the syntax to declare a single variable.
`var name type = initialvalue` is the syntax to declare a variable with 

```go
 var age int
 age = 100
 
 var age int = 100
```

#### Type inference
If the variable is declared using the syntax var name = initialvalue, Go will automatically infer the type of that variable from the initial value.
```go
 var age = 100 // type will be inferred
```

#### Short hand declaration 
Go also provides another concise way for declaring variables. This is known as short hand declaration and it uses := operator.
name := initialvalue is the short hand syntax to declare a variable.
Short hand syntax can only be used when at least one of the variables in the left side of := is newly declared !
```go
name, age := "naveen", 24  // short hand declaration
age, height := 24, 170  // ok
name, height := "naveen", 170  // Error! no new variables
```

#### Types
     bool
     Numeric Types
          int8, int16, int32, int64,  int
          uint8, uint16, uint32, uint64,  uint
          float32, float64
          complex64, complex128
          byte
          rune
     string

```go
// bool
a := true
b := false
c := a && b
d := a || b

// Integer
// Signed integers
int8: represents 8 bit signed integers 
size: 8 bits 
range: -128 to 127

int16: represents 16 bit signed integers 
size: 16 bits 
range: -32768 to 32767

int32: represents 32 bit signed integers 
size: 32 bits 
range: -2147483648 to 2147483647

int64: represents 64 bit signed integers 
size: 64 bits 
range: -9223372036854775808 to 9223372036854775807

int: represents 32 or 64 bit integers depending on the underlying platform. You should generally be using int to represent integers unless there is a need to use a specific sized integer. 
size: 32 bits in 32 bit systems and 64 bit in 64 bit systems. 
range: -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems

var a int = 89
b := 95

// Unsigned integers
uint8: represents 8 bit unsigned integers 
size: 8 bits 
range: 0 to 255

uint16: represents 16 bit unsigned integers 
size: 16 bits 
range: 0 to 65535

uint32: represents 32 bit unsigned integers 
size: 32 bits 
range: 0 to 4294967295

uint64: represents 64 bit unsigned integers 
size: 64 bits 
range: 0 to 18446744073709551615

uint : represents 32 or 64 bit unsigned integers depending on the underlying platform. 
size : 32 bits in 32 bit systems and 64 bits in 64 bit systems. 
range : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems


// Floating point types
float32: 32 bit floating point numbers 
float64: 64 bit floating point numbers(float64 is the default type for floating point values)

a, b := 5.67, 8.97

// Other numeric types
byte is an alias of uint8 
rune is an alias of int32

// String type
first := "Naveen"
last := "Ramanathan"
name := first +" "+ last

// Type Conversion
Go is very strict about explicit typing. There is no automatic type promotion or conversion. Lets look at what this means with an example.
i := 55      //int
j := 67.8    //float64
sum := i + j //int + float64 not allowed ! you must convert Type yourself !
//To fix the error, both i and j should be of the same type. Let's convert j to int. T(v) is the syntax to convert a value v to type T.
i := 55      //int
j := 67.8    //float64
sum := i + int(j) //j is converted to int

//Constants 
const a = 55 //allowed
a = 89 // Error: const reassignment not allowed
//The value of a constant should be known at compile time. Hence it cannot be assigned to a value returned by a function call since the function call takes place at run time.
var a = math.Sqrt(4)  //allowed
const b = math.Sqrt(4) //Error: b is a constant and the value of b needs to be know at compile time. but function math.Sqrt(4) will be evaluated only during run time.
//String Constants
t hello = "Hello World"  
//Boolean Constants
const trueConst = true
//(Note that const has no type, its type is untyped ?!)
```

#### Functions
The general syntax for declaring a function in go is
```go
func fname(parametername type) returntype {  
 //function body
}

func calculateBill(price int, no int) int {  
    var totalPrice = price * no
    return totalPrice
}

//The parameters and return type are optional in a function. no parameters and no return value.
func fname() {  
 //function body
}
```

#### Multiple return values
```go
func rectProps(length, width float64)(float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}

area, perimeter := rectProps(10.8, 5.6)
```

#### Named return values
It is possible to return named values from a function. If a return value is named, it can be considered as being declared as a variable in the first line of the function.
```go
func rectProps(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return //no explicit return value
}
Since area and perimeter are specified in the function declaration as return values, they are automatically returned from the function when a return statement in encountered.
```

#### Blank Identifier
_ is know as the blank identifier in Go. It can be used in place of any value of any type. use _ to ignore a value.
```go
func rectProps(length, width float64) (float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}

func main() {  
    area, _ := rectProps(10.8, 5.6) // perimeter is discarded
    fmt.Printf("Area %f ", area)
}
```

#### Packages
Every executable go application must contain a main function. This function is the entry point for execution. The main function should reside in the main package.
Source files belonging to a package should be placed in separate folders of their own. It is a convention in Go to name this folder with the same name of the package.

Any variable or function which starts with a capital letter are exported in go. Only exported functions and variables can be accessed from other packages.
```go
So lets create a folder named rectangle inside the src/hello/ folder. 
All files inside the rectangle folder should start with the line package rectangle as they all belong to the rectangle package.

GoProject/
         bin/
            hello   
         src/
            hello/  #this is the app project name
                 main.go  #this is main package&function as app entry           
                 rectangle/     #this is rectangle package
                          rectangle.go

import (  
    "hello/rectangle"  //import your package
 )
```

#### init function of a package
```go
Every package can contain a init function. (sort of like Python package __init__.py file)
The init function should not have any return type and should not have any parameters. 
The init function cannot be called explicitly in our source code. The init function looks like below

Package level variables are initialised first, init function is called next. 
A package can have multiple init functions (either in a single file or distributed across multiple files) 
and they are called in the order in which they are presented to the compiler.

func init() {  
}

if there is a init function in a main package, The order of initialisation of the main package is

The imported packages are first initialised. Hence rectangle package is initialised first.
Package level variables rectLen and rectWidth are initialised next.
init function is called.
main function is called then.
```

#### Use of blank identifier when import package
It is illegal in Go to import a package and not to use it anywhere in the code. 
Sometimes we need to import a package just to make sure the initialisation takes place even though we do not need to use any function or variable from the package. 
For example, we might need to ensure that the init function of the rectangle package is called even though we do not use that package anywhere in our code. 
The _ blank identifier can be used in this case too as show below.

```go
import (   
     _ "hello/rectangle" 
 )

```

#### if else statement
```go
num := 99
if num <= 50 {
       fmt.Println("number is less than or equal to 50")
} else if num >= 51 && num <= 100 {
       fmt.Println("number is between 51 and 100")
} else {
       fmt.Println("number is greater than 100")
}

//Note: The else statement should start in the same line after the closing curly brace } of the if statement. If not the compiler will throw Error !!

```

#### Loops
for is the only loop available in Go. Go doesn't have while or do while loops which are present in other languages like C. Go use for but it includes while...
```go
// for loop:
for i := 1; i <= 10; i++ {
        fmt.Printf(" %d",i)   
}
//The variables declared in a for loop are only available within the scope of the loop. Hence i cannot be accessed outside the body for loop !

// while loop:
i := 0
for i <= 10 { 
        fmt.Printf("%d ", i)
        i += 2
}

// dead loop:
for {  

}

```

#### switch 
```go
func main() {  
    finger := 4
	switch finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")  
    default:  //default case
        fmt.Println("incorrect finger number")
	}
}

func main() {  
    letter := "i"
	switch letter {
    case "a", "e", "i", "o", "u": //multiple expressions in case
        fmt.Println("vowel")
    default:
        fmt.Println("not a vowel")
	}
}

func main() {  
    num := 75
    switch { // expression is omitted
    case num >= 0 && num <= 50:
        fmt.Println("num is greater than 0 and less than 50")
    case num >= 51 && num <= 100:
        fmt.Println("num is greater than 51 and less than 100")
    case num >= 101:
        fmt.Println("num is greater than 100")
    }
}

```

#### Arrays and Slices
```go
var a [3]int  //int array with length 3 , a = [0, 0, 0]
a[0] = 12     // array index starts at 0
a[1] = 78
a[2] = 50

a := [3]int{12, 78, 50} // short hand declaration to create array

a := [3]int{12}         // a = [12, 0, 0]

a := [...]int{12, 78, 50} // ... makes the compiler determine the length
```

#### Arrays are value types in Go !!! Different from List in Python.
```go
Arrays in Go are value types and not reference types !!!
This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable. 
If changes are made to the new variable, it will not be reflected in the original array.

func main() {  
    a := [...]string{"USA", "China", "India", "Germany", "France"}
    b := a  // a value copy of a is assigned to b
    b[0] = "Singapore"
    fmt.Println("a is ", a) 
    fmt.Println("b is ", b) 
    // a is [USA China India Germany France]  
    // b is [Singapore China India Germany France]  
}

//Similarly when arrays are passed to functions as parameters, they are passed by value and the original array in unchanged.
func changeLocal(num [5]int) {  
    num[0] = 55
}

func main() {  
    num := [...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num)
    changeLocal(num)  //num is passed by value
    fmt.Println("after passing to function ", num)
    // num is  [5 6 7 8 8] still.
}
```

#### Iterate arrays using range
```go
a := [...]float64{67.7, 89.8, 21, 78}
    for i := 0; i < len(a); i++ {  //looping from 0 to the length of the array
        fmt.Printf("%d th element of a is %.2f\n", i, a[i])
    }

//to iterate an array by using the range form of the for loop. range returns both the index and the value at that index.
a := [...]float64{67.7, 89.8, 21, 78}
    sum := float64(0)
    for i, v := range a {  //range returns both the index and value
        fmt.Printf("%d the element of a is %.2f\n", i, v)
        sum += v
    }
    fmt.Println("\nsum of all elements of a",sum)

for _, v := range a {  //or ignores index  

}

```
#### multidimensional arrays
```go
func printarray(a [3][2]string) {  
		for _, v1 := range a {
			for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
			}
        fmt.Printf("\n")    
		}
}

func main() {  
a := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"},  //this comma is necessary. The compiler will complain if you omit this comma 
    }
    printarray(a)
    var b [3][2]string
    b[0][0] = "apple"
    b[0][1] = "samsung"
    b[1][0] = "microsoft"
    b[1][1] = "google"
    b[2][0] = "AT&T"
    b[2][1] = "T-Mobile"
    fmt.Printf("\n")
    printarray(b)
}
```

#### Slices
slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are the just references to existing arrays.
A slice with elements of type T is represented by []T.
```go
//creates and array and returns a slice reference, the slice is just a reference to that array !!
c := []int{6, 7, 8} 

//creates a slice from an exist array
a := [5]int{76, 77, 78, 79, 80}
var b []int = a[1:4]

# modifying a slice
A slice does not own any data of its own. It is just a representation of the underlying array. Any modifications done to the slice will be reflected in the underlying array.
func main() {  
    arr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
    slice := arr[2:5]
    fmt.Println("array before",arr)
	for i := range slice {
        slice[i] += 1
	}
    fmt.Println("array after",arr) 
    // arr is [57 89 91 83 101 78 67 69 59] now
}

# length & capcity of a slice
the length is elements number in slice, but the capcity is the length starting from left index of slice of origin array.
A slice can be re-sliced upto its capacity. Anything beyond that will cause the program to throw a run time error.

fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
fruitslice := fruitarray[1:3]
fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice))  //length of is 2 and capacity is 6

func main() {  
    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
    fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
    fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice)) //After re-slicing length is 6 and capacity is 6  
}
```
#### creating a slice using make
```go
func make([]T, len, cap) []T can be used to create a slice by passing the type, length and capacity. 
The capacity parameter is optional and defaults to the length. The make function creates an array and returns a slice reference to it.

i := make([]int, 5, 5)  //i = [0 0 0 0 0]
```
#### Appending to a slice
```go
As we already know arrays are restricted to fixed length and their length cannot be increased. 
Slices are dynamic and new elements can be appended to the slice using append function !! (like Python List ??)

The definition of append function is func append(s []T, x ...T) []T.
x ...T in the function definition means that the function accepts variable number of arguments for the parameter x. (variable parameter function)

****One question might be bothering you though. If slices are backed by arrays and arrays themselves are of fixed length then how come a slice is of dynamic length ?
Well what happens under the hoods is, when new elements are appended to the slice, a new array is created. 
The elements of the existing array are copied to this new array and a new slice reference for this new array is returned. 
The capacity of the new slice is now twice that of the old slice. Pretty cool right :). The following program will make things clear.

func main() {  
    cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
    cars = append(cars, "Toyota")
    fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
}


//The zero value of a slice type is nil. A nil slice has length and capacity 0. It is possible to append values to a nil slice using the append function.
func main() {  
    var names []string //zero value of a slice is nil
	if names == nil {
        names = append(names, "John", "Sebastian", "Vinay")
    }
}
 
//It is also possible to append one slice to another using the ... operator. 
func main() {  
    veggies := []string{"potatoes","tomatoes","brinjal"}
    fruits := []string{"oranges","apples"}
    food := append(veggies, fruits...)
}
   
```


#### Arrays are value types in Go !!! Different from List in Python.  (array是值类型)
#### Slice are reference types in Go !!! just like List in Python.    (slice是引用类型)
```go
Slices can be thought of as being represented internally by a structure type. This is how it looks,

type slice struct {  
    Length        int
    Capacity      int
    ZerothElement *byte  // a pointer to the zeroth element of the array. 
}

****A slice contains the length, capacity and a pointer to the zeroth element of the array. 
When a slice is passed to a function, even though its passed by value, the pointer variable will refer to the same underlying array. 
Hence when a slice is passed to a function as parameter, changes made inside the function are visible outside the function too. 

func subtactOne(numbers []int) {  
    for i := range numbers {
        numbers[i] -= 2
    }
}

func main() {  
    nos := []int{8, 7, 6}
    fmt.Println("slice before function call", nos)
    subtactOne(nos)                               //function modifies the slice actually
    fmt.Println("slice after function call", nos) //modifications are visible outside,  nos =  [6 5 4] now.
}

```
#### Multidimensional slices
```go
Similar to arrays, slices can have multiple dimensions.

func main() {  
     pls := [][]string {
            {"C", "C++"},
            {"JavaScript"},
            {"Go", "Rust"},
     }
     for _, v1 := range pls {
         for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
         }
        fmt.Printf("\n")
     }
}
```

#### little Memory Optimisation for slice and array
```go
Slices hold a reference to the underlying array. As long as the slice is in memory, the array cannot be garbage collected.
This might be of concern when it comes to memory management. 
Lets assume that we have a very large array and we are interested in processing only a small part of it. 
Henceforth we create a slice from that array and start processing the slice. 
The important thing to be noted here is that the array will still be in memory since the slice references it.

One way to solve this problem is to use the copy function func copy(dst, src []T) int to make a copy of that slice. (copy() copy value not reference, so...)
This way we can use the new slice and the original array can be garbage collected.

func countries() []string {  
    countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
    neededCountries := countries[:len(countries)-2]
    countriesCpy := make([]string, len(neededCountries))
    copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
    return countriesCpy
}

func main() {  
    countriesNeeded := countries()
    fmt.Println(countriesNeeded)
}
```

#### Variadic Functions(A variadic function is a function that can accept variable number of arguments)
Please note that only the last parameter of a function is allowed to be variadic.
```go
func find(num int, nums ...int) {    //Within the function find, the type of nums is equivalent to []int i.e, an integer slice.
    fmt.Printf("type of nums is %T\n", nums)  
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}

func main() {  
    find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87)
}

//The way variadic functions work is by converting the variable number of arguments passed, to a new slice of the type of the variadic parameter.

the variable number of arguments to the find function are 89, 90, 95. The find function expects a variadic int argument. 
Hence these three arguments will be converted by the compiler to a slice of type int []int{89, 90, 95} and then it will be passed to the find function.

```
#### Passing a slice to a variadic function
like python, you can't just pass slice into a variadic function, because the type is wrong, 
but you can split a slice into elements and then pass them, just like * operator in python !!
```go
func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}

// Error: can't directly pass a slice to a variadic function.
func main() {  
    nums := []int{89, 90, 95}
    find(89, nums)  // Error: nums here is a slice, not int type.
}


//There is a syntactic sugar which can be used to pass a slice to a variadic function. 
//You have to suffix the slice with ... If that is done, the slice will be divided into its elements and pass each of them.
func main() {  
    nums := []int{89, 90, 95}
    find(89, nums...)  // OK
}

#Note that:
#when you pass nums... to find(), it passes reference, so the nums will be changed actually in find() if you modify it. 
#------------------------------------------------------------------------------------------------------------------------------
#example 1:
func change(s ...string) {  
    s[0] = "Go"
    fmt.Println(s)
}
func main() {  
    welcome := []string{"hello", "world"}
    change(welcome...)
    fmt.Println(welcome)  // welcome = [Go world]
}

#example 2:
func change(s ...string) {  
    s[0] = "Go"     // s = [Go world] now
    s = append(s, "playground")  // append() will create a new array in mem and copy origin array values, so left s points the new created array here, not welcome array !!   
    fmt.Println(s)  // s = [Go world playground] now 
}
func main() {  
    welcome := []string{"hello", "world"}
    change(welcome...)
    fmt.Println(welcome)  // welcome = [Go world] not [Go world playground] !!
}
#------------------------------------------------------------------------------------------------------------------------------
```

#### Map
```go
func main() {  
    var personSalary map[string]int  //The zero value of a map is nil. If you try to add items to nil map, a run time panic will occur. 
    if personSalary == nil {
        fmt.Println("map is nil. Going to make one.")
        personSalary = make(map[string]int)  //Hence the map must be initialized using make function.
        }
}

//so better way to use a Map is to use make(map[string]int) directly:
personSalary := make(map[string]int)
personSalary["michaely"] = 25000

haha := map[string]int {
      "steve": 12000,
      "jamie": 15000,
}

It's not necessary that only string types should be keys. All comparable types such as boolean, integer, float, complex, string, ... can also be keys. 
If you would like to know more about comparable types, please visit http://golang.org/ref/spec#Comparison_operators


//What will happen if a element is not present? The map will return the zero value of the type of that element.
func main() {  
personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
              }
    personSalary["mike"] = 9000
    employee := "jamie"
    fmt.Println("Salary of", employee, "is", personSalary[employee])
    fmt.Println("Salary of joe is", personSalary["joe"])  //Salary of joe is 0   
}

//What if we want to know whether a key is present in a map or not.
value, ok := map[key]  // i don't like this in Go...shit...

value, ok := personSalary[newEmp]
if ok == true {
        fmt.Println("Salary of", newEmp, "is", value)
} else {
        fmt.Println(newEmp,"not found")   
}

// iterate Map
for k, v := range personSalary {
        fmt.Printf("personSalary[%s] = %d\n", k, v)   
}

//delete(map, key) is the syntax to delete key from a map.
delete(personSalary, "steve")

//Maps can't be compared using the == operator. The == can be only used to check if a map is nil.
if map1 == map2  // Error
if map1 == nil   // OK
One way to check whether two maps are equal is to compare each one's individual elements one by one... 
```

#### Maps are reference types 
```
Similar to slices, maps are reference types. When a map is assigned to a new variable, they both point to the same internal data structure.
Hence changes made in one will reflect in the other.
Similar is the case when maps are passed as parameters to functions. When any change is made to the map inside the function, it will be visible to the caller.
```

#### Strings
```go
//A string in Go is a slice of bytes. (like python String)
//Since a string is a slice of bytes, its possible to access each byte of a string.

func main() {  
    name := "Hello World"
    haha := name[0:5]
    lett := name[-1]
}

// (string & byte & rune difference in go??)


//iterate over the individual runes of a string. 
s := "Señor"
for index, rune := range s {
    fmt.Printf("%c starts at byte %d\n", rune, index)
}

//Constructing string from slice of bytes.
func main() {      
    byteSlice := []byte{67, 97, 102, 195, 169}  //byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
    s := string(byteSlice)
    fmt.Println(s)
}
//Constructing a string from slice of runes.
func main() {  
    runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    s := string(runeSlice)
    fmt.Println(s)
}
```

#### Strings are immutable (just like python String)
```go
//Strings are immutable in Go. Once a string is created its not possible to change it.
//you just change the variable which point to different string object.

func main() {  
    h := "hello"
    h[0] = 'a'  // Error: string is immutable : cannot assign to h[0]
}

//there is a tricky way of string immutability: convert strings to a slice of runes. 
//Then that slice is mutated with whatever changes needed and converted back to a new string.
s := "hello"
runes := []rune(s)
runes[0] = 'a'      // OK, because runes is a slice not a string
s = string(runes)   // s point to a new string now
```

#### Pointer in go
*A pointer is a variable which stores the memory address of another variable.*
*the reference type is using Pointer in essence*
```go
关于指针的3点基础知识:

// 1.Declaring pointer:
b := 255
var a *int = &b        //在声明指针变量的时候, *是指示当前变量为指针类型 !!

// 2.Using pointer:
fmt.Println("address of b is", a)
fmt.Println("address of b is", *a)     //而在使用指针变量的时候, *是取值操作符 !!

// 3.&永远是取址操作符
```

```go
The zero value of a pointer is nil.
Lets write one more program where we change the value in b using the pointer.

func main() {  
    b := 255
    a := &b
    fmt.Println("address of b is", a)
    fmt.Println("value of b is", *a)
    *a++                               // it's equals to b++
    fmt.Println("new value of b is", b)
}
```

#### Passing pointer to a function
```go
func change(val *int) {               //声明 函数的参数是指针类型时,也使用 *T 声明
    *val = 55
}

func main() {  
    a := 58
    fmt.Println("value of a before function call is",a)
    b := &a
    change(b)
    fmt.Println("value of a after function call is", a)
}
```

#### Do not pass a pointer to an array as a argument to a function. Use slice instead !!
```go
//Not Recommended: use "pointer" to an array as argument to function and make modification inside and outside:
//  Although this way of passing a pointer to an array as a argument to a function and making modification to it works, 
//  it is not the idiomatic way of achieving this in Go. We have slices for this.
func modify(arr *[3]int) {  
    (*arr)[0] = 100
}
func main() {  
    a := [3]int{89, 90, 91}
    modify(&a)
    fmt.Println(a)     // a = [100, 90, 91]
}

//Recommended: use "slice" to an array as argument to function and make modification inside and outside:
func modify(sls []int) {  
    sls[0] = 90
}
func main() {  
    a := [3]int{89, 90, 91}
    modify(a[:])       // a[:] is a slice point to a array, actually it's already pointer in background.
    fmt.Println(a)     // a = [100, 90, 91]
}    
```

#### Go does not support pointer arithmetic like pointer in C.
```go
b := [...]int{109, 110, 111}
p := &b
p++     // Error: go pointer can't do arithmetic
```

----------------------------------------
  # OOP in Go, by Structure, no Class.
----------------------------------------

#### Structure (A structure is a user defined type which represents a collection of fields)
*in go, we use structure to realize Class & OOP staff*
```go
//Declaring and use a named structure:
type Employee struct {  
    firstName string
    lastName  string
    age       int
}
e1 := Employee{firstName: "Sam", lastName: "Anderson", age: 25,}
e2 := Employee{"Thomas", "Paul", 29, 800}                          //args in order

//Creating anonymous structures, no struct name:
e3 := struct {
    firstName, lastName string
    age, salary         int
}{
    firstName: "Andreah",
    lastName:  "Nikola",
    age:       31,
    salary:    5000,
}

//When a struct is defined and it is not explicitly initialised with any value, 
//the fields of the struct are assigned their zero values by default.
type Employee struct {  
    firstName, lastName string
    age, salary         int
}
var e4 Employee  // e4 = {  0 0}  
```

#### Accessing individual fields of a struct
*The . operator is used to access the individual fields of a structure, just like Class*
```go
type Employee struct {  
    firstName, lastName string
    age, salary         int
}
e6 := Employee{"Sam", "Anderson", 55, 6000}
fmt.Println("First Name:", e6.firstName)

var e7 Employee
e7.firstName = "Jack"
e7.lastName = "Adams"
fmt.Println("Employee 7:", e7)
```

#### Pointers to a struct
```go
e8 := &Employee{"Sam", "Anderson", 55, 6000}
fmt.Println("First Name:", (*e8).firstName)

****actually, when a variable point to a reference type, it is a pointer !! ****
    e8 := &Employee{"Sam", "Anderson", 55, 6000}
    (*e8).firstName
equals to:
    e8 := Employee{"Sam", "Anderson", 55, 6000}
    e8.firstName
```

#### Nested structs
It is possible that a struct contains a field which in turn is a struct. 
```go
type Address struct {  
    city, state string
}
type Person struct {  
    name string
    age int
    address Address
}
func main() {  
    var p Person
    p.name = "Naveen"
    p.age = 50
    p.address = Address {
        city: "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p.name)
    fmt.Println("Age:",p.age)
    fmt.Println("City:",p.address.city)
    fmt.Println("State:",p.address.state)
}
```

#### Promoted fields
Fields that belong to a anonymous struct field in a structure are called promoted fields.     
since they can be accessed as if they belong to the structure which holds the anonymous struct field.     
```go
type Address struct {  
    city, state string
}
type Person struct {  
    name string
    age  int
    Address
}
func main() {  
    var p Person
    p.name = "Naveen"
    p.age = 50
    p.Address = Address{
        city:  "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.city)      //city is promoted field
    fmt.Println("State:", p.state)    //state is promoted field
}
```

#### Exported Structs and Fields
If a struct type starts with a capital letter, then it is a exported type and it can be accessed from other packages.    
Similarly if the fields of a structure start with caps, they can be accessed from other packages.     
```go
type Spec struct {  //exported struct  
    Maker string    //exported field
    model string      //unexported field
    Price int       //exported field
}
```

#### Structs Equality
```go
//Structs are value types and are comparable if each of their fields are comparable.    
//Two struct variables are considered equal if their corresponding fields are equal.    
name1 := name{"Steve", "Jobs"}
name2 := name{"Steve", "Jobs"}
if name1 == name2 {
    fmt.Println("name1 and name2 are equal")
} else {
    fmt.Println("name1 and name2 are not equal")
}

name3 := name{firstName:"Steve", lastName:"Jobs"}
name4 := name{}
name4.firstName = "Steve"
if name3 == name4 {
    fmt.Println("name3 and name4 are equal")
} else {
    fmt.Println("name3 and name4 are not equal")
}

//But, Struct variables are not comparable if they contain fields which are not comparable.
type image struct {  
    data map[int]int
}
func main() {  
    image1 := image{data: map[int]int{
        0: 155,
    }}
    image2 := image{data: map[int]int{
        0: 155,
    }}
    if image1 == image2 {  // Error: struct containing map[int]int cannot be compared
        fmt.Println("image1 and image2 are equal")
    }
}
```

#### Methods (function associated with a Type, like Class Methods)
A method is just a function with a special receiver type that is written between the func keyword and the method name.     
The receiver can be either struct type or non struct type. The receiver is available for access inside the method.     
func (t Type) methodName(parameter list) {         
}       
```go
type Employee struct {  
    name     string
    salary   int
}
// displaySalary() method has Employee as the receiver type
func (e Employee) displaySalary() {  
    fmt.Printf("Salary of %s is %d", e.name, e.salary)
}
func main() {  
    emp1 := Employee {
        name:     "michaely",
        salary:   25000,
    }
    emp1.displaySalary() //Calling displaySalary() method of Employee type
}

**** Go is not a pure object oriented programming language and it does not support classes. 
**** Hence methods on types is a way to achieve behaviour similar to classes.
```

#### Pointer receivers vs value receivers of Methods
```go
type Employee struct {  
    name string
    age  int
}
// Method with value receiver  
func (e Employee) changeName(newName string) {  
    e.name = newName
}
// Method with pointer receiver  
func (e *Employee) changeAge(newAge int) {  
    e.age = newAge
}
func main() {  
    e := Employee{
        name: "Mark Andrew",
        age:  50,
    }
    fmt.Printf("Employee name before change: %s", e.name)      //Mark Andrew
    e.changeName("Michael johson") 
    fmt.Printf("Employee name after change: %s", e.name)       //Mark Andrew

    fmt.Printf("Employee age before change: %d", e.age)        //50
    (&e).changeAge(51)  // equals to e.changeAge(51)
    fmt.Printf("Employee age after change: %d", e.age)         //51

    e.changeAge(52)     // will be interpreted as (&e).changeAge(52) by the language.
    fmt.Printf("Employee age after change: %d", e.age)         //52
}
```

#### When to use pointer receiver and when to use value receiver
```
Generally pointer receivers can be used when changes made to the receiver inside the method should be visible to the caller.
Pointers receivers can also be used in places where its expensive to copy a data structure. 
Consider a struct which has many fields. Using this struct as a value receiver in a method will need the entire struct to be copied which will be expensive. 
In this case if a pointer receiver is used, the struct will not be copied and only a pointer to it will be used in the method.
```

#### Methods of anonymous fields
Methods belonging to anonymous fields of a struct can be called as if they belong to the structure where the anonymous field is defined.
```go
type address struct {  
    city  string
    state string
}
func (a address) fullAddress() {  
    fmt.Printf("Full address: %s, %s", a.city, a.state)
}
type person struct {  
    firstName string
    lastName  string
    address
}
func main() {  
    p := person{
        firstName: "Elon",
        lastName:  "Musk",
        address: address {
            city:  "Los Angeles",
            state: "California",
        },
    }
    p.fullAddress() //accessing fullAddress method of address struct
}
```

#### Value receivers in methods .VS. value arguments in functions
```go
type rectangle struct {  
    length int
    width  int
}
func area(r rectangle) {  
    fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}
func (r rectangle) area() {  
    fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}
func main() {  
    r := rectangle{
        length: 10,
        width:  5,
    }
    area(r)   
    r.area()

    p := &r
    area(p)   //compilation Error: cannot use p (type *rectangle) as type rectangle in argument to area  
    p.area()  //OK, equals to r.area()
}
```

#### Pointer receivers in methods .VS. pointer arguments in functions.
```go
type rectangle struct {  
    length int
    width  int
}
func perimeter(r *rectangle) {  
    fmt.Println("perimeter function output:", 2*(r.length+r.width))

}
func (r *rectangle) perimeter() {  
    fmt.Println("perimeter method output:", 2*(r.length+r.width))
}
func main() {  
    r := rectangle{
        length: 10,
        width:  5,
    }
    p := &r //pointer to r
    perimeter(p)
    p.perimeter()

    perimeter(r)    //Error: cannot use r (type rectangle) as type *rectangle in argument to perimeter
    r.perimeter()   //OK, equals to p.perimeter()
}
```

#### Methods on non-struct types
```
It is also possible to define methods on non struct types but there is a catch. 
To define a method on a type, 
the definition of the receiver type of the method and the definition of the method must be in the same package !!
```
```go
package main
func (a int) add(b int) {   // Error: int type definition and method definition not in same package.
}
func main() {
}
//This is not allowed since the definition of the method add and the definition of type int are not in the same package. 
//This program will throw compilation error: cannot define new methods on non-local type int

//right way to do it:
//Create a type alias for the built-in type int and then create a method with this type alias as the receiver.
type myInt int
func (a myInt) add(b myInt) myInt {  
    return a + b
}
func main() {  
    num1 := myInt(5)
    num2 := myInt(10)
    sum := num1.add(num2)
    fmt.Println("Sum is", sum)
}

```

#### Interface (In Go, an interface is a set of method signatures)
    in the OOP world, "interface defines the behaviour of an object". 
    It only specifies what the object is supposed to do. 
    The way of achieving this behaviour (implementation detail) is upto the object.
    
    When a type provides definition for all the methods in the interface, it is said to implement the interface. 
    Interface specifies what methods a type should have and the type decides how to implement these methods.
    go interfaces are implemented implicitly if a type contains all the methods declared in the interface.
```go
//interface definition
type VowelsFinder interface {  
    FindVowels() []rune
}

type MyString string

//MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {  
    var vowels []rune
    for _, rune := range ms {
        if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
            vowels = append(vowels, rune)
        }
    }
    return vowels
}

func main() {  
    name := MyString("Sam Anderson")
    var v VowelsFinder
    v = name    // OK since MyString has implemented VowelsFinder interface.
    fmt.Printf("Vowels are %c", v.FindVowels())
}
```

#### Usage of Interface
```go
type SalaryCalculator interface {  
    CalculateSalary() int
}
type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}
type Contract struct {  
    empId  int
    basicpay int
}

//salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}
func main() {  
    pemp1 := Permanent{1, 5000, 20}
    pemp2 := Permanent{2, 6000, 30}
    cemp1 := Contract{3, 3000}
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)
}
```

#### Usage of Empty Interface
An interface which has zero methods is called empty interface. It is represented as interface{}. 
Since the empty interface has zero methods, all types implement the empty interface.
```go
func describe(i interface{}) {    // now you can pass any type here !!
    fmt.Printf("Type = %T, value = %v\n", i, i)
}
func main() {  
    s := "Hello World"
    describe(s)
    i := 55
    describe(i)
    strt := struct {
        name string
    }{
        name: "Naveen R",
    }
    describe(strt)
}
```

#### Type Assertion
Type assertion is used extract the underlying value of the interface.
i.(T) is the syntax which is used to get the underlying value of interface i whose concrete type is T.
```go
func assert(i interface{}) {  
    s := i.(int)      // get the underlying int value from i which type must be int !
    fmt.Println(s)
}
func main() {  
    var s interface{} = 56
    assert(s)         // OK

    var x interface{} = "Steven Paul"
    assert(x)         // Error: panic: interface conversion: interface {} is string, not int.
}


// If the concrete type of i is not T then ok will be false and v will have the zero value of type T 
// and the program will not panic !
func assert(i interface{}) {  
    v, ok := i.(int)         // if you receive returned error value, no panic then !!!
    fmt.Println(v, ok)       
}
func main() {  
    var s interface{} = 56   // 56 true  
    assert(s)
    var i interface{} = "Steven Paul"
    assert(i)                // no panic,  0 false  
}

//-------------------------------------------------------------------------------------------------------------
//****if you reveive the error value from a function, Go thought that you will handle the exception youself ??
//****so Go won't throw the panic then, right ?? SHIT...
//-------------------------------------------------------------------------------------------------------------
```

#### Type Switch
    A type switch is used to the compare the concrete type of an interface against multiple types specified in various case statements.
    It is similar to switch case. The only difference being the cases specify types and not values as in normal switch.
```go
func findType(i interface{}) {  
    switch i.(type) {
    case string:
        fmt.Printf("I am a string and my value is %s", i.(string))
    case int:
        fmt.Printf("I am an int and my value is %d", i.(int))
    default:
        fmt.Printf("Unknown type")
    }
}
func main() {  
    findType("Naveen")
    findType(77)
    findType(89.98)
}

//can also use interface with type
type Describer interface {  
    Describe()
}
type Person struct {  
    name string
    age  int
}
func (p Person) Describe() {  
    fmt.Printf("%s is %d years old", p.name, p.age)
}
func findType(i interface{}) {  
    switch v := i.(type) {
    case Describer:
        v.Describe()
    default:
        fmt.Printf("unknown type")
    }
}
func main() {  
    findType("Naveen")
    p := Person{
        name: "Naveen R",
        age:  25,
    }
    findType(p)
}
```

*TODO:*
#### Implementing interfaces using pointer receivers vs value receivers  
###### (please review "Pointer receivers vs value receivers of Methods" first)
*think of the difference here between them ??*
```go
type Describer interface {  
    Describe()
}
type Person struct {  
    name string
    age  int
}
func (p Person) Describe() {                            //implemented using value receiver  
    fmt.Printf("%s is %d years old\n", p.name, p.age)
}
type Address struct {  
    state   string
    country string
}
func (a *Address) Describe() {                          //implemented using pointer receiver  
    fmt.Printf("State %s Country %s", a.state, a.country)
}
// ------------------------------------------------------------------------------------------------
func main() {  
    var d1 Describer         // if `var d1 Person` here, opposite result !!??
    p1 := Person{"Sam", 25}
    d1 = p1                  //
    d1.Describe()
    p2 := Person{"James", 32}
    d1 = &p2                 //
    d1.Describe()

    var d2 Describer         // if `var d1 Address` here, opposite result !!??
    a1 := Address{"Washington", "USA"}    
    d2 = a1                  //
    d2.Describe()
    a2 := Address{"shanghai", "CN"}    
    d2 = &a2                 //
    d2.Describe()
}

//Explantion: ????
//This is because, the Describer interface is implemented using a Address Pointer receiver and 
//we are trying to assign a1 which is a value type and it has not implemented the Describer interface. 
//This will definitely surprise you since we learnt earlier that 
//methods with pointer receivers will accept both pointer and value receivers.
//Then why is the code opposite here ?
//
//The reason is that it is legal to call a pointer-valued method on anything that is already a pointer or whose address can be taken.
//The concrete value stored in an interface is not addressable and hence 
//it is not possible for the compiler to automatically take the address of a1 and hence this code fails.
// ------------------------------------------------------------------------------------------------
```

#### implementing multiple interfaces(A type can implement more than one interface)
```go
type SalaryCalculator interface {  
    DisplaySalary()
}
type LeaveCalculator interface {  
    CalculateLeavesLeft() int
}
type Employee struct {  
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}
func (e Employee) DisplaySalary() {  
    fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}
func (e Employee) CalculateLeavesLeft() int {  
    return e.totalLeaves - e.leavesTaken
}
func main() {  
    e := Employee {
        firstName: "Naveen",
        lastName: "Ramanathan",
        basicPay: 5000,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
    }
    var s SalaryCalculator = e
    s.DisplaySalary()
    var l LeaveCalculator = e
    fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
}
```

#### Embedding interfaces
*Although go does not offer inheritance, it is possible to create a new interfaces by embedding other interfaces.*
```go
type SalaryCalculator interface {  
    DisplaySalary()
}
type LeaveCalculator interface {  
    CalculateLeavesLeft() int
}

//Any type is said to implement EmployeeOperations interface if 
//it provides method definitions for the methods present in both SalaryCalculator and LeaveCalculator interfaces.
type EmployeeOperations interface {  
    SalaryCalculator
    LeaveCalculator
}
type Employee struct {  
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}
func (e Employee) DisplaySalary() {  
    fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}
func (e Employee) CalculateLeavesLeft() int {  
    return e.totalLeaves - e.leavesTaken
}
func main() {  
    e := Employee {
        firstName: "Naveen",
        lastName: "Ramanathan",
        basicPay: 5000,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
    }
    var empOp EmployeeOperations = e
    empOp.DisplaySalary()
    fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
}
```

#### Zero value of Interface
The zero value of a interface is nil. 
A nil interface has both its underlying value and as well as concrete type as nil.


-----------------------------------------------------------------------
      Concurrency in Go, use goroutine & channel,     
      just like coroutine & queue in Python but more simple/powerful.
-----------------------------------------------------------------------
## Concurrency(concurrency != parallel)
#### Goroutine introduce
    Goroutines are functions or methods that run concurrently with other functions or methods.
    Goroutines can be thought of as light weight threads. 

    They are only a few kb in stack size and the stack can grow and shrink according to needs of the application 
    whereas in the case of threads the stack size has to be specified and is fixed.

    The Goroutines are multiplexed to fewer number of OS threads. 
    There might be only one thread in a program with thousands of Goroutines. 
    If any Goroutine in that thread blocks say waiting for user input, then another OS thread is created and 
    the remaining Goroutines are moved to the new OS thread. 
    All these are taken care by the runtime and we as programmers are abstracted from these intricate details and 
    are given a clean API to work with concurrency.

    Goroutines communicate using channels. 
    Channels by design prevent race conditions from happening when accessing shared memory using Goroutines. 
    Channels can be thought of as a pipe using which Goroutines communicate.

#### start a goroutine
```go
func hello() {  
    fmt.Println("Hello world goroutine")
}
func main() {  
    go hello()                    //Now the hello() function will run concurrently along with the main() function.
    fmt.Println("main function")  //The main function runs in its own Goroutine and its called the main Goroutine.
}

//When a new Goroutine is started, the goroutine call returns immediately. 
//Unlike functions, the control does not wait for the Goroutine to finish executing. 
//The control returns immediately to the next line of code after the Goroutine call and 
//any return values from the Goroutine are ignored.
//The main Goroutine should be running for any other Goroutines to run. 
//If the main Goroutine terminates then the program will be terminated and no other Goroutine will run.
// so, the hello Goroutine did not get a chance to run at all... 

//let's make main goroutine wait for other goroutines running.
func numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}
func alphabets() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}
func main() {  
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("main terminated")
}
```
***one awesome graph explantion of goroutine***
![goroutine](https://golangbot.com/content/images/2017/07/Goroutines-explained.png "*one awesome graph explantion of goroutine*")

#### Channel
    Channels can be thought as pipes using which Goroutines communicate. 
    Similar to how water flows from one end to another in a pipe, 
    data can be sent from one end and received from the another end using channels.
    *Each channel has a type associated with it. 
    This type is the type of data that the channel is allowed to transport. 
    No other type is allowed to be transported using the channel.
    The zero value of a channel is nil if you just declare it. (var a chan int)
    *nil channels are no use and hence the channel has to be defined using make() similar to maps and slices.
    so, use make() more in go! `a := make(chan int)` 

```go
a := make(chan int)
data := <- a    // read from channel a  
a <- data       // write to channel a  
```

#### Send and Receive are blocking of unbuffered channel(queue_length=0) by default
    When a data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel.
    Similarly when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.
    This property of channels is what helps Goroutines communicate effectively without the use of explicit locks or conditional variables that are quite common in other programming languages.
```go
func hello(done chan bool) {  
    fmt.Println("Hello world goroutine")
    done <- true                    // block here until someone read from done channel 
}
func main() {  
    done := make(chan bool)
    go hello(done)
    <-done                          // block here until someone write to done channel 
    fmt.Println("main function")
}
```
```go
//example to split a task into different goroutines to run concurrently:
func calcSquares(number int, squareop chan int) {  
    sum := 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit
        number /= 10
    }
    squareop <- sum
}
func calcCubes(number int, cubeop chan int) {  
    sum := 0 
    for number != 0 {
        digit := number % 10
        sum += digit * digit * digit
        number /= 10
    }
    cubeop <- sum
} 
func main() {  
    number := 589
    sqrch := make(chan int)
    cubech := make(chan int)
    go calcSquares(number, sqrch)
    go calcCubes(number, cubech)
    squares, cubes := <-sqrch, <-cubech    //can also use 1 channel but read twice to realize same effect here
    fmt.Println("Final output", squares + cubes)
}
```

#### Deadlock
```go
func main() {  
    ch := make(chan int)
    ch <- 5               //Error: only write no read 
}
```

#### Unidirectional channels
    All the channels we discussed so far are bidirectional channels, that is data can be both sent and received on them. 
    It is also possible to create unidirectional channels, that is channels that only send or receive data.
    It is possible to convert a bidirectional channel to a send only or receive only channel but not the vice versa.
```go
func sendData(sendch chan<- int) {   //sendch is a write-only channel
    sendch <- 10
}
func main() {  
    chnl := make(chan int)
    go sendData(chnl)                //OK, automatical convertion
    fmt.Println(<-chnl)
}
```

#### Closing channels and for range loops on channels
*Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.*
```go
//Receivers can use an additional variable while receiving data from the channel to check whether the channel has been closed.
//If ok is false it means that we are reading from a closed channel. The value read from a closed channel will be the zero value of the channel's type. 
v, ok := <- ch  

func producer(chnl chan int) {  
    for i := 0; i < 10; i++ {
        chnl <- i
    }
    close(chnl)     // close channel after all writing task is done.
}
func main() {  
    ch := make(chan int)
    go producer(ch)
    for {
        v, ok := <-ch
        if ok == false {    //will be notified that the channel is closed, task already done.
            break
        }
        fmt.Println("Received ", v, ok)
    }
}
```

#### iterate a channel
    The for range form of the for loop can be used to receive values from a channel until it is closed.
```go
func producer(chnl chan int) {  
    for i := 0; i < 10; i++ {
        chnl <- i
    }
    close(chnl)
}
func main() {  
    ch := make(chan int)
    go producer(ch)
    for v := range ch {            // best way to read from channel until closed.
        fmt.Println("Received ",v)
    }
}
```
```go
//use multi goroutines to cooperate with each other example:
func digits(number int, dchnl chan int) {  
    for number != 0 {
        digit := number % 10
        dchnl <- digit
        number /= 10
    }
    close(dchnl)
}
func calcSquares(number int, squareop chan int) {  
    sum := 0
    dch := make(chan int)
    go digits(number, dch)
    for digit := range dch {
        sum += digit * digit
    }
    squareop <- sum
}
func calcCubes(number int, cubeop chan int) {  
    sum := 0
    dch := make(chan int)
    go digits(number, dch)
    for digit := range dch {
        sum += digit * digit * digit
    }
    cubeop <- sum
}
func main() {  
    number := 589
    sqrch := make(chan int)
    cubech := make(chan int)
    go calcSquares(number, sqrch)
    go calcCubes(number, cubech)
    squares, cubes := <-sqrch, <-cubech
    fmt.Println("Final output", squares+cubes)
}
```

#### Buffered Channel(queue_length > 0), and use buffered channel to realize a Worker Pool
    *send/receive to an unbuffered channel are blocking.*
    *send/receive to an buffered channel are only blocking when channel is full/empty.*

    It is possible to create a channel with a buffer.(just like a queue has a buffer) 
    Sends to a buffered channel are blocked only when the buffer is full. 
    Receives from a buffered channel are blocked only when the buffer is empty.
    ch := make(chan type, capacity)  //capacity > 0

```go
//example 1:
func main() {  
    ch := make(chan string, 2)
    ch <- "naveen"      
    ch <- "paul"
    fmt.Println(<- ch)
    fmt.Println(<- ch)
}
```
```go
//example 2:
func write(ch chan int) {  
    for i := 0; i < 5; i++ {
        ch <- i
        fmt.Println("successfully wrote", i, "to ch")
    }
    close(ch)
}
func main() {  
    ch := make(chan int, 2)
    go write(ch)
    time.Sleep(2 * time.Second)
    for v := range ch {
        fmt.Println("read value", v,"from ch")
        time.Sleep(2 * time.Second)

    }
}
```
```go
//example 3:
func main() {  
    ch := make(chan string, 2)
    ch <- "naveen"
    ch <- "paul"
    ch <- "steve"          //deadlock here...
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

#### Length & Capacity of a buffered channel
```go
//The capacity of a buffered channel is the number of values that the channel can hold. This is the value we specify when creating the buffered channel using the make() function.
//The length of the buffered channel is the number of elements currently queued in it.
func main() {  
    ch := make(chan string, 3)
    ch <- "naveen"
    ch <- "paul"
    fmt.Println("capacity is", cap(ch))
    fmt.Println("length is", len(ch))
    fmt.Println("read value", <-ch)
    fmt.Println("new length is", len(ch))
}
```

#### WaitGroup
    A WaitGroup is used to wait for a collection of Goroutines to finish executing. 
    The control is blocked until all Goroutines finish executing. 
    Lets say we have 3 concurrently executing Goroutines spawned from the main Goroutine. 
    The main Goroutines needs to wait for the 3 other Goroutines to finish before terminating. 
    This can be accomplished using WaitGroup.
    WaitGroup is a struct type actually. The way WaitGroup works is by using a counter. 
    (//structure is value or reference type when passing it ??)
    When we call Add on the WaitGroup and pass it an int, the WaitGroup's counter is incremented by the value passed to Add.
    The way to decrement the counter is by calling Done() method on the WaitGroup. 
    The Wait() methods blocks the Goroutine in which its called until the counter becomes zero.
```go
func process(i int, wg *sync.WaitGroup) {  
    fmt.Println("started Goroutine ", i)
    time.Sleep(2 * time.Second)
    fmt.Printf("Goroutine %d ended", i)
    wg.Done()
}
func main() {  
    no := 3
    var wg sync.WaitGroup
    for i := 0; i < no; i++ {
        wg.Add(1)
        go process(i, &wg)     //must use pointer here to change wg inside other function !! //structure is value or reference type when passing it??
    }
    wg.Wait()                  // block here until all goroutines in waitgroup finished.
    fmt.Println("All go routines finished executing")
}
```

### Worker Pool Implementation, by goroutine & buffered channel
**One of the important uses of buffered channel is the implementation of worker pool.**
    In general, a worker pool is a collection of threads which are waiting for tasks to be assigned to them.    
    Once they finish the task assigned, they make themselves available again for the next task.     
    more goroutines in worker pool, less time taken to finish jobs.
```go
package main
import (  
    "fmt"
    "math/rand"
    "sync"
    "time"
)
type Job struct {  
    id       int
    randomno int
}
type Result struct {  
    job         Job
    sumofdigits int
}
var jobs = make(chan Job, 10)  
var results = make(chan Result, 10)
func digits(number int) int {  
    sum := 0
    no := number
    for no != 0 {
        digit := no % 10
        sum += digit
        no /= 10
    }
    time.Sleep(2 * time.Second)
    return sum
}
func worker(wg *sync.WaitGroup) {  
    for job := range jobs {
        output := Result{job, digits(job.randomno)}
        results <- output
    }
    wg.Done()
}
func createWorkerPool(noOfWorkers int) {  
    var wg sync.WaitGroup
    for i := 0; i < noOfWorkers; i++ {
        wg.Add(1)
        go worker(&wg)
    }
    wg.Wait()
    close(results)
}
func allocate(noOfJobs int) {  
    for i := 0; i < noOfJobs; i++ {
        randomno := rand.Intn(999)
        job := Job{i, randomno}
        jobs <- job
    }
    close(jobs)
}
func result(done chan bool) {  
    for result := range results {
        fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
    }
    done <- true
}
func main() {  
    startTime := time.Now()
    noOfJobs := 100
    go allocate(noOfJobs)
    done := make(chan bool)
    go result(done)
    noOfWorkers := 10   //as the number of worker Goroutines increase, the total time taken to complete the jobs decreases.
    createWorkerPool(noOfWorkers)
    <-done
    endTime := time.Now()
    diff := endTime.Sub(startTime)
    fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
```
#### select
    The select statement is used to choose from multiple send/receive channel operations. 
    The select statement blocks until one of the send/receive operation is ready. 
    If multiple operations are ready, one of them is chosen at random. 
```go
func server1(ch chan string) {  
    time.Sleep(6 * time.Second)
    ch <- "from server1"
}
func server2(ch chan string) {  
    time.Sleep(3 * time.Second)
    ch <- "from server2"

}
func main() {  
    output1 := make(chan string)
    output2 := make(chan string)
    go server1(output1)
    go server2(output2)
    select {
    case s1 := <-output1:
        fmt.Println(s1)
    case s2 := <-output2:
        fmt.Println(s2)
    }
}
```
```go
func process(ch chan string) {  
    time.Sleep(10500 * time.Millisecond)
    ch <- "process successful"
}
func main() {  
    ch := make(chan string)
    go process(ch)
    for {
        time.Sleep(1000 * time.Millisecond)
        select {
        case v := <-ch:
            fmt.Println("received value: ", v)
            return
        default:  // The default case in a select statement is executed when none of the other case is ready.  
            fmt.Println("no value received")
        }
    }

}
```

#### race condition and Mutex(使用互斥锁避免竞态条件)
*learn how to solve race conditions using mutexes and channels.*
*please think about race conditions in multi threads/processes*

##### Critical section(临界区的概念 ?)
    Before jumping to mutex, it is important to understand the concept of critical section in concurrent programming. 
    when a program runs concurrently, the parts of code which modify shared resources should not be accessed by multiple Goroutines at the same time.
    This section of code which modifies shared resources is called critical section. 
    For example lets assume that we have some piece of code which increments a variable x by 1.
    x = x + 1 
    Internally the above line of code will be executed by the system in the following steps:
    (there are more technical details involving registers, how addition works and so on but for simplicity lets assume that these are the three steps),
        *step1: get the current value of x
        *step2: compute x + 1
        *step3: assign the computed value in step 2 to x
    you must realize that, the `x = x + 1` is not finished at 1 step immediately in deeper level... !!! 
    When these 3 steps are carried out by only one Goroutine, all is well.
    But when 2 Goroutines run this code concurrently, what happens ??

***2 awesome graph explantion of race condition***

![goroutine](https://golangbot.com/content/images/2017/08/cs5.png "*one awesome graph explantion of race condition*")

1.We have assumed the initial value of x to be 0. 
Goroutine 1 gets the initial value of x, computes x + 1 and before it could assign the computed value to x, the system context switches to Goroutine 2. 
Now Goroutine 2 gets the initial value of x which is still 0, computes x + 1. After this the system context switches again to Goroutine 1. 
Now Goroutine 1 assigns it's computed value 1 to x and hence x becomes 1. 
Then Goroutine 2 starts execution again and then assigns it's computed value, which is again 1 to x and hence x is 1 after both Goroutines execute.

![goroutine](https://golangbot.com/content/images/2017/08/cs-6.png "*one awesome graph explantion of race condition*")

2.a different scenario of what could happen: 
In the above scenario, Goroutine 1 starts execution and finishes all its three steps and hence the value of x becomes 1. 
Then Goroutine 2 starts execution. Now the value of x is 1 and when Goroutine 2 finishes execution, the value of x is 2.

#### race condition and Mutex introduce
    the two cases above you can see that the final value of x is 1 or 2 depending on how context switching happens. 
    This type of undesirable situation where the output of the program depends on the sequence of execution of Goroutines is called race condition.
    A Mutex is used to provide a locking mechanism to ensure that: 
    only one Goroutine is running the critical section of code at any point of time to prevent race condition from happening.
```go
//Mutex is available in the sync package. There are two methods defined on Mutex namely Lock and Unlock. 
//Any code that is present between a call to Lock and Unlock will be executed by only one Goroutine, thus avoiding race condition.
mutex.Lock()  
x = x + 1      //x = x + 1 now will be executed by only one Goroutine at any time thus preventing race condition.
mutex.Unlock()
//If one Goroutine already holds the lock and if a new Goroutine is trying to acquire a lock, the new Goroutine will be blocked until the mutex is unlocked.
```
```go
//example 1: no mutex, race condition happens.
var x  = 0  
func increment(wg *sync.WaitGroup) {  
    x = x + 1
    wg.Done()
}
func main() {  
    var w sync.WaitGroup
    for i := 0; i < 1000; i++ {
        w.Add(1)        
        go increment(&w)
    }
    w.Wait()
    fmt.Println("final value of x", x)    // x value will be different each time...
}

//example 2: use mutex, no race condition happens.
var x  = 0  
func increment(wg *sync.WaitGroup, m *sync.Mutex) {  
    m.Lock()
    x = x + 1
    m.Unlock()
    wg.Done()   
}
func main() {  
    var w sync.WaitGroup
    var m sync.Mutex
    for i := 0; i < 1000; i++ {
        w.Add(1)        
        go increment(&w, &m) //must pass the pointer of the mutex. 
    }                        //If passed value not pointer, each Goroutine will have its own copy of the mutex and the race condition will still occur!
    w.Wait()
    fmt.Println("final value of x", x)   // x = 1000, always
}

//example 3: Solving the race condition using channel.
var x  = 0  
func increment(wg *sync.WaitGroup, ch chan bool) {  
    ch <- true    // 
    x = x + 1     // ch here act just like a Lock. the 3 steps definitly run in 1 goroutine. awesome !
    <- ch         //
    wg.Done()   
}
func main() {  
    var w sync.WaitGroup
    ch := make(chan bool, 1)       
    for i := 0; i < 1000; i++ {
        w.Add(1)        
        go increment(&w, ch)
    }
    w.Wait()
    fmt.Println("final value of x", x)
}

//In general use channels when Goroutines need to communicate with each other and 
//use mutexes when only one Goroutine should access the critical section of code.
```


-----------------------------------------------------------------------
      Object Oriented Programming in Go, write code in OOP more
-----------------------------------------------------------------------

#### Structs Instead of Classes
```go
// just like Class:
type Employee struct {  
    FirstName   string
    LastName    string
    TotalLeaves int
    LeavesTaken int
 }
func (e Employee) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}

func main() {  
e := employee.Employee {
        FirstName: "Sam",
        LastName: "Adolf",
        TotalLeaves: 30,
        LeavesTaken: 20,
    
   }
    e.LeavesRemaining()
}
```

#### NewT() function to act like Constructor
```
    Go doesn't support constructors. 
    If the zero value of a type is not usable, it is the job of the programmer to unexport the type to prevent access from other packages and 
    also to provide a function named NewT(parameters) which initialises the type T with the required values.(you can write a NewEmployee() function here)
    It is a convention in Go to name a function which creates a value of type T to NewT(parameters). This will act like a constructor.(it's totally controlled by your code)
    If the package defines only one type, then its a convention in Go to name this function just New(parameters) instead of NewT(parameters).
```
```go
//employee.go
type employee struct {  
    firstName   string
    lastName    string
    totalLeaves int
    leavesTaken int
}
func NewEmployee(firstName string, lastName string, totalLeave int, leavesTaken int) employee {  
    e := employee {firstName, lastName, totalLeave, leavesTaken}
    return e
}
func (e employee) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}

//main.go
import "oop/employee"
func main() {  
    e := employee.NewEmployee("Sam", "Adolf", 30, 20)   
    e.LeavesRemaining()
}
```

#### Composition Instead of Inheritance
**Go does not support inheritance, however it does support composition.**    
**Composition can be achieved in Go is by embedding one struct type into another.**    
```go
//Composition can be achieved in Go is by embedding one struct type into another. 
type author struct {  
    firstName string
    lastName  string
    bio       string
}
func (a author) fullName() string {  
    return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {  
    title     string
    content   string
    author
}
func (p post) details() {  
    fmt.Println("Title: ", p.title)
    fmt.Println("Content: ", p.content)
    fmt.Println("Author: ", p.author.fullName())  //or: fmt.Println("Author: ", p.fullName()) 
    fmt.Println("Bio: ", p.author.bio)            //or: fmt.Println("Bio: ", p.bio)

//post struct has access to all the fields and methods of the author struct.
//Whenever one struct field is embedded in another, Go gives us the option to access the embedded fields as if they were part of the outer struct.
//This means that p.author.bio can be replaced with p.bio, p.author.fullName() can be replaced with p.fullName().(just like inheritance of Class)
```

#### embedding slice of structs
```go
type author struct {  
    firstName string
    lastName  string
    bio       string
}
func (a author) fullName() string {  
    return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}
type post struct {  
    title   string
    content string
    author
}
func (p post) details() {  
    fmt.Println("Title: ", p.title)
    fmt.Println("Content: ", p.content)
    fmt.Println("Author: ", p.fullName())
    fmt.Println("Bio: ", p.bio)
}

type website struct {  
  posts []post
}
func (w website) contents() {  
    fmt.Println("Contents of Website\n")
        for _, v := range w.posts {
        v.details()
        fmt.Println()
        }
 }
func main() {  
  author1 := author{
        "Naveen",
        "Ramanathan",
        "Golang Enthusiast",
         }
  post1 := post{
        "Inheritance in Go",
        "Go supports composition instead of inheritance",
        author1,
       }
  post2 := post{
        "Struct instead of Classes in Go",
        "Go does not support classes but methods can be added to structs",
        author1,
       }
  w := website{
        posts: []post{post1, post2, post3},
   }
    w.contents()
}
```

#### Polymorphism(所谓多态: 定义接口,接口有不同实现的type,写的时候用接口,真正用的时候却传入不同type)
#### (OOP的3大特性: 封装/继承/多态)
**A variable of type interface can hold any value which implements the interface. This property of interfaces is used to achieve polymorphism in Go.**
```go
type Income interface {  
    calculate() int
    source() string
}
type FixedBilling struct {  
    projectName string
    biddedAmount int
}
type TimeAndMaterial struct {  
    projectName string
    noOfHours  int
    hourlyRate int
}
func (fb FixedBilling) calculate() int {  
    return fb.biddedAmount
}
func (fb FixedBilling) source() string {  
    return fb.projectName
}
func (tm TimeAndMaterial) calculate() int {  
    return tm.noOfHours * tm.hourlyRate
}
func (tm TimeAndMaterial) source() string {  
    return tm.projectName
}
func calculateNetIncome(ic []Income) {  
    var netincome int = 0
        for _, income := range ic {
        fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
        netincome += income.calculate()
        }
    fmt.Printf("Net income of organisation = $%d", netincome)
}
func main() {  
    project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
    project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
    project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
    incomeStreams := []Income{project1, project2, project3}
    calculateNetIncome(incomeStreams)
}
```


-----------------------------------------------------------------------
    Go has no try-except-finally, but defer/panic/recover, SHIT...
-----------------------------------------------------------------------
#### Defer and Error Handling
**Defer statement is used to execute a function call just before the function where the defer statement is present returns.**
```go
func finished() {  
    fmt.Println("defer here in last")
}
func main() {  
    fmt.Println("main start 1")
    defer finished()
    fmt.Println("main start 2")
}
```
***The arguments of a deferred function are evaluated when the defer statement is executed and not when the actual function call is done.***
```go
func printA(a int) {  
    fmt.Println("value of a in deferred function", a)
}
func main() {  
    a := 5
    defer printA(a)  // a = 5 when defer shows up here, so when printA(a) really executed in last, a = 5, not a = 10 !! so you konw that defer has saved the value.
    a = 10
    fmt.Println("value of a before deferred function call", a)
}
```

#### Stack of defers(order of multi defer: the first the last/LIFO)
**When a function has multiple defer calls, they are added on to a stack and executed in Last In First Out (LIFO) order.**
```go
func main() {  
    name := "HELLO"
    for _, v := range []rune(name) {
        defer fmt.Printf("%c", v)      // "OLLEH"
    }
}
```

#### usage example of defer(defer act just like finally of try-except-finally)
```go
//1. not defer:
type rect struct {  
    length int
    width  int
 }
func (r rect) area(wg *sync.WaitGroup) {  
    if r.length < 0 {
        fmt.Printf("rect %v's length should be greater than zero\n", r)
        wg.Done()
        return
    }
    if r.width < 0 {
        fmt.Printf("rect %v's width should be greater than zero\n", r)
        wg.Done()
        return
    }
    area := r.length * r.width
    fmt.Printf("rect %v's area %d\n", r, area)
    wg.Done()
 }
func main() {  
    var wg sync.WaitGroup
    r1 := rect{-67, 89}
    r2 := rect{5, -67}
    r3 := rect{8, 9}
    rects := []rect{r1, r2, r3}
           for _, v := range rects {
               wg.Add(1)
               go v.area(&wg)
           }
    wg.Wait()
    fmt.Println("All go routines finished executing")
 }

//2. use defer in area():
func (r rect) area(wg *sync.WaitGroup) {  
    defer wg.Done()                             // code is more simple
    if r.length < 0 {
        fmt.Printf("rect %v's length should be greater than zero\n", r)
        return
        }
    if r.width < 0 {
        fmt.Printf("rect %v's width should be greater than zero\n", r)
        return   
    }
    area := r.length * r.width
    fmt.Printf("rect %v's area %d\n", r, area)
 }
```

## Error Handling
```
Errors are represented using the built-in error type.
If a function or method returns an error, then by convention it has to be the last value returned from the function.

The idiomatic way of handling error in Go is to compare the returned error to nil !!
A nil value indicates that no error has occurred and a non nil value indicates the presence of an error. 
```

```go
func main() {  
    f, err := os.Open("/test.txt")
        if err != nil {                  //compare err with nil in go, SHIT...
            fmt.Println(err)
            return
        }
    fmt.Println(f.Name(), "opened successfully")
 }
```

####  built-in Error type
**error is an interface type with the following definition, so you can simply write your own error with an Error() function**
```go
type error interface {  
    Error() string
 }
```

#### get more error information from your error type by adding more fields
```go
//if you custome your error type like this, you can get more clear error info about the error:
type PathError struct {  
    Op   string
    Path string
    Err  error
}
func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error()  }  
```

#### get more error information from your error type by adding more methods
```go
//if you custome your error type like this, you can get more clear error info about the error:
type DNSError struct {  
    ...
}
func (e *DNSError) Error() string {  
    ...
}
func (e *DNSError) Timeout() bool {  
    ... 
}
func (e *DNSError) Temporary() bool {  
    ... 
}
```

#### Directly compare error
```go
func main() {  
    files, error := filepath.Glob("[")
        if error != nil && error == filepath.ErrBadPattern {
        fmt.Println(error)
        return
        }
    fmt.Println("matched files", files)
}"]") }
```

## NOTE: Do not ignore error, Never ever ignore an error in Go !!!!

#### create our own custom errors which we can use in functions we create
```go
//Creating custom errors using the errors.New() function:
func circleArea(radius float64) (float64, error) {  
    if radius < 0 {
        return 0, errors.New("Area calculation failed, radius is less than zero")   //so others who use this function need to check & handle this error 
    }
    return math.Pi * radius * radius, nil
 }
func main() {  
    radius := -20.0
    area, err := circleArea(radius)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Area of circle %0.2f", area)
 }


//Providing more info about the error using your own customized XXError struct type:
//you just need to write a struct implement the error interface with an Error() function, then add your own fileds/methods.
type areaError struct {  
    err    string
    radius float64
 }
func (e *areaError) Error() string {  
    return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
 }
func circleArea(radius float64) (float64, error) {  
    if radius < 0 {
        return 0, &areaError{"radius is negative", radius}
    }
    return math.Pi * radius * radius, nil
 }
func main() {  
    radius := -20.0
    area, err := circleArea(radius)
    if err != nil {
        if err, ok := err.(*areaError); ok {
            fmt.Printf("Radius %0.2f is less than zero", err.radius)
            return
        }
        fmt.Println(err)
        return
    }
    fmt.Printf("Area of rectangle1 %0.2f", area)
}
```

#### Panic and Recover(panic-recover-defer just like try-except-finally)
    The idiomatic way to handle abnormal conditions in a program in Go is using errors. Errors are sufficient for most of the abnormal conditions arising in the program.
    But there are some situations where the program cannot simply continue executing after an abnormal situation. In this case we use panic to terminate the program. 
                      
    When a function encounters a panic, it's execution is stopped, any deferred functions are executed and then the control returns to it's caller.
    This process continues until all functions of the current goroutine have returned till the top caller. then print the stack trace  message and terminates. 

    you can regain control of a panicking program using recover.
    panic-recover is similar to try-except-finally in python except that it's rarely used and more elegant and results in clean code. SHIT...

    so When should we use panic but not return/check errors?
    Only in cases where the program just cannot continue execution should a panic and recover mechanism be used.(程序出现严重的异常而不能继续运行时直接panic退出程序)

```go
//The argument whatever passed to panic() will be printed when the program terminates.
func panic(interface{})   
```
```go
func fullName(firstName *string, lastName *string) {  
    if lastName == nil {
        panic("runtime error: last name cannot be nil")       //just like raise an exception in python...
    }
    fmt.Printf("%s %s\n", *firstName, *lastName)
 }
func main() {  
    firstName := "Elon"
    fullName(&firstName, nil)                     //the program will terminate/exit here because of a panic
    fmt.Println("returned normally from main")
 }
```

#### defer executed before panic
actually what panic() do is: return the panic to upper caller until main function then quit.       
so, you can easily know that defer will be executed before panic() in a function.
```go
func fullName(firstName *string, lastName *string) {  
    defer fmt.Println("deferred call in fullName")
    if lastName == nil {
        panic("runtime error: last name cannot be nil")
    }
    fmt.Printf("%s %s\n", *firstName, *lastName)
}
func main() {  
    defer fmt.Println("deferred call in main")
    firstName := "Elon"
    fullName(&firstName, nil)
    fmt.Println("returned normally from main")
}

//output:
//  deferred call in fullName  
//  deferred call in main  
//  panic: runtime error: last name cannot be nil
//  
//  goroutine 1 [running]:  
//  main.fullName(0x1042bf90, 0x0)  
//      /tmp/sandbox060731990/main.go:13 +0x280
//  main.main()  
//      /tmp/sandbox060731990/main.go:22 +0xc0
```

#### Recover
    Recover is useful only when called inside deferred functions. 
    call recover inside a deferred function stops the panicking sequence by restoring normal execution and retrieves the error value passed to panic.
    If recover is called outside the deferred function, it will not stop a panicking sequence.
    
```go
func recoverName() {  
    if r := recover(); r!= nil {                        // just like catch an exception
        fmt.Println("recovered from ", r)            
    }
 }
func fullName(firstName *string, lastName *string) {  
    defer recoverName()
    if lastName == nil {
        panic("runtime error: last name cannot be nil")
    }
    fmt.Printf("%s %s\n", *firstName, *lastName)
    fmt.Println("returned normally from fullName")
 }
func main() {  
    defer fmt.Println("deferred call in main")
    firstName := "Elon"
    fullName(&firstName, nil)
    fmt.Println("returned normally from main")
}
```

#### Panic, Recover and Goroutines
    Recover works only when it is called from the same goroutine. 
    It's not possible to recover from a panic that has happened in a different goroutine.     

```go
func recovery() {  
    if r := recover(); r != nil {
        fmt.Println("recovered:", r)
    }
}
func main() {  
    defer recovery()
    fmt.Println("Inside main goroutine")
    go b()
    time.Sleep(1 * time.Second)  // b goroutine start to run, then b goroutine panic, then program terminated/exited...
    fmt.Println("main here")
}
func b() {  
    fmt.Println("Inside B goroutine")
    panic("oops! B goroutine panicked")
}
```

#### Runtime panics
    Panics can also be caused by runtime errors such as array out of bounds access.
    This is equivalent to a call of the built-in function panic with an argument defined by interface type runtime.Error 

#### Getting stack trace after recover
**There is a way to print the stack trace using the PrintStack function of the Debug package after recover.**
```go
import (  
    "fmt"
    "runtime/debug"
)
func r() {  
    if r := recover(); r != nil {
        fmt.Println("Recovered", r)
        debug.PrintStack()              // print stack when needed
    }
}
```




***congratulations, you have finished golang course***
















======================================================================
        thanks to: https://golangbot.com/learn-golang-series/
======================================================================
[golangbot.com course](https://golangbot.com/learn-golang-series/)



