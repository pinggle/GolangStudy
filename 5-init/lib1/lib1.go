package lib1

import "fmt"

// 函数名第一个字母为大写,就是public;
func Lib1Api01() {
	fmt.Println("lib1, Lib1Api01() ...")
}

func init() {
	fmt.Println("lib1, init() ...")
}
