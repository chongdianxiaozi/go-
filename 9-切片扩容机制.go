package main

func main() {
	s := []int{1, 2}
	s = append(s, 3, 4, 5)
	println(cap(s))
}

// output：6
// no 5 and no 8
