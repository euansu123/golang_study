package datatype

import (
	"fmt"
	"strconv"
)

// 整型类型转换
func IntTypeConversion() {
	v1 := uint(16)   // 初始化 v1 类型为 unit
	v2 := int8(v1)   // 将 v1 转化为 int8 类型并赋值给 v2
	v3 := uint16(v2) // 将 v2 转化为 uint16 类型并赋值给 v3
	fmt.Println(v1, v2, v3)
	v4 := int16(-255)
	fmt.Println(v4)
}

// 整型浮点数类型转换
func IntTypeConversion2() {
	v1 := 99.99
	v2 := int(v1) // 浮点数转化为整型，小数点后的数字直接被抛弃
	fmt.Println(v2)

	v3 := 99
	v4 := float64(v3) // 整型转化为浮点数，直接调取对应的类型即可
	fmt.Println(v4)
}

// 整型和字符串类型转换
func IntTypeConversion3() {
	v1 := 65
	v2 := string(v1) // v2 = A
	v3 := 30028
	v4 := string(v3) // v4 = '界'
	fmt.Println(v2, v4)
}

// strconv包类型转换
func StrconvFunc() {
	v1 := "100"
	v2, _ := strconv.Atoi(v1) // 将字符串转化为整型，v2 = 100
	v3 := 100
	v4 := strconv.Itoa(v3) // 将整型转化为字符串, v4 = "100"
	v5 := "true"
	v6, _ := strconv.ParseBool(v5) // 将字符串转化为布尔型
	v5 = strconv.FormatBool(v6)    // 将布尔值转化为字符串
	v7 := "100"
	v8, _ := strconv.ParseInt(v7, 10, 64)  // 将字符串转化为整型，第二个参数表示进制，第三个参数表示最大位数
	v7 = strconv.FormatInt(v8, 10)         // 将整型转化为字符串，第二个参数表示进制
	v9, _ := strconv.ParseUint(v7, 10, 64) // 将字符串转化为无符号整型，参数含义同 ParseInt
	v7 = strconv.FormatUint(v9, 10)        // 将无符号整数型转化为字符串，参数含义同 FormatInt
	v10 := "99.99"
	v11, _ := strconv.ParseFloat(v10, 64) // 将字符串转化为浮点型，第二个参数表示精度
	v10 = strconv.FormatFloat(v11, 'E', -1, 64)
	q := strconv.Quote("Hello, 世界")       // 为字符串加引号
	q = strconv.QuoteToASCII("Hello, 世界") // 将字符串转化为 ASCII 编码
}
