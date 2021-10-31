package main

import (
	"fmt"
	"time"
)

/**
go build hello.go
go run hello.go
*/
func main() { // 函数的左括号({)一定和函数名在同一行;
	// 语句后面可加可不加分号;
	fmt.Println("hello go")
	time.Sleep(1 * time.Second)
	fmt.Println("sleep over")
}
