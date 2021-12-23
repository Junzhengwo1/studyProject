package main

import (
	"fmt"
	"strings"
)

/**
闭包的使用
：就是一个函数与其他相关的引用环境组合的一个整体
*/

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func MakeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	//累加器
	upper := AddUpper()
	fmt.Println(upper(2))
	fmt.Println(upper(1))

	suffix := MakeSuffix(".jpg")
	s := suffix("abc")
	fmt.Println(s) //abc.jpg
}
