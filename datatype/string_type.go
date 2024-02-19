package datatype

import "fmt"

func StringInit() {
	var str string        // 声明字符串变量
	str = "Hello World"   // 变量初始化
	str2 := "Hello World" // 也可以同时进行声明和初始化
	fmt.Println(str)
	fmt.Println(str2)
	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
}

// 字符串方法
func StringMethod() {
	// 引号打印输出
	label := `'Search' results for "Golang":`
	fmt.Println(label)
	// 多行字符串
	results := `Search results for "Golang":
	- Go
	- Golang
	- Golang Programming
	`
	fmt.Println(results)
	results = "Search results for \"Golang\":\n" +
		"- Go\n" +
		"- Golang\n" +
		"- Golang Programming\n"
	fmt.Printf("%s", results)
	str := "Hello World"
	//str[0] = 'M'
	str = "Mello World"
	fmt.Println(str)
}

// 字符串操作
func StringOperate() {
	str := "Hello"
	str = str + ",World"
	fmt.Println(str)
	str2 := "Hello"
	str2 += ",World"
	fmt.Println(str2)
	str3 := "Hello"
	str3 = str3 +
		",World"
	fmt.Println(str3)
}

// 字符串切片
func StringSlice() {
	str := "hello, world"
	str1 := str[:5]  // 获取索引5（不含）之前的子串
	str2 := str[7:]  // 获取索引7（含）之后的子串
	str3 := str[0:5] // 获取从索引0（含）到索引5（不含）之间的子串
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)
	fmt.Println("str3:", str3)
}

// 字符串遍历
func StringTraversal() {
	// 方式一：字节数组
	fmt.Println("方式一：字节数组===")
	str := "Hello, 世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，ch 类型为 byte
		fmt.Println(i, ch)
	}
	fmt.Println("=================")
	fmt.Println("方式二：unicode字符===")
	// 方式二：unicode字符
	str2 := "Hello, 世界"
	for i, ch := range str2 {
		fmt.Println(i, ch) // ch 的类型为 rune
	}
	fmt.Println("=================")
}
