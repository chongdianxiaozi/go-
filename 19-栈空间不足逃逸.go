package main

func Slice() {
	// s := make([]int, 1000, 1000)
	s := make([]int, 10000, 10000)
	for index, _ := range s {
		s[index] = index
	}
}
func main() {
	Slice()
}
