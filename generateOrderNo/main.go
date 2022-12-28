package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	no := GenerateOrderNo("ckb")
	noOld := GenerateElectricItemNo()
	fmt.Println(no)
	fmt.Println(noOld)
}

func GenerateOrderNo(currency string) string {
	date := time.Now().Format("20060102")
	ts := time.Now().UnixNano() / 1e6 % 1e5
	r := rand.Intn(1000)
	no := fmt.Sprintf("%s%s%07d%03d", currency, date[2:], ts, r)
	return no
}

func GenerateElectricItemNo() string {
	date := time.Now().Format("20060102")
	ts := time.Now().UnixNano() / 1e6 % 1e5
	r := rand.Intn(1000)
	no := fmt.Sprintf("ckb%s%d%03d", date, ts, r)
	return no
}