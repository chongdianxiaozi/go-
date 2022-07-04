package gotest_test

import (
	"gotest"
	"testing"
)

// sub1 子测试
func sub1(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3

	actual := gotest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}

}

// sub2 子测试
func sub2(t *testing.T) {
	var a = 2
	var b = 3
	var expected = 5

	actual := gotest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}

}

// sub3 子测试
func sub3(t *testing.T) {
	var a = 3
	var b = 4
	var expected = 7

	actual := gotest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}

}

// TestSub 内部调用sub1、sub2、sub3 三个子测试
func TestSub(t *testing.T) {
	// setup code

	t.Run("A=1", sub1)
	t.Run("A=2", sub2)
	t.Run("B=1", sub3)

	// tear-down code
}
