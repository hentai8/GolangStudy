package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api.prime.coinbase.com/v1/entities/entity_id/assets"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-CB-ACCESS-KEY", "$ACCESS_KEY")
	req.Header.Add("X-CB-ACCESS-PASSPHRASE", "$PASSPHRASE")
	req.Header.Add("X-CB-ACCESS-SIGNATURE", "$SIGNATURE")
	req.Header.Add("X-CB-ACCESS-TIMESTAMP", "$TIMESTAMP")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}