package main

import (
	"fmt"
	"golang_study/GoClass"
	"golang_study/constant"
	"golang_study/datatype"
	"golang_study/function"
	"golang_study/processcontrol"
	"golang_study/variable"
)

func variableFunc() {
	// 调用变量声明方法
	// variable.VariableDeclare()
	// 调用变量初始化方法
	// variable.VariableInit()
	// 测试变量初始化一个提前声明的变量
	// variable.VariableInitError()
	// 调用变量赋值
	// variable.VariableaAsignment()
	// 匿名变量调用
	variable.AnonymousVariables()
}

func constantFunc() {
	// 常量的基础代码
	// constant.ConstantDefine()
	// 常量的定义错误
	// constant.ConstantDefineError()
	// 预定义常量
	// constant.PredefineConstant()
	// 预定义常量的省略
	// constant.PredefineConstantOmit()
	// 常量枚举
	constant.EnumFunc()
}

func dataTypeFunc() {
	// 布尔类型初始化
	// datatype.BooleanInit()
	// 布尔类型False条件判断
	// datatype.BooleanJudge()
	// 整型类型编译错误
	// datatype.IntegerError()
	// 整型运算
	// datatype.IntegerOperation()
	// 自增/自减运算
	// datatype.IntegerSelfOperation()
	// 运算简写示例
	// datatype.IntegerOperationShort()
	// 比较运算符
	// datatype.IntegetCompare()
	// 位运算符
	// datatype.BitOperation()
	// 逻辑运算符
	// datatype.LogicOperation()
	// 浮点数定义
	// datatype.FloatInit()
	// 浮点数运算
	// datatype.FloatCalculate()
	// 浮点数比较
	// datatype.FloatCompare()
	// 复数的定义
	// datatype.ComplexInit()
	// 字符串的定义
	// datatype.StringInit()
	// 字符串方法
	// datatype.StringMethod()
	// 字符串操作
	// datatype.StringOperate()
	// 字符串切片
	// datatype.StringSlice()
	// 字符串遍历
	// datatype.StringTraversal()
	// 整型类型转换
	// datatype.IntTypeConversion()
	// 整型浮点数类型转换
	// datatype.IntTypeConversion2()
	// 整型字符串类型转换
	// datatype.IntTypeConversion3()
	// 数组初始化
	// datatype.ArrayInit()
	// 数组的遍历
	// datatype.ArrayTraversal()
	// 多维数组
	// datatype.MultiArray()
	// 创建切片
	// datatype.NewSlice()
	// 遍历切片
	// datatype.TraverseSlice()
	// 动态增加切片元素
	// datatype.AddElement()
	// 切片自动扩容
	// datatype.AutoExpansion()
	// 切片的内容复制
	// datatype.ContentCopy()
	// 切片动态删除元素
	// datatype.DynamicDelete()
	// 切片的数据共享
	// datatype.DataSharing()
	// 字典的简单实例
	// datatype.DictExample()
	// 字典的声明&初始化
	// datatype.DictInit()
	// 字典的遍历
	// datatype.DictTraversal()
	// 字典键值对调
	// datatype.DictSwap()
	// 字典键的排序
	// datatype.DictKeySort()
	// 字典值的排序
	// datatype.DictValueSort()
	// 指针类型的简单示例
	// datatype.PointerExample()
	// 指针的声明和初始化
	// datatype.PointerDeclarationAndInitialization()
	// 通过指针传值
	// datatype.PointerValueCopyExample()
	// 指针类型的转换
	// datatype.PointerTypeConversionExample()
	// 指针类型的运算
	datatype.PointerArithmeticExample()
}

