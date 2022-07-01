package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	// var s decimal.Decimal
	a := decimal.NewFromFloat(0.00)
	ss := decimal.NewFromInt(0)
	s, _ := decimal.NewFromString("0.00")

	b := decimal.NewFromFloat(0.000)
	fmt.Println(a.Equal(ss))
	fmt.Println(b.Float64())
	fmt.Println(s)
	fmt.Println(a==b)
	fmt.Println(a==s)
}
