package main

import "fmt"

/**
defer 类似c++中的析构函数;
多个defer是以压栈的形式执行,后压入的先执行;
1.多个defer是先进后出;
2.defer在return后执行;

*/

func deferFunc() int {
	fmt.Println("defer func called ...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called ...")
	return 0
}

/**
return 和 defer 哪个函数先触发?
return 先执行, defer 后执行;
defer 是在函数执行完毕后,最后执行的操作;
*/
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	// defer依次入栈,调用顺序,依次出栈;(先进后出)
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")
	defer fmt.Println("main end3")

	fmt.Println("main call func1")
	fmt.Println("main call func2")

	returnAndDefer()

}
