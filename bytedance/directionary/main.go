package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Request struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
}

type Response struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func main() {
	client := &http.Client{}
	var data Request
	data.TransType = "en2zh"
	data.Source = "good"
	//var data = strings.NewReader(`{"trans_type":"en2zh","source":"good"}`)

	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(dataJson))
	req, err := http.NewRequest("POST", "https://lingocloud.caiyunapp.com/v1/dict", strings.NewReader(string(dataJson)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", "_gcl_au=1.1.74046197.1684131810; _ga=GA1.2.808227757.1684131810; _gid=GA1.2.215716807.1684131810; _gat_gtag_UA_185151443_2=1; _ga_65TZCJSDBD=GS1.1.1684131809.1.0.1684131810.0.0.0; _ga_R9YPR75N68=GS1.1.1684131809.1.0.1684131810.59.0.0")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("app-name", "xy")
	req.Header.Set("device-id", "6561bcf10d6dae965b54da230dc3c4a0")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)

	var r Response
	err = json.Unmarshal(bodyText, &r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("1%+v\n", r)
}
