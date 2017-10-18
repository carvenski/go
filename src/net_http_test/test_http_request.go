package main
import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main(){

	httpGet()

}

func httpGet() {
	resp, err := http.Get("http://www.github.com/yxzoro/go")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}