func flowControlFunc() {
	// 条件语句
	// processcontrol.ConditionDemo()
	// 分支语句
	// processcontrol.SwitchDemo()
	// 分支语句的合并
	// processcontrol.SwitchMergeDemo()
	// 循环语句
	// processcontrol.ForLoop()
	// 多重循环
	// processcontrol.MultiLoop()
	// for-range 结构
	// processcontrol.ForRangeLoop()
	// 基于判断条件实现循环
	// processcontrol.BasedOnCondition()
	// break语句和continue语句
	// processcontrol.JumpExample()
	// 标签
	// processcontrol.BreakLabelExample()
	// continue 标签
	// processcontrol.ContinueLabelExample()
	// goto语句
	processcontrol.GotoExample()
}

// 函数调用
func functionCall() {
	// 调用函数
	//function.Myfunc(1, 2, 3, 4, 5, 6)
	//fmt.Println("函数调用")
	//function.Myfunc(1, 2, 3)
	//fmt.Println("变长参数调用...")
	// 变长参数支持传递切片
	//slice := []int{1, 2, 3, 4, 5, 6}
	//function.Myfunc(slice...)
	//function.Myfunc(slice[1:3]...)

	// 可以传递任意类型任意数量的函数
	// function.Myfunc2(1, "1", true)

	// 多返回值
	//x, y := -1, 2
	//z, err := function.AddFunc(&x, &y)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Printf("%d + %d = %d", x, y, z)

	// 命名返回值
	//x, y := -1, 2
	//z, err := function.AddFunc2(&x, &y)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Printf("%d + %d = %d", x, y, z)

	// 匿名函数
	// function.AnonymousFunc()

	// 高阶函数实现装饰器
	function.Multiply(2, 8)
	a := 2
	b := 8
	decorator := function.ExecTime(function.Multiply)
	c := decorator(a, b)
	fmt.Printf("%d x %d = %d\n", a, b, c)
}

func classInitFunc() {
	// 类的初始化
	//student := GoClass.NewStudent(1, "南歌", false, 100)
	//fmt.Println(student)
	//// male未定义，会默认为male的零值false
	//student2 := GoClass.NewStudentPart(1, "南歌", 100)
	//fmt.Println(student2)

	// 成员方法调用
	// 值方法
	//student3 := GoClass.NewStudent(1, "南歌", false, 100)
	//fmt.Println("Name:", student3.GetName())

	// 指针方法
	student4 := GoClass.NewStudent(1, "南歌", false, 100)
	fmt.Println("Name:", student4.GetName())
	student4.SetName("euansu")
	fmt.Println("Name:", student4.GetName())

}

func classFeaturesFunc() {

	//animal := GoClass.Animal{"中华田园犬"}
	//dog := GoClass.Dog{animal}
	////fmt.Println(dog.GetName())
	////fmt.Println(dog.Call())
	////fmt.Println(dog.FavorFood())
	//
	//fmt.Print(dog.Animal.Call())
	//fmt.Println(dog.Call())
	//fmt.Print(dog.Animal.FavorFood())
	//fmt.Println(dog.FavorFood())
}
func classProblems() {
	type Dog struct {
		Animal *GoClass.Animal
		Pet    *GoClass.Pet
	}

	animal := GoClass.Animal{"中华田园犬"}
	pet := GoClass.Pet{"宠物狗"}
	dog := Dog{&animal, &pet}
	fmt.Println(dog.Pet.GetName())
	fmt.Println(dog.Animal.GetName())

	//animal := GoClass.Animal{"中华田园犬"}
	//pet := GoClass.Pet{"宠物狗"}
	//dog := GoClass.Dog{&animal, &pet}
	//fmt.Println(dog.pet.GetName())
	//fmt.Println(dog.animal.GetName())
}

func main() {
	// 变量的代码示例
	// variableFunc()

	// 常量的代码示例
	//constantFunc()

	// 数据类型的代码示例
	// dataTypeFunc()

	// 流程控制语句的代码示例
	// flowControlFunc()

	// 函数调用示例
	// functionCall()

	// 类调用示例
	// classInitFunc()

	// 类的特性
	// classFeaturesFunc()

	// 类的问题
	classProblems()

}
