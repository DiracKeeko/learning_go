package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))

	// 如果不关心输出格式，只想看看输出值，或许只是为了调试，可以用Println为我们格式化输出。
	//fmt.Println(os.Args[1:])
}
