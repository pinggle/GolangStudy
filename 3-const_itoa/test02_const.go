package main

import "fmt"

// const定义枚举类型
const (
	// 可以在const()添加一个关键字iota,第一行的iota的默认值是0,每行的iota都会累加1;
	BEIJING  = iota // iota = 0
	SHANGHAI        // iota = 1
	SHENZHEN        // iota = 2
)

const (
	// iota只能在const定义中;
	a, b = iota + 1, iota + 2 // iota=0, a=0+1=1, b=0+2=2;
	c, d                      // iota=1, c=1+1=2, d=1+2=2;
	e, f                      // iota=2, e=2+1=3, f=2+2=4;
	g, h = iota * 2, iota * 3 // iota=3, g=3*2=6, h=3*3=9;
	i, j                      // iota=4, i=4*2=8, j=4*3=12
)

func main() {

	//常量(只读属性)
	const length int = 10

	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING, " SHANGHAI = ", SHANGHAI, " SHENZHEN = ", SHENZHEN)
	fmt.Println("a = ", a, " b = ", b)
	fmt.Println("c = ", c, " d = ", d)
	fmt.Println("e = ", e, " f = ", f)
	fmt.Println("g = ", g, " h = ", h)
	fmt.Println("i = ", i, " j = ", j)
}
