package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

type CheckoutCreate struct {
	LocalPrice `json:"local_price"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PricingType string `json:"pricing_type"`
	RedirectURL string `json:"redirect_url"`
	CancelURL   string `json:"cancel_url"`
}

type LocalPrice struct {
	Currency string `json:"currency"`
	Amount float64 `json:"amount"`
}


func main() {
	var checkoutCreate CheckoutCreate
	checkoutCreate.LocalPrice.Currency = "BCH"
	checkoutCreate.LocalPrice.Amount = 0.0005
	checkoutCreate.Name = "qwq"
	checkoutCreate.Description = "qaq"
	checkoutCreate.PricingType = "fixed_price"
	checkoutCreate.RedirectURL = "https://charge/completed/page"
	checkoutCreate.CancelURL = "https://charge/canceled/page"

	checkoutCreateString, err := json.Marshal(checkoutCreate)
	url := "https://api.commerce.coinbase.com/charges"
	method := "POST"

	payload := strings.NewReader(string(checkoutCreateString))

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("X-CC-Api-Key", "2c30c770-632b-4478-a1e6-7f09e7b97eee")
  req.Header.Add("Accept", "application/json")
  req.Header.Add("X-CC-Version", "2018-03-22")
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}