package function

import (
	"errors"
	"fmt"
)

// 变长参数，能够传递任意个数的参数，但要求类型统一
func Myfunc(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

// 变长参数，可以传递任意个数的参数，类型无要求
func Myfunc2(params ...interface{}) {
	for _, param := range params {
		fmt.Println(param)
	}
}

// 多返回值
func AddFunc(a, b *int) (int, error) {
	if *a < 0 || *b < 0 {
		err := errors.New("仅支持非负整数的相加")
		return 0, err
	}
	return *a + *b, nil
}

// 指定返回值变量
func AddFunc2(a, b *int) (c int, err error) {
	if *a < 0 || *b < 0 {
		err = errors.New("仅支持非负整数的相加")
		return
	}
	c = *a + *b
	return
}

func AnonymousFunc() {
	//// 匿名函数
	//// 将匿名函数赋值给变量
	//sum := func(a, b int) int {
	//	return a + b
	//}
	//// 调用匿名函数 add
	//fmt.Println(sum(1, 2))
	//
	//// 也可以在定义的时候,直接调用匿名函数
	//func(a, b int) {
	//	fmt.Println(a + b)
	//}(1, 2)

	// 匿名函数的使用场景
	// 1.保证局部变量的安全性
	//var j int = 1
	//f := func() {
	//	var i int = 1
	//	fmt.Printf("i value is %d,j value is %d\n", i, j)
	//}
	//f()
	//j += 2
	//f()

	// 2.将匿名函数作为函数参数
	//add := func(a, b int) int {
	//	return a + b
	//}
	//func(call func(int, int) int) {
	//	fmt.Println(call(1, 2))
	//}(add)

	// 3.将函数作为返回值类型
	// 此时返回的是匿名函数
	addFunc := defaultAdd(1, 2)
	// 这里才会真正的执行加法操作
	fmt.Println(addFunc())
}

func defaultAdd(a, b int) func() int {
	return func() int {
		return a + b
	}
}
