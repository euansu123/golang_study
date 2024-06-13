package processcontrol

import "fmt"

// break与continue
func JumpExample() {
	// break循环
	sum := 0
	i := 0
	for {
		i++
		if i > 100 {
			break
		}
		sum += i
	}
	fmt.Println("sum:", sum)
	fmt.Println("i:", i)
	fmt.Println("continue语句")
	// continue语句
	for i = 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}

// 标签
func BreakLabelExample() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
ITERATOR1:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := arr[i][j]
			if j > 1 {
				break ITERATOR1
			}
			fmt.Println(num)
		}
	}
}

func ContinueLabelExample() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
ITERATOR1:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := arr[i][j]
			if j > 1 {
				continue ITERATOR1
			}
			fmt.Println(num)
		}
	}
}

// goto语句
func GotoExample() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := arr[i][j]
			if j > 1 {
				goto EXIT
			}
			fmt.Println(num)
		}
	}
EXIT:
	fmt.Println("EXIT.")
}
