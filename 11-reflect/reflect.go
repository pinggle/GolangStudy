package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", this)
}

func reflectNum(arg interface{}) {
	fmt.Println("type:", reflect.TypeOf(arg))
	fmt.Println("type:", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.2345
	reflectNum(num)

	user := User{1, "Hero", 18}
	user.Call()
	DoFiledAndMethod(user)

}

func DoFiledAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("input type is :", inputType.Name())

	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("value is :", inputValue)

	//通过type获取里面的字段
	//1.获取interface的reflect.type,通过type得到NumField,进行遍历;
	//2.得到每个field,数据类型
	//3.通过field有一个Interface()方法得到对应的value;
	for i := 0; i < inputType.NumField(); i++ {
		fieldNode := inputType.Field(i)
		fieldValue := inputValue.Field(i)
		fmt.Printf("%s: %v = %v\n", fieldNode.Name, fieldNode.Type, fieldValue)
	}

	//通过type获取里面的方法;
	//TODO:这里的函数打印为什么为空?
	for j := 0; j < inputType.NumMethod(); j++ {
		m := inputType.Method(j)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
