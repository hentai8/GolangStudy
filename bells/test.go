package main

import (
	"errors"
	"fmt"
	"github.com/asmcos/requests"
	"strings"
)

func main() {
	getReq := requests.Requests()
	getReq.SetTimeout(20)
	resp, err := getReq.Get("http://54.184.85.102:15000" + "/metrics")
	if err != nil {
		fmt.Printf("failed to get %s block height, err: %v \n", "ssv", err.Error())
	} else {
		resString := resp.Text()
		if strings.Contains(resString, "ssv_node_status 1") {
			fmt.Println(1, nil)
		} else {
			err := errors.New("ssv node status is 0")
			fmt.Println(0, err)
		}
	}
}
