

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

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



