package pool

type Pool chan *Object

func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(Object)
	}

	return &p
}

// 对象池的典型用法就是 线程池/连接池

// Usage
// Given below is a simple lifecycle example on an object pool.

p := pool.New(2)

select {
case obj := <-p:
	obj.Do( /*...*/ )

	p <- obj
default:
	// No more objects left — retry later or fail
	return
}