package processcontrol

import "fmt"

// 条件语句
func ConditionDemo() {
	score := 100
	if score > 85 {
		fmt.Println("优秀")
	} else if score > 70 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}
