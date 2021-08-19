package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	// SuperMan类继承了Human类的方法
	Human
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat().............")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()~~~")
}

func main() {
	h := Human{"Zhang3", "male"}
	h.Eat()
	h.Walk()
	s := SuperMan{Human{"li4", "female"}, 88}
	s.Walk()
	s.Eat()
	s.Fly()
}
