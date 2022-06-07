package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api.commerce.coinbase.com/checkouts/checkout_id"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-CC-Api-Key", "2c30c770-632b-4478-a1e6-7f09e7b97eee")
	req.Header.Add("X-CC-Version", "2018-03-22")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}