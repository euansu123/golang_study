package constant

import "fmt"

func ConstantDefine() {
	// 常量的定义
	const Pi float64 = 3.14159265358979323846
	const zero = 0.0 // 无类型浮点常量
	const (          // 通过一个 const 关键字定义多个常量，和 var 类似
		size int64 = 1024
		eof        = -1 // 无类型整型常量
	)
	const u, v float32 = 0, 3   // u = 0.0, v = 3.0，常量的多重赋值
	const a, b, c = 3, 4, "foo" // a = 3, b = 4, c = "foo", 无类型整型和字符串常量
	fmt.Println("常量值的打印===")
	fmt.Printf("Pi value is %f\n", Pi)
	fmt.Printf("zero value is %f\n", zero)
	fmt.Printf("size value is %d, eof value is %d\n", size, eof)
	fmt.Printf("u value is %f, v value is %f\n", u, v)
	fmt.Printf("a value is %d, b value is %d, c value is %s\n", a, b, c)
}

func GetNumber() int {
	return 100
}

// 常量的定义错误，执行该代码时会在编译时候报错，提示GetNumber()不是常量
//func ConstantDefineError() {
//	const num = GetNumber()
//}

func PredefineConstant() {
	// 预定义常量
	const (
		c0 = iota // c0 = 0
		c1 = iota // c1 = 1
		c2 = iota // c2 = 2
	)
	// iota 被重置为0
	const (
		u = iota * 2 // u = 0
		v = iota * 2 // v = 2
		w = iota * 2 // w = 4
	)
	// iota 被重置为0
	const x = iota // x = 0
	const y = iota // y = 0
	fmt.Println("预定义变量打印===")
	fmt.Printf("c0 value is %d, c1 value is %d, c2 value is %d\n", c0, c1, c2)
	fmt.Printf("u value is %d, v value is %d, w value is %d\n", u, v, w)
	fmt.Printf("x value is %d, y value is %d\n", x, y)
}

func PredefineConstantOmit() {
	const (
		c0 = iota
		c1
		c2
	)
	const (
		u = iota * 2
		v
		w
	)
	fmt.Println("预定义常量定义的时候可以进行省略===")
	fmt.Printf("c0 value is %d, c1 value is %d, c2 value is %d\n", c0, c1, c2)
	fmt.Printf("u value is %d, v value is %d, w value is %d\n", u, v, w)
}

func EnumFunc() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
	fmt.Println("枚举常量===")
	fmt.Printf("Sunday is %d,Monday value is %d,Tuesday value is %d,Wednesday value is %d,Thursday value is %d,Friday value is %d,Saturday value is %d,numberOfDays value is %d\n", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, numberOfDays)
}
