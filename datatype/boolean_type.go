package datatype

import "fmt"

func BooleanInit() {
	var v1 bool
	v1 = true
	v2 := (1 == 2) // v2 也会被推导为 bool 类型
	fmt.Println("布尔类型初始化打印===")
	fmt.Println(v1, v2)
}

// 编译报错示例
func BooleanJudge() {
	// 不同类型的值，不能够使用 == 或者 != 进行比较
	// b := (false == 0)

}
