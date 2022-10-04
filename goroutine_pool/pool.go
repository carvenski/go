package main

import (
	"log"
    ants "github.com/panjf2000/ants/v2"
    echo "github.com/labstack/echo/v4"
)

func main() {
	log.Println("== goroutine pool")
	log.Println(echo)
	log.Println(ants)
}

// task function and arg
type Arg struct {
	a int
	b int
}

func task(a *Arg) int {
	return a.a + a.b
}

// Pool
type Pool struct {
	min int
	max int
}

func (p Pool) Init() {

}

func (p Pool) Submit() {

}
