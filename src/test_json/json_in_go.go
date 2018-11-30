package main

import (
	"encoding/json"
	"fmt"
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

func receive_json(resp http.ResponseWriter, req *http.Request) {
	var req_json map[string]interface{}

	// 1.parse args in req json:
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&req_json)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("--> receive request json:", req_json)
	fmt.Println("--> get name:", req_json["name"].(string))
	fmt.Println("--> get money:", req_json["money"].(float64))
	fmt.Println("--> get list:", req_json["list"].([]interface{}))
	fmt.Println("--> get list[0]:", req_json["list"].([]interface{})[0].(float64))
	fmt.Println("--> get map:", req_json["map"].(map[string]interface{}))
	fmt.Println("--> get map[k]:", req_json["map"].(map[string]interface{})["k"].(string))

	// 2.use simplejson to parse req json:

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
	if err != nil {
		fmt.Println(err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(resp_str)
}

func main() {
	http.HandleFunc("/r", receive_json)
	http.HandleFunc("/s", send_json)
	fmt.Println("Server Started at :8000")
	http.ListenAndServe(":8000", nil)
}
