package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// 练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
	fmt.Println("execute command is: ", os.Args[0])

	// 练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
	for argIndex, argValue := range os.Args[1:] {
		fmt.Println("args[", argIndex, "]=", argValue)
	}

	// 练习 1.3： 做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。

	fmt.Println(strings.Join(os.Args[1:], " "))
}
