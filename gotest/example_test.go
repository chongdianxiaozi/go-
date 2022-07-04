package gotest_test

import "gotest"

// 检测单行输出
func ExampleSayHello() {
	gotest.SayHello()
	// OutPut: Hello go
}

// 检测多行输出
func ExampleSayGoodbye() {
	gotest.SayGoodbye()
	// OutPut:
	// Hello,
	// goodbye
}

// 检测乱序输出
func ExamplePrintNames() {
	gotest.PrintNames()
	// Unordered output:
	// Kobe
	// Duncan
	// Curry
	// Jingyu
}
