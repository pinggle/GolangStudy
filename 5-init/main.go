package main

/**
文件分布:
 ~/go/src/GolangStudy/5-init  tree
.
├── lib1
│	└── lib1.go
├── lib2
│	└── lib2.go
├── lib3
│	└── lib3.go
├── lib4
│	└── lib4.go
└── main.go

**/
// 导入自定义包;
import (
	"GolangStudy/5-init/lib1"
	// 用下划线导入,表示匿名导入包,无法使用包的方法,但是会执行包内的init函数
	_ "GolangStudy/5-init/lib2"
	// Lib3Library 别名,给导入包起个别名,使用别名可调用包内函数;
	Lib3Library "GolangStudy/5-init/lib3"
	// . 导入包中的全部方法到本包的作用域中,包中的方法可直接调用,不需要fmt.API来调用;
	. "GolangStudy/5-init/lib4"
)

func main() {
	// 调用自定义包中的自定义函数;
	lib1.Lib1Api01()
	Lib3Library.Lib3Api01()
	Lib4Api01()
}
