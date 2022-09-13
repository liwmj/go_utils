package go_utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/**
{
    "ip":"172.104.94.100",
    "location":{
        "country_code":"JP",
        "country_name":"Japan",
        "province":"Tokyo",
        "city":"Shinagawa",
        "latitude":"35.6130514",
        "longitude":"139.7344198"
    }
}
*/
type WanIP struct {
	IP       string   `json:"ip"`
	Location Location `json:"location"`
}

type Location struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	Province    string `json:"province"`
	City        string `json:"city"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

func GetWanIP() *WanIP {
	rsp, err := http.Get("https://api.myip.la/en?json")
	PanicError(err)
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	PanicError(err)

	// 解析 JSON 数据使用 json.Unmarshal([]byte(JSON_DATA),JSON对应的结构体) ,也就是说我们在解析 JSON 的时候需要确定 JSON 的数据结构
	wip := &WanIP{}
	json.Unmarshal([]byte(body), &wip)

	return wip
}
