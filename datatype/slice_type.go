package datatype

import "fmt"

func NewSlice() {
	// 基于数组创建切片
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	q2 := months[3:6]     // 第二季度
	summer := months[5:8] // 夏季
	fmt.Println("q2:", q2)
	fmt.Println("summer:", summer)

	// 基于切片创建切片
	firsthalf := months[:6]
	q1 := firsthalf[:3]
	fmt.Println("q1:", q1)
	// 可以创建超过切片的元素
	q3 := q1[:12]
	fmt.Println("q3:", q3)

	// 直接创建切片
	mySlice := make([]int, 5)
	mySlice2 := make([]int, 5, 10)
	mySlice3 := []int{1, 2, 3, 4, 5}
	fmt.Println("mySlice:", mySlice, "len:", len(mySlice), "cap:", cap(mySlice))
	fmt.Println("mySlice2:", mySlice2, "len:", len(mySlice2), "cap:", cap(mySlice2))
	fmt.Println("mySlice3:", mySlice3, "len:", len(mySlice3), "cap:", cap(mySlice3))
}

// 切片的遍历
func TraverseSlice() {
	mySlice := make([]int, 5)
	for i := 0; i < len(mySlice); i++ {
		fmt.Println("mySlice[", i, "] =", mySlice[i])
	}
	fmt.Println("range 方式遍历切片")
	// range 方式遍历
	for i, v := range mySlice {
		fmt.Println("mySlice[", i, "] =", v)
	}
}

// 动态增加元素
func AddElement() {
	// len() 长度函数
	// cap() 容量函数
	var oldSlice = make([]int, 5, 10)
	fmt.Println("len(oldSlice):", len(oldSlice))
	fmt.Println("cap(oldSlice):", cap(oldSlice))
	fmt.Println("动态增加元素")
	newSlice := append(oldSlice, 1, 2, 3)
	fmt.Println("oldSlice:", oldSlice, "len:", len(oldSlice), "cap:", cap(oldSlice))
	fmt.Println("newSlice:", newSlice, "len:", len(newSlice), "cap:", cap(newSlice))
	slice2 := []int{1, 2, 3, 4, 5}
	// 注意append()后面的...不能省略
	slice3 := append(newSlice, slice2...)
	fmt.Println("slice3:", slice3, "len:", len(slice3), "cap:", cap(slice3))
}

// 自动扩容
func AutoExpansion() {
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := append(oldSlice, 6, 7, 8, 9)
	fmt.Println("oldSlice:", oldSlice, "len:", len(oldSlice), "cap:", cap(oldSlice))
	fmt.Println("newSlice:", newSlice, "len:", len(newSlice), "cap:", cap(newSlice))
}

// 内容复制
func ContentCopy() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}
	// 复制slice1到slice2，复制slice1的前三个元素到slice2中
	fmt.Println("复制slice1到slice2")
	copy(slice2, slice1)
	fmt.Println("slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1))
	fmt.Println("slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := []int{6, 7, 8}
	fmt.Println("复制slice4到slice3")
	// 复制slice4到slice3，复制slice4的所有元素到slice3中
	copy(slice3, slice4)
	fmt.Println("slice3:", slice3, "len:", len(slice3), "cap:", cap(slice3))
	fmt.Println("slice4:", slice4, "len:", len(slice4), "cap:", cap(slice4))
}

// 切片的动态删除
func DynamicDelete() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 = slice1[:len(slice1)-5] // 删除 slice3 尾部 5 个元素
	fmt.Println("slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1))
	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 = slice2[5:] // 删除 slice3 头部 5 个元素
	fmt.Println("slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))
	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("使用append实现元素的删除")
	fmt.Println(slice3)
	slice4 := append(slice3[:0], slice3[3:]...) // 删除开头三个元素
	fmt.Println("slice4:", slice4, "len:", len(slice4), "cap:", cap(slice4))
	fmt.Println("slice3:", slice3, "len:", len(slice3), "cap:", cap(slice3))

	slice5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice6 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice7 := append(slice5[:3], slice6[6:]...)
	fmt.Println("slice7:", slice7, "len:", len(slice7), "cap:", cap(slice7))
	fmt.Println("slice6:", slice6, "len:", len(slice6), "cap:", cap(slice6))
	fmt.Println("slice5:", slice5, "len:", len(slice5), "cap:", cap(slice5))

	fmt.Println("使用copy实现元素的删除")
	slice8 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice9 := make([]int, len(slice3)-3)
	copy(slice9, slice3[3:]) // 删除开头前三个元素
	fmt.Println("slice9:", slice9, "len:", len(slice9), "cap:", cap(slice9))
	fmt.Println("slice8:", slice8, "len:", len(slice8), "cap:", cap(slice8))
}

// 切片的数据共享问题
func DataSharing() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:3]
	slice2[1] = 6
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("避免切片的数据共享问题")
	slice3 := make([]int, 4)
	slice4 := slice3[1:3]
	slice3 = append(slice3, 0)
	slice3[1] = 2
	slice4[1] = 6
	fmt.Println("slice3:", slice3)
	fmt.Println("slice4:", slice4)

	slice5 := make([]int, 4, 5)
	slice6 := slice5[1:3]
	slice5 = append(slice5, 0)
	slice5[1] = 2
	slice6[1] = 6
	fmt.Println("slice5:", slice5)
	fmt.Println("slice6:", slice6)
}
