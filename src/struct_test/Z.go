package main

import (
	"fmt"
)

type X struct {
	Id   int
	Name string
}

func haha(id int) X {
	c := X{id, "haha"}
	return c
}

func haha2(id int) *X {
	c := X{id, "haha2"}
	return &c
}

func haha3(id int) map[string]int {
	c := map[string]int{"haha": id}
	return c
}
func haha4(id int) *map[string]int {
	c := map[string]int{"haha": id}
	return &c
}

func main() {
	a1 := haha(99)
	fmt.Println(a1.Id)
	a2 := haha2(100)
	fmt.Println(a2.Id)
	a3 := haha3(101)
	fmt.Println(a3["haha"])
	a4 := *haha4(102)
	fmt.Println(a4["haha"])
}
