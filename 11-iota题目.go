package main

import "fmt"

const (
	/*a = 1 << iota
	b
	c
	d = iota
	e = 1e6*/

	bit0, mask0 = 1 << iota, 1<<iota - 1
	bit1, mask1
	_, _
	bit3, mask3
)

func main() {
	/*fmt.Println(a) // 1
	fmt.Println(b) // 2
	fmt.Println(c) // 4
	fmt.Println(d) // 3
	fmt.Println(e) // 1000000*/

	fmt.Println(bit0)
	fmt.Println(mask0)
	fmt.Println(bit1)
	fmt.Println(mask1)
	fmt.Println(bit3)
	fmt.Println(mask3)

}
