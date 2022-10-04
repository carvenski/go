package main

import "log"

// 使用slice构造一个Queue对象
type Queue struct {
	queue []int
}
func (q *Queue) Push(v int) {
	log.Println(q.queue) 
	log.Println("Push ing", v) 	
	q.queue = append(q.queue, v)
	log.Println(q.queue, "\n") 
}
func (q *Queue) Pop() int {
	log.Println(q.queue) 
	v := q.queue[len(q.queue)-1]
	log.Println("Pop ing", v) 	
	q.queue = q.queue[0:(len(q.queue)-1)]
	log.Println(q.queue, "\n") 
	return v
}

func main(){
	q := &Queue{[]int{1,2,3}}
	q.Push(4)
	q.Push(5)
	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
}