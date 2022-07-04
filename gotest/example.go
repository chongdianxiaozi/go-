package gotest

import "fmt"

// 打印一行字符串
func SayHello() {
	fmt.Println("Hello go")
}

// 打印两行字符串
func SayGoodbye() {
	fmt.Println("Hello,")
	fmt.Println("goodbye")
}

// 打印姓名
func PrintNames() {
	students := make(map[int]string, 4)
	students[1] = "Kobe"
	students[2] = "Duncan"
	students[3] = "Curry"
	students[4] = "Jingyu"

	for _, value := range students {
		fmt.Println(value)
	}
}
