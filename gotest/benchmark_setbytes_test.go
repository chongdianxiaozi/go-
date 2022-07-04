package gotest_test

import (
	"testing"
	"time"
)

func BenchmarkSetBytes(b *testing.B) {
	b.SetBytes(1024 * 1024)
	for i := 0; i < b.N; i++ {
		time.Sleep(1 * time.Second) // 模拟待测函数
	}
}
