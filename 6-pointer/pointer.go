package main

import "fmt"

func swapWithValue(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

func swapWithPointer(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	var a int = 10
	var b int = 20

	swapWithValue(a, b)
	fmt.Println("a = ", a, "; b = ", b)

	a = 10
	b = 20
	swapWithPointer(&a, &b)
	fmt.Println("a = ", a, "; b = ", b)

	// 二级指针
	var pp **int
	var p *int
	p = &a
	pp = &p
	fmt.Println("p = ", p, "; &a =", &a)
	fmt.Println("pp = ", pp, "; *pp =", *pp, "; &p = ", &p)

}
