import (
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

// golang中的装饰器是不能装饰任意函数的,
// 一个装饰器必须针对一类有固定参数和返回值的函数来写,因为类型不是动态的,
// 那么做法就是在程序的最外层给每个请求定义一个Controller函数,它们的参数和返回都相同,
// 然后转装饰器就可以加在它们上面了.

// HTTP Server SafeHandler decorator
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

// add for loop decorator
func loop(f func()) func() {
	return func() {
		for {
			f()
			time.Sleep(time.Second * 1)
		}
	}
}



