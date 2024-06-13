package datatype

import (
	"fmt"
	"unsafe"
)

// 指针的简单示例
func PointerExample() {
	a := 100
	var ptr *int // 声明指针类型
	ptr = &a     // 初始化指针类型值为变量 a
	fmt.Println(ptr)
	fmt.Println(*ptr)
}

// 指针的声明和初始化
func PointerDeclarationAndInitialization() {
	var ptr *int
	fmt.Println(ptr)

	a := 100
	ptr = &a
	fmt.Println(ptr)
	fmt.Println(*ptr)

	b := 100
	ptr2 := &b
	fmt.Printf("%p\n", ptr2)
	fmt.Printf("%d\n", *ptr2)
	fmt.Println("通过new初始化指针")
	ptr3 := new(int)
	fmt.Println(ptr3)
	fmt.Println(*ptr3)
	*ptr3 = 100
	fmt.Printf("%p\n", ptr3)
	fmt.Printf("%d\n", *ptr3)
}

func swap(a, b int) {
	a, b = b, a
	fmt.Println(a, b)
}
func pointerSwap(a, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

// 值拷贝
func PointerValueCopyExample() {
	a := 10
	b := 20
	fmt.Println("直接进行值拷贝")
	swap(a, b)
	fmt.Println(a, b)
	fmt.Println("通过指针进行值交换")
	pointerSwap(&a, &b)
	fmt.Println(a, b)

}

// 指针类型的转换
func PointerTypeConversionExample() {
	i := 10
	var p *int = &i
	// int类型的指针p -> unsafe.Pointer类型的指针 -> float32类型的指针
	var fp *float32 = (*float32)(unsafe.Pointer(p))
	*fp = *fp * 10
	fmt.Println(i)
}

// 指针运算实现
func PointerArithmeticExample() {
	arr := [3]int{1, 2, 3}
	ap := &arr
	// unsafe.Sizeof 数组元素偏移量
	// ap由unsafe.Pointer -> uintptr -> unsafe.Pointer
	fmt.Println("arr[0]的偏移量（内存大小）", unsafe.Sizeof(arr[0]))
	fmt.Println("arr的偏移量（内存大小）", unsafe.Sizeof(arr))
	sp := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ap)) + unsafe.Sizeof(arr[0])))
	*sp += 3
	fmt.Println(arr)
}
