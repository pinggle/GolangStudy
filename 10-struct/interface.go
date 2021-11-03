package main

import "fmt"

/**
interface 通用万用类型 interface{}
空接口
所有基本类型都实现了interface{}
*/
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// 给 interface{} 提供 "类型断言" 的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf("value type is %T\n", value)
	}
}

// AnimalIF 接口, 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// Cat 具体的类(猫),只用实现interface的方法,然后指向即可使用;
type Cat struct {
	color string //猫的颜色
}

func (this *Cat) GetType() string {
	fmt.Println("Cat type")
	return "Cat"
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

///
type Dog struct {
	color string //猫的颜色
}

func (this *Dog) GetType() string {
	fmt.Println("dog type")
	return "Dog"
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println("color=", animal.GetColor())
	fmt.Println("type=", animal.GetType())
}
func main() {
	var animal AnimalIF //接口的数据类型,父类指针
	animal = &Cat{"Green"}
	animal.Sleep() // 调用猫的sleep方法;

	animal = &Dog{"Yellow"}
	animal.Sleep()

	cat := Cat{"Red"}
	dog := Dog{"gray"}
	showAnimal(&cat)
	showAnimal(&dog)

	// 万能数据类型 interface{} 测试
	myFunc("fsdfdf")
	myFunc(1000)
}
