package datatype

import "fmt"

// 数组初始化
func ArrayInit() {
	var a [8]byte // 长度为8的数组，每个元素为一个字节
	fmt.Println("array a:", a)
	fmt.Println("======================")
	var b [3][3]int // 二维数组（9宫格）
	fmt.Println("array b:", b)
	fmt.Println("======================")
	var c [3][3][3]float64 // 三维数组（立体的9宫格）
	fmt.Println("array c:", c)
	fmt.Println("======================")
	var d = [3]int{1, 2, 3} // 声明时初始化
	fmt.Println("array d:", d)
	fmt.Println("======================")
	var e = new([3]string) // 通过 new 初始化
	fmt.Println("array e:", e)
	fmt.Println("======================")
	// 一次性声明以及初始化
	array := [3]int{1, 2, 3}
	fmt.Println("array:", array)
	// 省略数组长度的声明
	array2 := [...]int{1, 2, 3}
	fmt.Println("array2:", array2)
	// 数组初始化时，未填满的位置以数据类型的零值进行填充
	array3 := [5]int{1, 2, 3}
	fmt.Println("array3:", array3)
	//  使用索引号初始化数组
	array4 := [5]int{1: 3, 3: 5}
	fmt.Println("array4:", array4)
}

// 数组的遍历
func ArrayTraversal() {
	var arr = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(arr[5]) 越界异常
	arr[0] = 100
	fmt.Println("arr[0]:", arr[0])
	fmt.Println("for循环 遍历数组")
	// for循环遍历数组
	for i := 0; i < len(arr); i++ {
		fmt.Println("Element", i, "of arr is", arr[i])
	}
	fmt.Println("range 遍历数组")
	// range遍历数组
	for i, v := range arr {
		fmt.Println("Element", i, "of arr is", v)
	}
	fmt.Println("range 遍历数组只获取索引值")
	// range 遍历数组只获取索引值
	for i, _ := range arr {
		fmt.Println("index is ", i)
	}
	fmt.Println("range 遍历数组只获取元素值")
	// range 遍历数组只获取元素值
	for _, v := range arr {
		fmt.Println("value is ", v)
	}
}

// 多维数组
func MultiArray() {
	// 声明二维数组
	var multi [9][9]string
	// 二维数组元素赋值
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			m := i + 1
			n := j + 1
			if n > m {
				// 九九乘法表
				// 1*1 = 1
				// 1*2 = 2 2*2 = 4
				// 1*3 = 3 2*3 = 6 3*3 = 9
				// 当n>m，跳出循环
				continue
			}
			multi[i][j] = fmt.Sprintf("%d*%d=%d", n, m, m*n)
		}
	}
	// 二维数组打印
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if j > i {
				continue
			}
			fmt.Printf(multi[i][j] + " ")
			if j == i {
				fmt.Println()
			}
		}
	}
	fmt.Println("九九乘法表优化打印===")
	// 打印九九乘法表
	for _, v1 := range multi {
		for _, v2 := range v1 {
			fmt.Printf("%-8s", v2) // 位宽为8，左对齐
		}
		fmt.Println()
	}
}
