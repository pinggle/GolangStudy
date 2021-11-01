package main

import "fmt"

// 声明一种新的数据类型 myint，是int的一个别名;
type myint int

// 定义一个结构体;
type Book struct {
	title string
	auth  string
}

func printBook(book Book) {
	fmt.Println("title=", book.title, "auth=", book.auth)
}

// 如果类名首字母大写,表示其他包也能够访问
type Hero struct {
	//如果类的属性大写,表示其他包也能够访问
	Name  string
	Ad    int
	Level int
}

// GetName绑定到Hero结构体;
func (this *Hero) GetName() {
	fmt.Println("Name = ", this.Name)
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func (this *Hero) Show() {
	fmt.Println("Name=", this.Name, "Ad=", this.Ad, "Level=", this.Level)
}

func main() {
	var a myint = 10
	fmt.Println("a = ", a)

	var book01 Book
	book01.title = "golang"
	book01.auth = "zhangsan"
	fmt.Printf("%v\n", book01)
	printBook(book01)

	// 创建一个对象;
	hero := Hero{Name: "zhangsan", Ad: 100, Level: 10}
	hero.Show()
	hero.SetName("wangwu")
	hero.Show()
}
