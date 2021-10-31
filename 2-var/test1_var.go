package main

import "fmt"

// 方法一,方法二,方法三可以声明全局变量;
var global_a = 100

func main() {
	// 方法一:声明一个变量,默认值为0;
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("a = %T\n", a)

	// 方法二:声明一个变量,初始化一个值;
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("b = %T\n", b)

	// 方法三:在初初始化的时候,可以省去数据类型,通过值自动匹配当前你的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("c = %T\n", c)

	var bb string = "abcd"
	fmt.Println("bb = ", bb)
	fmt.Printf("bb = %T\n", bb)

	var cc = "dabcd"
	fmt.Println("cc = ", cc)
	fmt.Printf("cc = %T\n", cc)

	// 方法四(最常用)省去var关键字,直接自动匹配
	// := 只能用在内部函数,不能声明全局变量;
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("e = %T\n", e)

	f := "abcd"
	fmt.Println("f = ", f)
	fmt.Printf("f = %T\n", f)

	g := 3.1415926
	fmt.Println("g = ", g)
	fmt.Printf("g = %T\n", g)

	fmt.Println("global_a = ", global_a)

	var xx, yy int = 100, 200
	var kk, ll = 100, "abcd"
	fmt.Println("xx = ", xx, ",yy = ", yy)
	fmt.Println("kk = ", kk, ",ll = ", ll)

	// 多行的多变量声明;
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ",jj = ", jj)
}
