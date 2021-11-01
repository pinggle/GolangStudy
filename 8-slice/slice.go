package main

import "fmt"

/**
定长数组和变长数组(slice)

*/

func printArray(inputArray []int) {
	// 引用传递;
	for index, value := range inputArray {
		fmt.Println("index=", index, ", value=", value)
	}

	if len(inputArray) > 0 {
		inputArray[0] = 999
	}
}

func judgeSliceIsEmpty(inputSlice []int) bool {
	var isEmpty bool = false
	if inputSlice == nil {
		fmt.Println("inputSlice 是一个空切片")
		isEmpty = true
	} else {
		fmt.Println("inputSlice 是有空间的")
	}
	return isEmpty
}

func main() {
	// 固定长度的数组;
	var myArray01 [10]int

	myArray02 := [10]int{1, 2, 3, 4}
	myArray03 := [4]int{11, 22, 33, 44}

	// 按照索引遍历数组;
	for i := 0; i < len(myArray01); i++ {
		fmt.Println("array[", i, "]: ", myArray01[i])
	}

	// 使用range关键字依次遍历;
	for index, value := range myArray02 {
		fmt.Println("index =", index, ", value: ", value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray01 type = %T\n", myArray01)
	fmt.Printf("myArray02 type = %T\n", myArray02)
	fmt.Printf("myArray03 type = %T\n", myArray03)

	// 变长数组,切片 slice
	myDynamicArray := []int{1, 2, 3, 4, 5}
	fmt.Printf("myDynamicArray type = %T\n", myDynamicArray)
	fmt.Println("before index[0] =", myDynamicArray[0])
	printArray(myDynamicArray) //传递引用;
	fmt.Println("after index[0] =", myDynamicArray[0])

	// 声明slice是一个切片,并且初始化,默认值为1，2，3
	//slice1 := []int{1,2,3}
	// 声明slice是一个切片;(容量为0)
	var slice01 []int
	fmt.Printf("before slice1 len=%d, slie=%v\n", len(slice01), slice01)
	// 开辟3个容量;
	slice01 = make([]int, 3)
	fmt.Printf("after slice1 len=%d, slie=%v\n", len(slice01), slice01)

	// 声明切片,并分配空间;
	var slice02 []int = make([]int, 3)
	slice03 := make([]int, 3)
	fmt.Printf("init slice02 len=%d, slie=%v\n", len(slice02), slice02)
	fmt.Printf("init slice03 len=%d, slie=%v\n", len(slice03), slice03)
	// 声明slice是一个切片;(容量为0)
	var slice04 []int

	fmt.Println("slice01 is empty : ", judgeSliceIsEmpty(slice01))
	fmt.Println("slice04 is empty : ", judgeSliceIsEmpty(slice04))

	/**
	切片容量的追加:
	1.切片的长度和容量不同,长度表示左指针至右指针之间的距离,容量表示左指针至底层数组末尾的距离;
	2.切片扩容机制,append的时候,如果长度增加后超过容量,则容量增加2倍;
	*/
	var numbers = make([]int, 3, 5)
	// len=3,cap=5,slice=[0 0 0]
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	// 向numbers切片追加1个元素1,numbers len=4,cap=5;
	numbers = append(numbers, 1)
	// len=4,cap=5,slice=[0 0 0 1]
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	// 向numbers切片追加1个元素2,numbers len=5,cap=5;
	numbers = append(numbers, 2)
	// len=5,cap=5,slice=[0 0 0 1 2]
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	// 向numbers切片追加1个元素3,numbers len=5,cap=5;
	numbers = append(numbers, 3)
	// len=6,cap=10,slice=[0 0 0 1 2 3] -- 容量自动增加cap;
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)

	/**
	切片的截取;
	*/
	s01 := []int{1, 2, 3} //len=3,cap=3
	s02 := s01[0:2]       // [1, 2]
	fmt.Println(s02)
	// copy 深拷贝;
	s03 := make([]int, 3)
	copy(s03, s01)
	s01[0] = 100
	fmt.Println(s01)
	fmt.Println(s03)

}
