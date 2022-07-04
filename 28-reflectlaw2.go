package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	var y float64 = v.Interface().(float64)
	fmt.Println("value:", y)
}
