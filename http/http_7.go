package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"
    "strings"
)

type CurrencyPrice struct {
    Code int64 `json:"code"`
    Data struct {
        Price struct {
            Usd map[string]string `json:"USD"`
        } `json:"price"`
    } `json:"data"`
    Message string `json:"message"`
    Status  string `json:"status"`
}

func main() {
    //duration := time.Second * 60 * 5
    //timer := time.NewTimer(duration)
    //go func() {
    //    for {
    //        select {
    //        case <-timer.C:
    //var currencyParams Currency.GetCurrenciesParams
    //currencyParams.PageSize = 100
    //currencies, err := Currency.GetCurrencies(currencyParams)
    //if err != nil {
    //    fmt.Println(err)
    //    timer.Reset(duration)
    //    continue
    //}
    var currencies []string
    currencies = append(currencies, "ckb")
    for _, currency := range currencies {
        coin := strings.ToLower(currency)
        err, dailyPrice := GetDailyPrice(coin)
        fmt.Println(dailyPrice)
        if err != nil {
            fmt.Println(err)
            continue
        }
        //basicCurrency := currency.BasicCurrencyStruct
        //basicCurrency.Price = dailyPrice
        //err = Currency.ChangeCurrencyById(currency.Id, basicCurrency)
        //if err != nil {
        //    fmt.Println(err)
        //    timer.Reset(duration)
        //    continue
        //}

        //                timer.Reset(duration)
        //            }
        //        }
        //}()
    }
}

func GetDailyPrice(coin string) (err error, price float64) {
    var currencyPrice CurrencyPrice

    url := "http://8.212.8.124:3000/price?base=usd&symbol=" + coin
    method := "GET"

    client := &http.Client{}
    req, err := http.NewRequest(method, url, nil)

    if err != nil {
        fmt.Println(err)
        return
    }
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
    err = json.Unmarshal(body, &currencyPrice)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(currencyPrice)
    if currencyPrice.Code != 200 {
        err := errors.New("failed to get price of the currency: " + coin)
        return err, 0
    }

    for name, value := range currencyPrice.Data.Price.Usd {
        if strings.ToUpper(name) == strings.ToUpper(coin) {
            fmt.Println(value)
            price, err = strconv.ParseFloat(value, 64)
            if err != nil {
                fmt.Println(err)
                return
            }
            fmt.Println(price)
            fmt.Println("end")
        }
    }
    return
}
