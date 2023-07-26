package main

import (
	"fmt"
	"os"
)

func main() {
	// # 变量声明
	/*
		var 变量名字 类型 = 表达式

			其中“类型”或“= 表达式”两个部分可以省略其中的一个。
			如果省略的是类型信息，那么将根据初始化表达式来推导变量的类型信息。
			如果初始化表达式被省略，那么将用零值初始化该变量。
	*/
	// 零值初始化机制可以确保每个声明的变量总是有一个良好定义的值，因此在Go语言中不存在未初始化的变量。
	var s string
	fmt.Println(s) // ""

	// 也可以在一个声明语句中同时声明一组变量，或用一组初始化表达式声明并初始化一组变量。
	// 如果省略每个变量的类型，将可以声明多个类型不同的变量（类型由初始化表达式推导）：
	var i, j int // int, int
	// var boo, flo, str = true, 2.3, "four" // bool, float64, string

	// # 简短变量声明

	// 在函数内部，有一种称为简短变量声明语句的形式可用于声明和初始化局部变量。
	// 它以“名字 := 表达式”形式声明变量，变量的类型根据表达式来自动推导。
	// 下面是lissajous函数中的三个简短变量声明语句
	/*
		anim := gif.GIF{LoopCount: nframes}
		freq := rand.Float64() * 3.0
		t := 0.0
	*/

	//因为简洁和灵活的特点，简短变量声明被广泛用于大部分的局部变量的声明和初始化。
	//var形式的声明语句往往是用于需要显式指定变量类型的地方，或者因为变量稍后会被重新赋值而初始值无关紧要的地方。

	x := 100                  // an int
	var boiling float64 = 100 // a float64
	var names []string
	var err error
	fmt.Println(x, boiling, names, err)

	i, j = j, i // 交换 i 和 j 的值

	// 和普通var形式的变量声明语句一样，简短变量声明语句也可以用函数的返回值来声明和初始化变量，像下面的os.Open函数调用将返回两个值：
	f, err := os.Open("fileName")
	fmt.Println(f)
}
