package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (user *User) Call() {
	fmt.Println("user is called")
	fmt.Printf("%v\n", user)
}

func DoFiledAndMethod(input interface{}) {
	//获取input的type和value
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	//通过Type获取里面的字段
	//（1）获取interface的reflect.Type，通过Type得到NumField，进行遍历
	//（2）得到每个field，数据类型
	//（3）通过field有一个Interface()方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// xiamiandefangfa chucuole

	//通过type获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		//println()
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func main() {
	user := User{1, "hentai8", 18}
	DoFiledAndMethod(user)
	//DoFiledAndMethod1(&user)

}
