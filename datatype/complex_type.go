package datatype

import "fmt"

// 复数的定义
func ComplexInit() {
	var complexValue1 complex64
	complexValue1 = 1.10 + 10i         // 由两个 float32 实数构成的复数类型
	complexValue2 := 1.10 + 10i        // 和浮点型一样，默认自动推导的实数类型是 float64，所以 complexValue2 是 complex128 类型
	complexValue3 := complex(1.10, 10) // 与 complexValue2 等价
	fmt.Println("complexValue1:", complexValue1)
	fmt.Println("complexValue2:", complexValue2)
	fmt.Println("complexValue3:", complexValue3)

	real := real(complexValue1) // 获取复数的实部
	imag := imag(complexValue1) // 获取复数的虚部
	fmt.Println("real:", real, "imag:", imag)
}
