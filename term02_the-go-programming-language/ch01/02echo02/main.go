package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", "" // 使用 := 运算符进行短变量声明和初始化。
	// 短变量声明的语法形式是 name := value，在函数内部，短变量声明方式是允许的。

	// ↓ 在Go语言中，_（下划线）被用作匿名变量（anonymous variable）。它用于表示一个值被接收但不被使用，起到占位符的作用。
	for _, arg := range os.Args[1:] {
		// ↑ 这里 (_, arg) (通过短变量声明)接收了 range 的返回值 [索引, 在该索引处的元素值]
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
