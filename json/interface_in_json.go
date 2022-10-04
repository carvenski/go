package main

import (
	"encoding/json"
	"fmt"
)

type Map map[string]interface{}
type List []interface{}

func checkErr(err error) {
	if err != nil {
		fmt.Println("err=:", err)
	}
}

func main() {
	var obj Map = Map{}
	obj["name"] = "michael"
	obj["money"] = 50000
	obj["list"] = []int{1, 2, 3, 4, 5}
	obj["map"] = map[string]string{"k": "v"}
	obj_str, err := json.Marshal(obj)
	checkErr(err)
	fmt.Println(string(obj_str))
	// 会自动识别golang的数据结构并正确的序列化成json对应的的数据结构
	//{"list":[1,2,3,4,5],"map":{"k":"v"},"money":50000,"name":"michael"}
}

/*
=== 关于golang中http使用json做参数和返回值的技巧 ===
对应的resp的json可以这么来定义:
map[string]interface{}
而真正传的时候可以传各种数据结构来填充interface{},
json序列化时会自动识别并转换的.
{
	"code": 0,
	"errmsg": "",
	"data": real_data_structure,
}

request的json参数一样是map[string]interface{},
然后做个类型转换就行了
request_json["xx1"].(string)
request_json["xx2"].([]string)

这样就可以使用一个统一的数据结构来描述所有的json值了,而不用一一写结构体.
*/