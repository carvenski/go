package main

import (
	"encoding/json"
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"net/http"
)

/*
req json example:
curl --header "Content-Type: application/json" \
  --request POST \
  --data \
  '{"name":"michael","money": 50000,"list":[1,2,3,4,5],"map":{"k": "v"}}' \
  http://localhost:8000/r

resp json example:
curl http://localhost:8000/s
{"name": "michael", "money": 50000, "list":[1,2,3,4,5], "map":{"k": "v"}}
*/

type Map map[string]interface{}
type List []interface{}

func checkErr(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

func receive_json_1(resp http.ResponseWriter, req *http.Request) {
	var req_json map[string]interface{}

	// 1.use type assertion parse args in req json:
	fmt.Println("[use type assertion]:")
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&req_json)
	checkErr(err)
	fmt.Println("--> receive request json:", req_json)
	fmt.Println("--> get name:", req_json["name"].(string))
	fmt.Println("--> get money:", req_json["money"].(float64))
	fmt.Println("--> get list:", req_json["list"].([]interface{}))
	fmt.Println("--> get list[0]:", req_json["list"].([]interface{})[0].(float64))
	fmt.Println("--> get map:", req_json["map"].(map[string]interface{}))
	fmt.Println("--> get map[k]:", req_json["map"].(map[string]interface{})["k"].(string))
}

func receive_json_2(resp http.ResponseWriter, req *http.Request) {
	// 2.use simplejson to parse req json:
	//   see https://godoc.org/github.com/bitly/go-simplejson
	// actually, simplejson use type assertion inside too...
	fmt.Println("\n\n[use simplejson]:")
	js, err := simplejson.NewFromReader(req.Body)
	checkErr(err)
	fmt.Println("--> receive request json:", js)
	name, err := js.Get("name").String()
	checkErr(err)
	fmt.Println("--> get name:", name)
}

func send_json(resp http.ResponseWriter, req *http.Request) {
	var resp_json map[string]interface{}

	//return json resp:
	resp_json = map[string]interface{}{}
	resp_json["name"] = "michael"
	resp_json["money"] = 50000
	resp_json["list"] = []int{1, 2, 3, 4, 5}
	resp_json["map"] = map[string]string{"k": "v"}
	fmt.Println("--> send response json:", resp_json)
	resp_str, err := json.Marshal(resp_json)
	checkErr(err)
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(resp_str)
}

func main() {
	http.HandleFunc("/r1", receive_json_1)
	http.HandleFunc("/r2", receive_json_2)
	http.HandleFunc("/s", send_json)
	fmt.Println("Server Started at :8000")
	http.ListenAndServe(":8000", nil)
}
