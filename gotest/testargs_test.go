package gotest_test

import (
	"flag"
	"testing"
)

// 演示如何解析-args参数
func TestArgs(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}

	argList := flag.Args() // 返回 -args 后面的所有参数，以切片表示，每个元素代表一个参数
	for _, arg := range argList {
		if arg == "hello" {
			t.Log("Running in hello.")
		} else {
			t.Log("Running in ...")
		}
	}
}
