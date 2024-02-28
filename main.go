package main

import (
	"golang_study/constant"
	"golang_study/datatype"
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
	datatype.DataSharing()
}

func main() {
	// 变量的代码示例
	// variableFunc()

	// 常量的代码示例
	//constantFunc()

	// 数据类型的代码示例
	dataTypeFunc()
}
