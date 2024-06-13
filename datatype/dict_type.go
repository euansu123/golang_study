package datatype

import (
	"fmt"
	"sort"
)

// 字典示例
func DictExample() {

	var tempMap map[string]int
	tempMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	k := "two"
	v, ok := tempMap[k]
	// ok，如果key存在，则ok为true，否则ok为false
	if ok {
		fmt.Printf("the element of key %q:%d", k, v)
	} else {
		fmt.Println("Not found!")
	}
}

// 字典的声明和初始化
func DictInit() {
	testMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println("testMap:", testMap)
	var tempMap = make(map[string]int)
	tempMap["one"] = 1
	fmt.Println("tempMap:", tempMap)
}

// 字典的遍历
func DictTraversal() {
	testMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for key, value := range testMap {
		fmt.Println(key, value)
	}
	fmt.Println("只获取字典的值")
	// 只获取字典的值
	for _, value := range testMap {
		fmt.Println(value)
	}
	fmt.Println("只获取字典的键")
	// 只获取字典的键
	for key := range testMap {
		fmt.Println(key)
	}
}

// 键值对调
func DictSwap() {
	testMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	newMap := make(map[int]string)

	for k, v := range testMap {
		newMap[v] = k
	}

	for k, v := range newMap {
		fmt.Println(k, v)
	}
}

// 字典排序
func DictKeySort() {
	testMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	// 创建切片
	keys := make([]string, 0)
	for key := range testMap {
		keys = append(keys, key)
	}
	// 切片排序
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(v, testMap[v])
	}
}

func DictValueSort() {
	testMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	// 创建切片
	values := make([]int, 0)
	for _, value := range testMap {
		values = append(values, value)
	}
	// 对字典的值进行排序
	sort.Ints(values)
	for _, v := range values {
		fmt.Println(v)
	}
}
