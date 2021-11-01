package main

import "fmt"

func printMap(cityMap map[string]string) {
	// cityMap 是一个引用传递;(指针传递)
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func main() {

	//===> 第1种声明方式;
	//声明map01是一种map类型,key是string,value是string
	var map01 map[string]string
	if map01 == nil {
		fmt.Println("map01 是一个空map")
	}

	// 在使用map前,需要先用make给map分配数据空间
	map01 = make(map[string]string, 10)
	map01["one"] = "java"
	map01["two"] = "c++"
	map01["three"] = "python"
	fmt.Println(map01)

	//===> 第2种声明方式;
	map02 := make(map[int]string)
	map02[1] = "java"
	map02[2] = "c++"
	map02[3] = "python"
	fmt.Println(map02)

	//===> 第3种声明方式;
	map03 := map[string]string{
		"one":   "java",
		"two":   "c++",
		"three": "python",
	}
	fmt.Println(map03)

	//
	cityMap := make(map[string]string)
	cityMap["china"] = "beijing"
	cityMap["japan"] = "tokyo"
	cityMap["usa"] = "newyork"
	//遍历
	printMap(cityMap)

	// 删除
	delete(cityMap, "japan")

	// 修改
	cityMap["usa"] = "dc"

	fmt.Println("----------------")
	printMap(cityMap)

}
