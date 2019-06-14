package main

import (
	"fmt"
	"protobuf/myproto"
	"github.com/golang/protobuf/proto"
)

func checkErr(err error)  {
	if err != nil {
		fmt.Println("error => %s", err)
	}
}

func main()  {
	pbObj := &myproto.VMInfoConfig{Cpu: 1, Memory: 2}

	pbStr, err := proto.Marshal(pbObj)
	checkErr(err)
	fmt.Println("pb => str")
	fmt.Println(pbStr)

	pbObj2 := &myproto.VMInfoConfig{}
    proto.Unmarshal(pbStr, pbObj2)
	fmt.Println("str => pb")
	fmt.Println(pbObj2)

}









