package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	err, num := TransferHashStringToInt("PH/S")
	if err != nil {
		return
	}
	fmt.Println(num)

}

func TransferHashStringToInt(s string) (err error, num int) {
	switch strings.ToUpper(s) {
	case "H/S":
		num = 1
	case "KH/S":
		num = 1000
	case "MH/s":
		num = 1000000
	case "GH/S":
		num = 1000000000
	case "TH/S":
		num = 1000000000000
	case "PH/S":
		num = 1000000000000000
	default:
		num = 0
	}
	if num == 0 {
		return errors.New("common.WRONGPARAMS"), num
	}
	return
}