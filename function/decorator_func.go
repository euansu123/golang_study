package function

import (
	"fmt"
	"time"
)

func Multiply(a, b int) int {
	return a * b
}

// 为函数设置别名提高代码可读性
type MultiPlyFunc func(a, b int) int

// 高阶函数实现装饰器
func ExecTime(f MultiPlyFunc) MultiPlyFunc {
	return func(a, b int) int {
		start := time.Now() // 开始时间
		fmt.Println(start)
		c := f(a, b) //执行函数运算
		end := time.Since(start)
		fmt.Println("执行耗时：", end)
		return c
	}
}
