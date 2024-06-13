package processcontrol

import "fmt"

// for循环语句
func ForLoop() {
	// 基础的for循环语句
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("基础的for循环:", sum)
	// 无限循环
	sum2 := 0
	j := 0
	for {
		j++
		if j > 100 {
			break
		}
		sum2 += j
	}
	fmt.Println("无限循环:", sum2)
}

// 多重循环
func MultiLoop() {
	array := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	fmt.Println(array)
}

// for-range 结构
func ForRangeLoop() {
	array := []int{1, 2, 3, 4, 5, 6}
	for k, v := range array {
		fmt.Println(k, v)
	}
	fmt.Println("只获取键")
	// 只获取键
	for k := range array {
		fmt.Printf("%d ", k)
	}
	fmt.Println()
	fmt.Println("只获取值")
	// 只获取值
	for _, v := range array {
		fmt.Printf("%d ", v)
	}
}

// 基于判断条件实现循环
func BasedOnCondition() {
	sum := 0
	i := 0
	for i < 100 {
		i++
		sum += i
	}
	fmt.Println(sum)
}
