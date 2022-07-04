package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	t := reflect.TypeOf(x)
	fmt.Println("type:", t)

	v := reflect.ValueOf(x)
	fmt.Println("value:", v)

}
