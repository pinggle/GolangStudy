package main

import "fmt"

// 返回一个参数;
func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

// 返回两个参数;
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 600, 777
}

// 返回(带参数名)的返回值
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}

// 返回(带参数名)的返回值
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}

// 返回(带参数名)的返回值
func foo5(a string, b int) (r1 string, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 给有名称的返回值变量赋值
	r1 = a
	r2 = b
	return
}

func main() {

	var tmpString string = "123"
	var tmpParam int = 100

	// foo1 return : 100
	fmt.Println("foo1 return :", foo1(tmpString, tmpParam))
	ret1, ret2 := foo2(tmpString, tmpParam)
	// foo2 return ret1 : 600 , ret2 : 777
	fmt.Println("foo2 return ret1 :", ret1, ", ret2 :", ret2)
	ret1, ret2 = foo3(tmpString, tmpParam)
	// foo3 return ret1 : 1000 , ret2 : 2000
	fmt.Println("foo3 return ret1 :", ret1, ", ret2 :", ret2)
	ret1, ret2 = foo4(tmpString, tmpParam)
	// foo4 return ret1 : 1000 , ret2 : 2000
	fmt.Println("foo4 return ret1 :", ret1, ", ret2 :", ret2)
	retString, retVal := foo5(tmpString, tmpParam)
	// foo5 return retString : 123 , retVal : 100
	fmt.Println("foo5 return retString :", retString, ", retVal :", retVal)

}
