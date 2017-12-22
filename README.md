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

#### Structure







































