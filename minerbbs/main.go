package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Prices struct {
	Data struct {
		Points map[string]Price
	} `json:"data"`
	Status struct {
		CreditCount  int64  `json:"credit_count"`
		Elapsed      string `json:"elapsed"`
		ErrorCode    string `json:"error_code"`
		ErrorMessage string `json:"error_message"`
		Timestamp    string `json:"timestamp"`
	} `json:"status"`
}

type Price struct {
	V []float64 `json:"v"`
}

func main()  {
	err, monthlyPrice := GetMonthlyPrice()
	if err != nil{
		fmt.Println(err)
	}

	err, monthlyDifficulty := GetMonthlyDifficulty()
	if err != nil{
		fmt.Println(err)
	}

	CalculateCKBLastMonthOutput(1000000, 5000, 270, 0.00005, 0.00005, "GH/s", monthlyDifficulty, monthlyPrice)


	//duration := time.Second * 90
	//timer := time.NewTimer(duration)
	//go func() {
	//	for {
	//		select {
	//		case <-timer.C:
	//			err, monthlyPrice := GetMonthlyPrice()
	//			if err != nil{
	//				fmt.Println(err)
	//				timer.Reset(duration)
	//				continue
	//			}
	//
	//			err, monthlyDifficulty := GetMonthlyDifficulty()
	//			if err != nil{
	//				fmt.Println(err)
	//				timer.Reset(duration)
	//				continue
	//			}
	//
	//			CalculateCKBLastMonthOutput(1000000000, monthlyDifficulty, monthlyPrice)
	//			timer.Reset(duration)
	//		}
	//	}
	//}()
}


func GetMonthlyPrice() (err error, monthlyPrice float64) {
	var prices Prices

	url := "https://api.coinmarketcap.com/data-api/v3/cryptocurrency/detail/chart?id=4948&range=1M"
	method := "GET"

	client := &http.Client {
	}
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
	err = json.Unmarshal(body, &prices)
	if err != nil {
		fmt.Println(err)
		return
	}

	err, monthlyPrice = CalculateMonthlyPrice(prices)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func CalculateMonthlyPrice(prices Prices) (err error, monthlyPrice float64) {
	var sum float64
	for _, value := range prices.Data.Points {
		sum += value.V[0]
	}
	fmt.Println(sum)
	monthlyPrice = sum/float64(len(prices.Data.Points))
	return
}

func GetMonthlyDifficulty() (err error, monthlyDifficulty float64) {
	result := new([][]interface{})

	url := "https://minerbbs.com/tools/pos/tubiao/jsonnd.php?id=517"
	method := "GET"

	client := &http.Client {
	}
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
	err = json.Unmarshal(body, &result)

	//fmt.Println(result)
	//fmt.Println(len(*result))
	var timeAll []float64
	var difficultyAll []float64
	for _, v := range *result {
		//fmt.Println(v[0])
		//fmt.Println(reflect.TypeOf(v[0]))
		//fmt.Println(reflect.TypeOf(v[1]))
		time := v[0].(float64)
		difficulty := v[1].(float64)
		//fmt.Println(time)
		//fmt.Println(difficulty)
		timeAll = append(timeAll, time)
		difficultyAll = append(difficultyAll, difficulty)
	}
	err, monthlyDifficulty = CalculateMonthlyDifficulty(timeAll, difficultyAll)
	//fmt.Println(string(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func CalculateMonthlyDifficulty(timeAll []float64, difficultyAll []float64) (err error , monthlyDifficulty float64){
	endTime := float64(time.Now().UnixNano()/1000000)
	startTime := endTime - 2592000000
	var startIndex int
	var endIndex int
	for i, v := range timeAll {
		if v <= startTime {
			startIndex = i
		}
		if v <= endTime {
			endIndex = i
		}
	}
	fmt.Println(startIndex)
	fmt.Println(endIndex)
	fmt.Println(difficultyAll[startIndex:endIndex])
	difficultyLastMonth := difficultyAll[startIndex:endIndex]
	var sum float64
	for _, v := range difficultyLastMonth{
		sum += v
	}
	monthlyDifficulty = sum/float64(len(difficultyLastMonth))
	fmt.Println(monthlyDifficulty)
	return
}



func CalculateCKBLastMonthOutput(userDifficulty float64, userPrice float64, userDuration float64, hashrateFee float64, electricityFee float64, hashrateUnit string, monthlyDifficulty float64, monthlyPrice float64)  {
	var userDifficultyAbsolute float64
	if hashrateUnit == "H/s" {
		userDifficultyAbsolute = userDifficulty * 1
	}
	if hashrateUnit == "KH/s" {
		userDifficultyAbsolute = userDifficulty * 1000
	}
	if hashrateUnit == "MH/s" {
		userDifficultyAbsolute = userDifficulty * 1000000
	}
	if hashrateUnit == "GH/s" {
		userDifficultyAbsolute = userDifficulty * 1000000000
	}
	if hashrateUnit == "TH/s" {
		userDifficultyAbsolute = userDifficulty * 1000000000000
	}
	if hashrateUnit == "PH/s" {
		userDifficultyAbsolute = userDifficulty * 1000000000000000
	}

	fmt.Println("userDifficulty:", userDifficulty)
	fmt.Println("userDifficultyAbsolute:", userDifficultyAbsolute)
	fmt.Println("monthlyDifficulty:", monthlyDifficulty)
	fmt.Println("monthlyPrice:", monthlyPrice)
	monthlyHistoricalOutput := userDifficultyAbsolute/monthlyDifficulty*monthlyPrice*1234.68096637*(30*24*60*60/15) - (hashrateFee+electricityFee)*userDifficulty
	fmt.Println(monthlyHistoricalOutput)
	monthlyHistoricalOutputRate := monthlyHistoricalOutput/userPrice
	fmt.Println(monthlyHistoricalOutputRate)
	AnnualizedHistoricalOutputRate := monthlyHistoricalOutput*12/(userPrice*365/userDuration)
	fmt.Println(AnnualizedHistoricalOutputRate)
	daysOfInvestmentRecovery := int64(userPrice * 30 / monthlyHistoricalOutput)
	fmt.Println(daysOfInvestmentRecovery)

}

