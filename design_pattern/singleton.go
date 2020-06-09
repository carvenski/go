package singleton

import "sync"

type singleton map[string]string

var (
	// sync.Once的用法,保证只调用一次函数,正好用来实现单例模式
	once sync.Once

	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}

// s := singleton.New()

// s["this"] = "that"

// s2 := singleton.New()

// fmt.Println("This is ", s2["this"])
// // This is that
