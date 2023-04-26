package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	// 使用随机数之前需要设置随机数种子
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	//fmt.Println("The secrete number is ", secretNumber)

	fmt.Println("Please input your guess number: ")
	// 读取一行的输入，转成一个只读的流
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred: ", err)
			continue
		}

		// 替换其中的\n
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("An error occurred: ", err)
			continue
		}
		fmt.Println("Your guess is ", guess)
		if guess > secretNumber {
			fmt.Println("Too big")
		} else if guess < secretNumber {
			fmt.Println("Too small")
		} else {
			fmt.Println("Bingo")
			break
		}
	}
}
