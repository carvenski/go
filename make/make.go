package main

import (
	"log"
)

func main(){
	log.Println("== test make func ==")
	// slice
	l := make([]int, 10)
	// map
	m := make(map[string]int)
	// chan
	c := make(chan int, 1)
	log.Printf("slice= %v, len= %v", l, len(l))	
	log.Printf("map= %v", m)
	log.Printf("chan= %v", c)	
}