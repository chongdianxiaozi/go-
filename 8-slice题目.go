package main

import "fmt"

/*func AddElement(slice []int, e int) []int {
	return append(slice, e)
}*/

func main() {
	/*var array [10]int
	var slice = array[5:6]
	fmt.Println("length of slice：", len(slice))
	fmt.Println("capacity of slice：", cap(slice))
	fmt.Println(&slice[0] == &array[5])*/
	fmt.Println("--------------------------")
	/*var slice []int
	slice = append(slice, 1, 2, 3)
	fmt.Println(cap(slice))
	newSlice := AddElement(slice, 4)
	fmt.Println(cap(newSlice))
	fmt.Println(&slice[0] == &newSlice[0])*/
	fmt.Println("--------------------------")

	orderLen := 5
	order := make([]uint16, 2*orderLen)

	pollOrder := order[:orderLen:orderLen]
	lockOrder := order[orderLen:][:orderLen:orderLen]

	fmt.Println("len(pollOrder)=", len(pollOrder))
	fmt.Println("cap(pollOrder)=", cap(pollOrder))
	fmt.Println("len(lockOrder)=", len(lockOrder))
	fmt.Println("cap(lockOrder)=", cap(lockOrder))

}
