/*
对文件做拷贝、打印、搜索、排序、统计或类似事情的程序都有一个差不多的程序结构：
	一个处理输入的循环，在每个元素上执行计算处理，在处理的同时或最后产生输出。
*/

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
