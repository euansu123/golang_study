package processcontrol

import "fmt"

// 分支语句示例
func SwitchDemo() {
	score := 100
	switch {
	case score > 85:
		fmt.Println("优秀")
	case score > 70 && score <= 85:
		fmt.Println("良好")
	case score >= 60 && score <= 70:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}

// 分支语句的合并
func SwitchMergeDemo() {
	month := 1
	switch month {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12:
		fmt.Println("冬季")
		fallthrough
	case 1, 2:
		fmt.Println("冬季")
	}
}
