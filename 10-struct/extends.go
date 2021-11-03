package main

import (
	"fmt"
)

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.eat() ...")
}

func (this *Human) Walk() {
	fmt.Println("Human.walk() ...")
}

//======== 继承类声明

type SuperMan struct {
	Human // SuperMan类继承了Human的类方法
	level int
}

// 重定义父类的Eat方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat() ...")
}

//子类新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly() ...")
}

func (this *SuperMan) Print() {
	fmt.Println("name=", this.name, "sex=", this.sex, "level=", this.level)
}

func main() {
	fmt.Println("继承示例")

	human := Human{"zhangsan", "female"}
	human.Eat()
	human.Walk()

	// 定义子类对象继承Human
	var superMan SuperMan
	superMan.name = "超人1号"
	superMan.sex = "男"
	superMan.level = 100

	superMan.Walk()
	superMan.Eat()
	superMan.Fly()

	superMan.Print()
}
