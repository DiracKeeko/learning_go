package main

import "fmt"

/* 
	如果 main 包依赖的包中定义了 init 函数，或者是 main 包自身定义了 init 函数，
	那么 Go 程序在这个包初始化的时候，就会自动调用它的 init 函数，
	因此这些 init 函数的执行就都会发生在 main 函数之前。
*/
func init() { 
	fmt.Println("init invoked")
}

// Go 语言中有一个特殊的函数：main 包中的 main 函数，也就是 main.main，它是所有 Go 可执行程序的用户层执行逻辑的入口函数。Go 程序在用户层面的执行逻辑，会在这个函数内按照它的调用顺序展开。
// 和 main.main 函数是一个无参数无返回值的函数
func main() {
	init() // 编译报错 Go中不能手动显式调用 init，否则就会收到编译错误
	
  // 用户层执行逻辑
}