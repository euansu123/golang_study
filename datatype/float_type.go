package datatype

import (
	"fmt"
	"math"
)

func FloatInit() {
	// 浮点数定义
	var floatValue1 float32
	floatValue1 = 10
	floatValue2 := 10.0 // 如果不加小数点，floatValue2 会被推导为整型而不是浮点型
	floatValue3 := 1.1e-10
	fmt.Println("定义的浮点数打印===")
	fmt.Println("floatValue1:", floatValue1)
	fmt.Println("floatValue2:", floatValue2)
	fmt.Println("floatValue3:", floatValue3)
	floatValue1 = float32(floatValue2)
	floatValue4 := 0.1
	floatValue5 := 0.7
	floatValue6 := floatValue4 + floatValue5
	fmt.Println(floatValue6)
}

// 浮点数运算

func FloatCalculate() {
	floatValue4 := 0.1
	floatValue5 := 0.7
	floatValue6 := floatValue4 + floatValue5
	fmt.Println("floatValue4 + floatValue5 = ", floatValue6)

	epsilon := 1e-10
	sum := 0.8
	if math.Abs(sum-floatValue6) < epsilon {
		fmt.Println("sum and c are approximately equal")
	} else {
		fmt.Println("sum and c are not equal")
	}
}

// 浮点数的比较

func FloatCompare() {
	floatValue1 := 0.1
	floatValue2 := 0.1
	p := 0.00001
	// 判断 floatValue1 与 floatValue2 是否相等
	if math.Dim(float64(floatValue1), floatValue2) < p {
		fmt.Println("floatValue1 和 floatValue2 相等")
	}
}
