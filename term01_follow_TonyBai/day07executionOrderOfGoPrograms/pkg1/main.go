package pkg1

import "fmt"

/* 
	除了 main 包外，其他包也可以拥有自己的名为 main 的函数或方法。
	但按照 Go 的可见性规则（小写字母开头的标识符为非导出标识符），非 main 包中自定义的 main 函数仅限于包内使用，就像下面代码这样，这是一段在非 main 包中定义 main 函数的代码片段：
*/
func Main() {
    main()
}

/* 
	main 函数就主要是用来在包 pkg1 内部使用的，它是没法在包外使用的。
*/
func main() {
    fmt.Println("main func for pkg1")
}
