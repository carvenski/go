package main

import (
	"fmt"
)

// slice的头部和尾部和中间的CRUD操作
func main() {
	var a []int

	a = []int{0, 1, 2}
	fmt.Println(a)

	// 尾部添加
	a = append(a, 1)               // 尾部追加1个元素
	a = append(a, 1, 2, 3)         // 尾部追加多个元素
	a = append(a, []int{1,2,3}...) // 尾部追加一个切片, 切片需要解包

	// 头部添加
	a = append([]int{0}, a...)        // 在开头添加1个元素
	a = append([]int{-3,-2,-1}, a...) // 在开头添加多个元素

	// 尾部删除
	N := 2
	a = a[:len(a)-1]   // 删除尾部1个元素
	a = a[:len(a)-N]   // 删除尾部N个元素

	// 头部删除
	a = a[1:]   // 删除开头1个元素
	a = a[N:]   // 删除开头N个元素

	// 中间添加
	i, x := 1, 5
	a = append(a[:i], append([]int{x}, a[i:]...)...)     // 在第i个位置插入1个元素
	a = append(a[:i], append([]int{1,2,3}, a[i:]...)...) // 在第i个位置插入多个元素

	// 中间删除
	a = append(a[:i], a[i+1:]...) // 删除中间1个元素
	a = append(a[:i], a[i+N:]...) // 删除中间N个元素


}



