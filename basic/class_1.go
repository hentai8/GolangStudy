package main

import "fmt"

// 如果类名首字母大写，表示其他包也能访问
type Hero struct {
	// 如果说类的属性首字母大写，表示该属性是对外能够访问的，否则的话只能够内部访问
	Name  string
	Level int
	AD    int
}

/*
func (this Hero) Show()  {
	fmt.Println("hero = ", this)
}

func (this Hero) GetName()  {
	fmt.Println("Name = ",this.Name)
}

func (this Hero) SetName(newName string)  {
	// this是调用该方法的对象的一个副本（拷贝）
	this.Name = newName
}

*/

func (this *Hero) Show() {
	fmt.Println("hero = ", this)
}

func (this *Hero) GetName() {
	fmt.Println("Name = ", this.Name)
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	hero := Hero{Name: "wuke", AD: 100, Level: 1}
	hero.GetName()
	hero.Show()
	hero.SetName("naotan")
	hero.Show()
}
