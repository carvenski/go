package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/url"

	"strings"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"encoding/base64"
	"io/ioutil"

	"myproxy/httputil"
)

type Handler struct {
	backends []string
}

func (this *Handler) getBackend() string {
	// random select
	return this.backends[rand.Intn(len(this.backends))]
}

func getDataReqId(r *http.Request) string {
	// 目前只cache 13率的data请求
	if r.Method == "POST" && r.Body != nil && 
		strings.Contains(r.URL.String(), "v5/design/report/share/widget/data?") && 
		strings.Contains(r.URL.String(), "reportId=86547fecb7ba4ad59b2ea68534ea9d60") {

		var reqId string = ""
		var buf []byte
		buf, _ = ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

		r2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		var req_json map[string]interface{}
		decoder := json.NewDecoder(r2)
		_ = decoder.Decode(&req_json)		

		// 清空timeStamp/sessionId
		req_json["timeStamp"] = ""
		req_json["sessionId"] = ""
		// 点击地图的5个data请求：使用整个body生成reqId
		if req_json["tableName"].(string) == "BI_雷达图_本年" {
			// req_json["tableName"].(string) == "本年累计本月上月应收实收保费" || 
			// req_json["tableName"].(string) == "个险13率基本统计汇总表1.0" {
				log.Println("=> Map id")
				h := md5.New()
			    // h.Write(buf)
			    h.Write([]byte(req_json["tableName"].(string)))
			    reqId += string(h.Sum(nil))
		// 其他的data请求,使用其中几个字段来生成reqId
		} else {
			return ""
			log.Println("=> Other id")
			if req_json["dimensions"] != nil {
				dimensions, _ := json.Marshal(req_json["dimensions"].(map[string]interface{}))
				reqId += base64.StdEncoding.EncodeToString(dimensions)
			}
			if req_json["dimensionGroups"] != nil {
				dimensionGroups, _ := json.Marshal(req_json["dimensionGroups"].(map[string]interface{}))
				reqId += base64.StdEncoding.EncodeToString(dimensionGroups)
			}
			if req_json["measures"] != nil {
				measures, _ := json.Marshal(req_json["measures"].([]interface{}))
				reqId += base64.StdEncoding.EncodeToString(measures)
			}
			if req_json["page"] != nil {
				page, _ := json.Marshal(req_json["page"].(float64))
				reqId += string(page)
			}
			if req_json["resultFilter"] != nil {
				resultFilter, _ := json.Marshal(req_json["resultFilter"].([]interface{}))
				reqId += base64.StdEncoding.EncodeToString(resultFilter)
			}

		}
		log.Println("parse reqId ok=") //, reqId)
		return reqId
	}
	return ""
}

func (this *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// add r.ReqId
	if reqId := getDataReqId(r); reqId != "" {
		r.ReqId = reqId
	} else {
		r.ReqId = ""
	}

	backend, err := url.Parse(this.getBackend())
	//log.Println("proxy to backend=", backend)
	if err != nil {
		log.Fatalln(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(backend)
	r.Host = backend.Host
	proxy.ServeHTTP(w, r)
	//log.Println(r.backendAddr + r.Method + r.URL.String() + r.Proto + r.UserAgent())
}

func main() {
	bind := ":8002"
	backends := []string{
		"http://21.50.131.35:37799",
	}
	log.Printf("Proxy Listen on %s", bind)
	log.Printf("Proxy Forward to %s", backends)
	handler := &Handler{backends: backends}
	err := http.ListenAndServe(bind, handler)
	if err != nil {
		log.Fatalln("err: ", err)
	}
}
