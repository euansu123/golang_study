package datatype

import "fmt"

func IntegerError() {
	var intValue1 int8
	intValue2 := 8 // Go语言会将未声明的数字默认为int类型
	// intValue1 = intValue2       // 编译错误，intValue1是int8类型，intValue2是int类型
	_ = intValue1
	intValue1 = int8(intValue2) // 正确，将intValue2转换为int8类型
	fmt.Printf("intValue1:%d, intValue2:%d\n", intValue1, intValue2)
}

func IntegerOperation() {
	// 整型运算
	var intValue1 int8
	intValue2 := 128
	_ = intValue1
	intValue1 = 81
	intValue3 := int(intValue1) + intValue2
	fmt.Println("intValue3:", intValue3)
}

func IntegerSelfOperation() {
	// 自增、自减运算
	intValue1 := 10
	intValue1++ // 有效，intValue1 的值变成 11
	intValue1-- // 有效，intValue1 的值变成 10
	// intValue1 = intValue1++ // 无效，编译报错
	// --intValue1  // 无效，编译报错
}

func IntegerOperationShort() {
	intValue1, intValue2 := 10, 20
	fmt.Println("intValue1:", intValue1)
	intValue1 += intValue2
	fmt.Println("intValue1:", intValue1)
	intValue1 -= intValue2
	fmt.Println("intValue1:", intValue1)
	intValue1 *= intValue2
	fmt.Println("intValue1:", intValue1)
	intValue1 /= intValue2
	fmt.Println("intValue1:", intValue1)
	intValue1 %= 3
	fmt.Println("intValue1:", intValue1)
}

// 整型的比较运算
func IntegetCompare() {
	intValue1, intValue2 := 10, 20
	if intValue1 <= intValue2 {
		fmt.Println("intValue1 <= intValue2")
	}
	var intValue3 int8 = 10
	if intValue3 == 10 {
		fmt.Println("intValue3 == 10")
	}

}

// 位运算符
func BitOperation() {
	var intValue1 uint8
	var intValue2 uint8
	intValue1 = 255                                            // 1111 1111
	intValue2 = 0                                              // 0000 0000
	fmt.Println("intValue1 & intValue2:", intValue1&intValue2) // 按位与，0
	fmt.Println("intValue1 | intValue2:", intValue1|intValue2) // 按位或，255
	fmt.Println("intValue1 ^ intValue2:", intValue1^intValue2) // 按位异或，255
	fmt.Println("^intValue1:", ^intValue1)                     // 按位取反 0
	fmt.Println("intValue1 << 1:", intValue1<<1)               // 左移1位，254
	fmt.Println("intValue1 >> 1:", intValue1>>1)               // 右移1位，127
}

// 逻辑运算符
func LogicOperation() {
	intValue1, intValue2 := 10, 20
	if intValue1 > 0 && intValue2 > 0 {
		fmt.Println("intValue1 > 0 && intValue2 > 0")
	}

	if intValue1 > 15 || intValue2 > 15 {
		fmt.Println("intValue1 > 15 || intValue2 > 15")
	}

	fmt.Println("!(intValue1 > 15):", !(intValue1 > 15))
}
