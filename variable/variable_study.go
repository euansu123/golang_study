package variable

import "fmt"

func VariableDeclare() {
	// 定义变量及打印
	var v1 int      // 整型
	var v2 string   // 字符串
	var v3 bool     // 布尔类型
	var v4 [10]int  // 数组，数组元素类型为整型
	var v5 struct { // 结构体，成员变量 f 的类型为64位浮点型号
		f float64
	}
	var v6 *int            // 指针，指向整型
	var v7 map[string]int  // map（字典），key为字符串类型，value为整型
	var v8 func(a int) int // 函数，参数类型为整型，返回类型为整型
	fmt.Println("变量值打印开始===")
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)
	fmt.Println(v6)
	fmt.Println(v7)
	fmt.Println(v8)
	fmt.Println("变量值打印结束===")
}

func VariableInit() {
	var v1 int = 10 // 方式一，常规的初始化操作
	var v2 = 10     // 方式二，此时变量类型会被编译器自动推导出来
	v3 := 10        // 方式三，可以省略 var，编译器可以自动推导出v3的类型
	fmt.Println("变量值打印开始===")
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Printf("变量值打印结束===")
}

func VariableInitError() {
	//var i int
	//i := 10
}

func VariableaAsignment() {
	// 变量声明及变量初始化
	var v1 int
	v2 := 10
	v3 := 15
	// 赋值
	v1 = 10
	// 多重赋值
	v2, v3 = v3, v2
	fmt.Println("变量值打印开始===")
	fmt.Printf("v1 value is: %d\n", v1)
	fmt.Printf("v2 value is %d\n", v2)
	fmt.Printf("v3 value is %d\n", v3)
	fmt.Println("变量值打印结束===")
}

func GetName() (userName, nickName string) {
	return "南歌", "EuanSu"
}

func AnonymousVariables() {
	// 匿名变量
	// GetName返回两个值，分别是username和nickName，
	// 其中nickName是string类型，而GetName()函数中没有使用nickName，
	// 此时nickName就是一个匿名变量，它不占用命名空间，不会分配内存，
	// 因此不会产生任何开销。
	_, nickName := GetName()
	fmt.Printf("nickName value is %s\n", nickName)
}
