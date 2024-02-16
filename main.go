package main

import (
	"golang_study/constant"
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

func main() {
	// 变量的基础代码
	// variableFunc()

	// 常量的基础代码
	constantFunc()
}
