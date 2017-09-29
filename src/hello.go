package main

import "fmt"

func main() {
	fmt.Println("========hello go=========")
	p := test()
	fmt.Println(p.name)
	fmt.Println(p.age)

}

type Person struct {
	name string
	age  int
}

func test() Person {
	p := Person{name: "haha", age: 24}
	p.name = "hi"
	return p
}
