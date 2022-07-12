package main

import (
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

type res struct {
	Timestamp   int64
	Data        []*MiningPoolResp
	Mu          sync.Mutex
}

type MiningPoolResp struct {
	Name    string `json:"name"`
	Height  int64 `json:"height"`
	Algo    string `json:"algo"`
	Symbol  string `json:"symbol"`
}

type MiningPoolStatsCache struct {
	Timestamp   int64
	Data        map[string]*MiningPoolResp
	Mu          sync.Mutex
}

func main() {
	var c MiningPoolStatsCache
	c.UpdateNetworkHeight()
}

func (c *MiningPoolStatsCache) UpdateNetworkHeight() bool {
	resp, err := http.Get("https://miningpoolstats.stream/")
	if err != nil {
		fmt.Println("http get mining pool stats page error: ", err)
		return false
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read http mining pool stats response error", err)
		return false
	}
	str := string(body)
	reg := regexp.MustCompile("var last_time = \"[0-9]{10}\";")
	data := reg.Find([]byte(str))
	str = string(data)
	reg = regexp.MustCompile("[0-9]{10}")
	data = reg.Find([]byte(str))
	ts := string(data)
	req := requests.Requests()
	req.SetTimeout(10)
	p := requests.Params{
		"t":ts,
	}
	fmt.Println("t:", ts)
	response, err := req.Get("https://data.miningpoolstats.stream/data/coins_data.js", p)
	if err != nil {
		println("error call miningpoolstats api: ", err)
	}

	// fmt.Println("response: ", response.Content())

	var tmp res

	err = json.Unmarshal(response.Content(), &tmp)

	// fmt.Println("tmp:", tmp)
	if err != nil {
		println("error unmarshal mining pool stats api response, err: ", err.Error())
		return false
	}
	for _, v := range tmp.Data {
		c.Data[strings.ToLower(v.Symbol)] = v
	}
	c.Timestamp = time.Now().Unix()
	return true
}