package main

import "fmt"

func fun1() {
	i := 0
	i, j := 1, 2
	fmt.Printf("i = %d, j = %d\n", i, j) // i = 1, j = 2
}

// 不能通过编译
/*func fun2(i int) {
	i := 0
	fmt.Println(i)
}*/

func fun3() {
	i, j := 0, 0
	if true {
		j, k := 1, 1
		fmt.Printf("j = %d, k = %d\n", j, k)
	}
	fmt.Printf("i = %d, j = %d\n", i, j)
	// 运行结果
	//j = 1, k = 1
	//i = 0, j = 0
}

func main() {
	fun1()
	// fun2(1) // no new variables on left side of :=
	fun3()
}
