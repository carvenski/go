package main

import (
	"log"
)

func main()  {
	log.Println("--> test list")
	a := []int{1,2,3}
	log.Println("a", a)
	b := a
	b = append(b, 100)
	log.Println("b", b)
	log.Println("a", a)

	log.Println("--> test map")
	aa := map[string]int{"x": 1, "y":2}
	log.Println("aa", aa)
	bb := aa
	bb["z"] = 3
	log.Println("bb", bb)
	log.Println("aa", aa)

	log.Println("--> test struct")
	type Test struct{
		p int		
	}
	aaa := Test{1}
	log.Println("aaa", aaa)
	bbb := aaa
	bbb.p = 100
	log.Println("bbb", bbb)
	log.Println("aaa", aaa)

	
}

// --> test list
// a [1 2 3]
// b [1 2 3 100]
// a [1 2 3]
// list 是值类型！

// --> test map
// aa map[x:1 y:2]
// bb map[x:1 y:2 z:3]
// aa map[x:1 y:2 z:3]
// map 是引用类型！

// --> test struct
// aaa {1}
// bbb {100}
// aaa {1}
// struct 是值类型！

// *******************************************************************
// *******************************************************************
// 总结：
// go里面的 基本类型 + 数组 + list + struct 都是 值类型，不是引用类型 !!
// go里面的引用类型是 map + channel !!

// 其实也很好理解为什么go里面大多数类型都设计成值类型而不是引用类型的 => 因为go有指针啊!如果需要引用类型,自己用指针啊...
// *******************************************************************
// *******************************************************************


