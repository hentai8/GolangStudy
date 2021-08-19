package main

import "fmt"

// 本质上是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 具体的类 重写AnimalIF接口
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "cat"
}

// 具体的类 重写AnimalIF接口
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimalSleep(animal AnimalIF) {
	// 多态
	animal.Sleep()
}

func main() {
	// 接口的数据类型，父类指针
	var animal AnimalIF
	// 多态的现象

	animal = &Cat{"Green"}
	// 调用的就是cat的Sleep()方法
	animal.Sleep()

	animal = &Dog{"Yellow"}
	animal.Sleep()

	cat := Cat{"blue"}
	dog := Dog{"white"}

	showAnimalSleep(&cat)
	showAnimalSleep(&dog)

}
