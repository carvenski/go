package main

import (
	"fmt"
)

func modify_array(x [5]int)  {
	x[1] = 10000
	fmt.Println("array in func: ", x)
}

func modify_list(x []int)  {
	x[1] = 10000
	fmt.Println("list in func: ", x)
}

func modify_map(x map[string]int)  {
	x["m"] = 10000
	fmt.Println("map in func: ", x)
}

type Test struct{
	p int		
	q int
}

func modify_struct(x Test)  {
	x.q = 10000
	fmt.Println("struct in func: ", x)
}

func main()  {
	fmt.Println("--> test array")
	array_a := [5]int{1,2,3}
	fmt.Println("array a=", array_a)
	array_b := array_a
	array_b[0] = 100
	fmt.Println("array b=", array_b)
	fmt.Println("array a=", array_a)       // 数组是 值类型 的

	fmt.Println("--> test list append")
	list_a := []int{1,2,3}
	fmt.Println("list a=", list_a)
	list_b := list_a
	// (注意: append函数并不是修改当前list,而是使用当前的list创建一个新的list再返回...草, list是引用类型的)
	list_b = append(list_b, 100)  
	fmt.Println("list b=", list_b)
	fmt.Println("list a=", list_a)

	fmt.Println("--> test list")
	list_aa := []int{1,2,3}
	fmt.Println("list a=", list_aa)
	list_bb := list_aa
	list_bb[0] = 100                       // list是 引用类型 的.
	fmt.Println("list b=", list_bb)
	fmt.Println("list a=", list_aa)

	fmt.Println("--> test map")
	map_a := map[string]int{"x": 1, "y":2}
	fmt.Println("map a=", map_a)
	map_b := map_a
	map_b["z"] = 3
	fmt.Println("map b=", map_b)
	fmt.Println("map a=", map_a)           // map是 引用类型 的

	fmt.Println("--> test struct")
	struct_a := Test{1, 2}
	fmt.Println("struct a=", struct_a)
	struct_b := struct_a
	struct_b.p = 100
	fmt.Println("struct b=", struct_b)
	fmt.Println("struct a=", struct_a)      // struct是 值类型 的
	
	fmt.Println("--> test array in func")
	fmt.Println("array a=", array_a)
	modify_array(array_a)
	fmt.Println("array a=", array_a)

	fmt.Println("--> test list in func")
	fmt.Println("list a=", list_aa)
	modify_list(list_aa)
	fmt.Println("list a=", list_aa)

	fmt.Println("--> test map in func")
	fmt.Println("map a=", map_a)
	modify_map(map_a)
	fmt.Println("map a=", map_a)

	fmt.Println("--> test struct in func")
	fmt.Println("struct a=", struct_a)
	modify_struct(struct_a)	
	fmt.Println("struct a=", struct_a)

}

/* 输出
--> test array
array a= [1 2 3 0 0]
array b= [100 2 3 0 0]
array a= [1 2 3 0 0]
--> test list append
list a= [1 2 3]
list b= [1 2 3 100]
list a= [1 2 3]
--> test list
list a= [1 2 3]
list b= [100 2 3]
list a= [100 2 3]
--> test map
map a= map[x:1 y:2]
map b= map[x:1 y:2 z:3]
map a= map[x:1 y:2 z:3]
--> test struct
struct a= {1 2}
struct b= {100 2}
struct a= {1 2}
--> test array in func
array a= [1 2 3 0 0]
array in func:  [1 10000 3 0 0]
array a= [1 2 3 0 0]
--> test list in func
list a= [100 2 3]
list in func:  [100 10000 3]
list a= [100 10000 3]
--> test map in func
map a= map[x:1 y:2 z:3]
map in func:  map[x:1 y:2 z:3 m:10000]
map a= map[x:1 y:2 z:3 m:10000]
--> test struct in func
struct a= {1 2}
struct in func:  {1 10000}
struct a= {1 2}
*/



