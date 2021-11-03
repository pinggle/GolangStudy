package main

import (
	"fmt"
	"runtime"
	"time"
)

func SayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
		if i > 3 {
			runtime.Goexit()
		} else {
			runtime.Gosched()
		}
	}
}

func SayWorld() {
	for i := 0; i < 10; i++ {
		fmt.Println("World")
		if i > 3 {
			runtime.Goexit()
		} else {
			runtime.Gosched()
		}
	}
}

func main() {

	fmt.Println("CPU核数:", runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	go SayHello()
	go SayWorld()
	time.Sleep(5 * time.Second)
}
