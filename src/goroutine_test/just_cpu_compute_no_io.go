package main
import ("fmt"
	"time"
)

func main(){
        fmt.Println("========start")
	for i := 0; i<10; i++{
           go t(i)
        }
        fmt.Println("========middle")
        time.Sleep(time.Second*3)
        fmt.Println("========end")
}

func t(i int){
	fmt.Println("------1--", i) //just cpu compute, not io block, why current goroutine also yield itself to let other goroutine run ?? 
	fmt.Println("------2--", i)
	fmt.Println("------3--", i)
	fmt.Println("------4--", i)
}


